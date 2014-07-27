package haproxy

import (
	"log"
	"os"
	"os/exec"
)

type Server struct {
	Socket  string
	cmd     *exec.Cmd
	process *os.Process
}

func (h *Server) Start(start chan int, stop chan int) {
	h.cmd = exec.Command("/usr/local/bin/haproxy", "-f", "config/haproxy.conf")
	h.cmd.Stdout = os.Stdout
	h.cmd.Stderr = os.Stderr

	err := h.cmd.Start()

	if err != nil {
		log.Fatal(err)
	}

	h.process = h.cmd.Process

	start <- 1
	err = h.cmd.Wait()
	log.Println(err)
	stop <- 1
}

func (h *Server) GracefulRestart() {

}

//func (h *Haproxy) AddFrontend() {

//}

func (h *Server) WriteConfig(haproxyConfig *GlobalBlock) {

}

func (h *Server) ReadConfig(haproxyConfig *GlobalBlock) {

}
