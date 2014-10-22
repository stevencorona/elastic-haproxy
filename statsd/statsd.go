package statsd

import (
	"fmt"
	"github.com/cyberdelia/statsd"
	"github.com/stevencorona/elastic-haproxy/haproxy"
	"time"
)

func SendMetrics(server *haproxy.Server) {

	c, err := statsd.Dial("localhost:8125")

	// TODO: Instead of dying completely, should put this on a retry
	// and try again.
	if err != nil {
		log.Fatal("Could not connect to statsd")
	}

	for {
		info := server.GetInfo()
		c.Gauge("current_connections", info.CurrConns, 1)
		c.Gauge("cum_connections", info.CumConns, 1)
		c.Flush()
		time.Sleep(1 * time.Second)
	}
}
