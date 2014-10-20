package haproxy

import (
	"io/ioutil"
	"strings"
)

// This needs to be changed to spit out JSON into our own format and just
// run the config data through the templating engine
func (h *Server) WriteConfig(configData []string) (err error) {
	configString := strings.Join(configData, "\n")

	err = ioutil.WriteFile("./config/haproxy-test.conf", []byte(configString), 0644)

	return err
}
