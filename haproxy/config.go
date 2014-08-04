package haproxy

import (
	"io/ioutil"
)

func (h *Server) WriteConfig(configData []string) (err error) {
	configString := strings.Join(configData, "\n")

	err := ioutil.WriteFile("./config/haproxy-test.conf", []byte(configString), 0644)

	return err
}
