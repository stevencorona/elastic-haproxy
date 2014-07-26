package statsd

import (
	"github.com/stevencorona/elastic-haproxy"
)

func main() {
	config := LoadConfig()

	// loop every X seconds
	// read stats from haproxy server
	// post them to statsd
}
