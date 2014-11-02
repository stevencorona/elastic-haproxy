package haproxy

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"sync"
	"time"
)

type Event int
type Action int

// Events that can be fired from the event channel
const (
	HasStarted Event = 1 << iota
	HasStopped
	HasReloaded
)

// Actions that can be received from the action channel
const (
	WantsReload Action = 1 << iota
	WantsStop
)

type Server struct {
	Socket     string
	ActionChan chan Action
	cmd        *exec.Cmd
	sync.RWMutex
}

const BINARYPATH = "/usr/local/bin/haproxy"
const CONFIGPATH = "config/haproxy.conf"

func (h *Server) createProcess() {
	h.cmd = exec.Command(BINARYPATH, "-f", CONFIGPATH)
}

func (h *Server) setupStdout() {
	h.cmd.Stdout = os.Stdout
	h.cmd.Stderr = os.Stderr
}

func (h *Server) runProcess() error {
	return h.cmd.Start()
}

func (h *Server) Start(notify chan Event, action chan Action) {
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
		switch <-action {
		case WantsReload:
			h.reloadProcess()
			notify <- HasReloaded
		case WantsStop:
			h.stopProcess()
			notify <- HasStopped
			return
		}
	}
}

func (h *Server) reloadProcess() error {
	h.Lock()

	// Grab pid of current running process
	old := h.cmd
	pid := strconv.Itoa(h.cmd.Process.Pid)

	// Start a new process, telling it to replace the old process
	// This will signal the current process and tell it shutdown, which is why
	// we don't need to do it here.
	cmd := exec.Command(BINARYPATH, "-f", CONFIGPATH, "-sf", pid)

	// Start the new process and check for errors. We bail out if there is
	// an error and DON'T replace the old process.
	err := cmd.Start()

	if err != nil {
		return err
	}

	// No errors? Replace the old process
	h.cmd = cmd
	h.finishOrKill(10*time.Second, old)

	return nil
}

func (h *Server) finishOrKill(waitFor time.Duration, old *exec.Cmd) {
	// Create a channel and wait for the old process
	// to finish itself
	didFinish := make(chan error, 1)
	go func() {
		didFinish <- old.Wait()
	}()

	// Wait for the didFinish channel or force kill the process
	// if it takes longer than 10 seconds
	select {
	case <-time.After(waitFor):
		log.Println("manually killing process")
		if err := old.Process.Kill(); err != nil {
			log.Println("failed to kill ", err)
		}
	case err := <-didFinish:
		if err != nil {
			log.Println("process finished with error", err)
		}
	}
}

func (h *Server) stopProcess() error {
	return h.cmd.Process.Kill()
}
