package node

import (
	"go.uber.org/atomic"
	"golang.org/x/sys/unix"
	"syscall"
)

type PHPAppNode struct {
	pid int

	agent *AppAgent

	stopped atomic.Bool
}

func (a *PHPAppNode) Pid() int {
	return a.pid
}

func (a *PHPAppNode) Stop() {
	if a.stopped.CAS(false, true) {
		a.agent.Close()
	}
}

func (a *PHPAppNode) ForceStop() {
	if a.pid != 0 {
		err := syscall.Kill(a.pid, syscall.SIGKILL)
		if err != unix.ESRCH {

		}
	}
}


