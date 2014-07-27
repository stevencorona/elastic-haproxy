package haproxy

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"sync"
	"syscall"
)

const HasStarted = 1
const HasStopped = 2
const WantsReload = 4
const WantsStop = 8

type Server struct {
	Socket     string
	ActionChan chan int
	cmd        *exec.Cmd
	sync.RWMutex
}

func (h *Server) createProcess() {
	h.cmd = exec.Command("/usr/local/bin/haproxy", "-f", "config/haproxy.conf")
}

func (h *Server) setupStdout() {
	h.cmd.Stdout = os.Stdout
	h.cmd.Stderr = os.Stderr
}

func (h *Server) runProcess() error {
	return h.cmd.Start()
}

func (h *Server) Start(notify chan int, action chan int) {
	h.Lock()
	h.createProcess()
	h.setupStdout()

	err := h.runProcess()

	if err != nil {
		log.Fatal(err)
	}

	h.ActionChan = action
	h.Unlock()

	notify <- HasStarted

	// Wait for a stop signal, reload signal, or process death
	for {

		wants := <-action
		h.Lock()

		switch wants {
		case WantsReload:
			log.Println("Replacing process")
			h.reloadProcess()
			h.Unlock()
			notify <- 1
			log.Println("Process has been replaced")
		case WantsStop:
			h.stopProcess()
			h.Unlock()
			break
		}
	}

	notify <- HasStopped
}

func (h *Server) reloadProcess() error {
	// Grab pid of current running process
	pid := strconv.Itoa(h.cmd.Process.Pid)

	// Start a new process, telling it to replace the old process
	cmd := exec.Command("/usr/local/bin/haproxy", "-f", "config/haproxy.conf", "-sf", pid)

	// Start the new process and check for errors. We bail out if there is
	// an error and DON'T replace the old process.
	err := cmd.Start()

	if err != nil {
		return err
	}

	// No errors? Replace the old process
	h.cmd = cmd

	return nil
}

func (h *Server) stopProcess() error {
	return h.cmd.Process.Signal(syscall.SIGUSR1)
}
