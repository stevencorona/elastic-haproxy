package statsd

import (
	"fmt"
	"github.com/cyberdelia/statsd"
	"github.com/stevencorona/elastic-haproxy/haproxy"
	"time"
)

func SendMetrics(server *haproxy.Server) {

	c, err := statsd.Dial("localhost:8125")

	fmt.Println(err)

	for {
		info := server.GetInfo()
		c.Gauge("current_connections", info.CurrConns, 1)
		c.Gauge("cum_connections", info.CumConns, 1)
		fmt.Println("current_connections", info.CurrConns)
		fmt.Println("cum_connections", info.CumConns)
		c.Flush()
		time.Sleep(1 * time.Second)
	}
}
