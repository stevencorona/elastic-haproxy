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

type Server struct {
	Socket string
	Reload chan int
	cmd    *exec.Cmd
	sync.RWMutex
}

func (h *Server) Start(notify chan int, reload chan int) {
	h.Lock()

	h.cmd = exec.Command("/usr/local/bin/haproxy", "-f", "config/haproxy.conf")
	h.cmd.Stdout = os.Stdout
	h.cmd.Stderr = os.Stderr

	err := h.cmd.Start()

	if err != nil {
		log.Fatal(err)
	}

	h.Reload = reload

	h.Unlock()

	notify <- HasStarted

	for {

		<-reload

		h.Lock()

		oldPid := strconv.Itoa(h.cmd.Process.Pid)
		newCmd := exec.Command("/usr/local/bin/haproxy", "-f", "config/haproxy.conf", "-sf", oldPid)
		h.cmd = newCmd
		err := h.cmd.Start()

		if err != nil {
			log.Fatal(err)
		}

		h.Unlock()

	}

	notify <- HasStopped
}

func (h *Server) Shutdown() {
	h.Lock()
	defer h.Unlock()

	err := h.cmd.Process.Signal(syscall.SIGUSR1)
	if err != nil {
		log.Println(err)
	}
}

func (h *Server) GracefulRestart(start chan int, stop chan int) {
	h.Lock()
	defer h.Unlock()

	h.Start(start, stop)

}

//func (h *Haproxy) AddFrontend() {

//}

func (h *Server) WriteConfig(haproxyConfig *GlobalBlock) {

}

func (h *Server) ReadConfig(haproxyConfig *GlobalBlock) {

}
