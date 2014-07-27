package haproxy

type Server struct {
	Socket string
}

func (h *Server) GracefulRestart() {

}

//func (h *Haproxy) AddFrontend() {

//}

func (h *Server) WriteConfig(haproxyConfig *GlobalBlock) {

}

func (h *Server) ReadConfig(haproxyConfig *GlobalBlock) {

}
