package haproxy

import (
	"log"
	"os"
	"os/exec"
	"sync"
	"syscall"
)

type Server struct {
	Socket  string
	cmd     *exec.Cmd
	process *os.Process
	sync.RWMutex
}

func (h *Server) Start(start chan int, stop chan int) {
	h.Lock()

	h.cmd = exec.Command("/usr/local/bin/haproxy", "-f", "config/haproxy.conf")
	h.cmd.Stdout = os.Stdout
	h.cmd.Stderr = os.Stderr

	err := h.cmd.Start()

	if err != nil {
		log.Fatal(err)
	}

	h.process = h.cmd.Process
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

	err := h.process.Signal(syscall.SIGUSR1)
	if err != nil {
		log.Println(err)
	}
}

func (h *Server) GracefulRestart() {

}

//func (h *Haproxy) AddFrontend() {

//}

func (h *Server) WriteConfig(haproxyConfig *GlobalBlock) {

}

func (h *Server) ReadConfig(haproxyConfig *GlobalBlock) {

}
