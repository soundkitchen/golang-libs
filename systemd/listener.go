package systemd

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

// the first passed file descriptor is fd 3
// defined in `systemd/src/systemd/sd-daemon.h`
// see also "https://github.com/systemd/systemd/blob/master/src/systemd/sd-daemon.h"
const listenFdStart = 3

// create new listener with systemd socket environment.
// unsupported multiple listeners yet.
func Listener() (net.Listener, error) {
	defer os.Unsetenv("LISTEN_PID")
	defer os.Unsetenv("LISTEN_FDS")
	defer os.Unsetenv("LISTEN_FDNAMES")
	// check pid.
	pid, err := strconv.Atoi(os.Getenv("LISTEN_PID"))
	if err != nil {
		return nil, err
	}
	cPid := os.Getpid()
	if pid != cPid {
		return nil, fmt.Errorf("mismatch pid. expected %d but actual %d.", pid, cPid)
	}
	// check listeners.
	nfds, err := strconv.Atoi(os.Getenv("LISTEN_FDS"))
	if err != nil {
		return nil, err
	}
	if nfds < 1 {
		return nil, fmt.Errorf("can't find any listeners.")
	}
	if nfds > 1 {
		return nil, fmt.Errorf("unsupport multiple listeners.")
	}
	// create new listener.
	fp := os.NewFile(uintptr(listenFdStart), "LISTEN_FD_0")
	defer fp.Close()
	return net.FileListener(fp)
}
