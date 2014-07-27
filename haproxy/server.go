package haproxy

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"sync"
	"syscall"
)

type Server struct {
	Socket string
	cmd    *exec.Cmd
	sync.RWMutex
}

func (h *Server) Start(start chan int, stop chan int) {
	h.Lock()

	if h.cmd == nil {
		h.cmd = exec.Command("/usr/local/bin/haproxy", "-f", "config/haproxy.conf")
	} else {
		pid := strconv.Itoa(h.cmd.Process.Pid)
		h.cmd = exec.Command("/usr/local/bin/haproxy", "-f", "config/haproxy.conf", "-sf", pid)
	}

	h.cmd.Stdout = os.Stdout
	h.cmd.Stderr = os.Stderr

	err := h.cmd.Start()

	if err != nil {
		log.Fatal(err)
	}

	h.Unlock()

	start <- 1
	err = h.cmd.Wait()

	if err != nil {
		log.Println(err)
	}

	stop <- 1
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
