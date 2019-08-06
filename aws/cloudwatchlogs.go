package aws

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

// create new io.Writer for cloudwatch logs.
func NewLogsWriter(p client.ConfigProvider, groupName string, streamName string) (*logsWriter, error) {
	l := &logsWriter{
		service:       cloudwatchlogs.New(p),
		groupName:     groupName,
		streamName:    streamName,
		sequenceToken: "0",
		events:        make([]*cloudwatchlogs.InputLogEvent, 0),
		eventsLock:    &sync.Mutex{},
		closeSignal:   make(chan bool),
		closeWaiter:   &sync.WaitGroup{},
	}
	// create log group if not exists.
	err := CreateLogGroup(l.service, l.groupName)
	if err != nil {
		return nil, err
	}
	// check latest sequence.
	streams, err := DescribeLogStreams(l.service, l.groupName, l.streamName)
	if err != nil {
		return nil, err
	}
	switch len(streams) {
	case 0:
		err = CreateLogStream(l.service, l.groupName, l.streamName)
		if err != nil {
			return nil, err
		}
	case 1:
		stream := streams[0]
		if stream.UploadSequenceToken != nil {
			l.sequenceToken = *stream.UploadSequenceToken
		}
	default:
		return nil, errors.New("too many streams found.")
	}
	l.closeWaiter.Add(1)
	go l.processEventsLoop()
	return l, nil
}

// declare cloudwatch logs writer.
type logsWriter struct {
	service       *cloudwatchlogs.CloudWatchLogs
	groupName     string
	streamName    string
	sequenceToken string
	events        []*cloudwatchlogs.InputLogEvent
	eventsLock    *sync.Mutex
	closeSignal   chan bool
	closeWaiter   *sync.WaitGroup
}

// implements `io.WriteCloser` interface.
func (l *logsWriter) Write(p []byte) (n int, err error) {
	event := &cloudwatchlogs.InputLogEvent{
		// need millsecond format.
		Timestamp: aws.Int64(time.Now().UTC().UnixNano() / 1000 / 1000),
		Message:   aws.String(string(p)),
	}
	l.pushEvent(event)
	n = len(p)
	err = nil
	return
}

// implements `io.WriteCloser` interface.
func (l *logsWriter) Close() error {
	close(l.closeSignal)
	l.closeWaiter.Wait()
	// flush events.
	l.sendEvents()
	return nil
}

// representation string format.
func (l *logsWriter) String() string {
	return fmt.Sprintf("CloudWatchLogger: group=%s stream=%s sequence=%s events=%d",
		l.groupName, l.streamName, l.sequenceToken, len(l.events))
}

// push new log event.
func (l *logsWriter) pushEvent(event *cloudwatchlogs.InputLogEvent) {
	l.eventsLock.Lock()
	defer l.eventsLock.Unlock()
	l.events = append(l.events, event)
}

// pop log events.
func (l *logsWriter) popEvents() []*cloudwatchlogs.InputLogEvent {
	l.eventsLock.Lock()
	defer l.eventsLock.Unlock()
	events := l.events[:]
	l.events = nil
	return events
}

// send events to cloudwatch logs.
func (l *logsWriter) sendEvents() {
	events := l.popEvents()
	// send events if needed.
	if len(events) < 1 {
		return
	}
	input := &cloudwatchlogs.PutLogEventsInput{
		LogGroupName:  aws.String(l.groupName),
		LogStreamName: aws.String(l.streamName),
		SequenceToken: aws.String(l.sequenceToken),
		LogEvents:     events,
	}
	output, err := l.service.PutLogEvents(input)
	// FIXME: how handle this error?
	if err == nil {
		l.sequenceToken = *output.NextSequenceToken
	}
}

// run send event loop.
func (l *logsWriter) processEventsLoop() {
	defer l.closeWaiter.Done()
	t := time.NewTicker(time.Second)
	defer t.Stop()
	for {
		select {
		case _, ok := <-l.closeSignal:
			if !ok {
				return
			}
		case _ = <-t.C:
			l.sendEvents()
		}
	}
}

// short hand for create log group.
func CreateLogGroup(service *cloudwatchlogs.CloudWatchLogs, name string) error {
	input := &cloudwatchlogs.CreateLogGroupInput{
		LogGroupName: aws.String(name),
	}
	_, err := service.CreateLogGroup(input)
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			switch awsErr.Code() {
			case cloudwatchlogs.ErrCodeResourceAlreadyExistsException:
				return nil
			}
		}
		return err
	}
	return nil
}

// short hand for describe log groups.
func DescribeLogGroups(service *cloudwatchlogs.CloudWatchLogs, name string) ([]*cloudwatchlogs.LogGroup, error) {
	input := &cloudwatchlogs.DescribeLogGroupsInput{
		LogGroupNamePrefix: aws.String(name),
	}
	output, err := service.DescribeLogGroups(input)
	if err != nil {
		return nil, err
	}
	return output.LogGroups, nil
}

// short hand for create log stream.
func CreateLogStream(service *cloudwatchlogs.CloudWatchLogs, groupName string, streamName string) error {
	input := &cloudwatchlogs.CreateLogStreamInput{
		LogGroupName:  aws.String(groupName),
		LogStreamName: aws.String(streamName),
	}
	_, err := service.CreateLogStream(input)
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			switch awsErr.Code() {
			case cloudwatchlogs.ErrCodeResourceAlreadyExistsException:
				return nil
			}
		}
		return err
	}
	return nil
}

// short hand for describe log streams.
func DescribeLogStreams(service *cloudwatchlogs.CloudWatchLogs, groupName string, streamName string) ([]*cloudwatchlogs.LogStream, error) {
	input := &cloudwatchlogs.DescribeLogStreamsInput{
		Limit:               aws.Int64(1),
		LogGroupName:        aws.String(groupName),
		LogStreamNamePrefix: aws.String(streamName),
	}
	output, err := service.DescribeLogStreams(input)
	if err != nil {
		return nil, err
	}
	return output.LogStreams, nil
}
