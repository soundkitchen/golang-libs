package errors

import (
	"fmt"
)

// store multiple errors.
type MultiError map[string]error

// implements Error interface.
func (m MultiError) Error() string {
	l := len(m)
	var extra string
	if l > 1 {
		extra = fmt.Sprintf(" (and %d others)", l-1)
	}
	var message string
	for k, err := range m {
		message = fmt.Sprintf("%s: %s%s", k, err, extra)
		break
	}
	return message
}
