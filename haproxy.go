package main

type Haproxy struct {
	Socket string
}

func (h *Haproxy) GracefulRestart() {

}

//func (h *Haproxy) AddFrontend() {

//}

func (h *Haproxy) WriteConfig(haproxyConfig *GlobalBlock) {

}

func (h *Haproxy) ReadConfig(haproxyConfig *GlobalBlock) {

}
