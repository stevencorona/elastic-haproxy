# Elastic HAProxy

Amazon ELB is awesome. It's cheap and easy to get started with, but can be limiting for advanced or high performance architectures. Sometimes you just need a little bit more control.

**Elastic HAProxy is a modern wrapper for HAProxy.**

* Mostly compatible ELB HTTP Api
* Reports key Haproxy metrics to Statsd
* Dynamically update frontends and backends (with zero downtime reloads)
* Multi-node coordination, scale out to N load balancers
* Etcd, Route53, and Docker support
* Angular dashboard with full control

All of this is done outside of the HAProxy code path, so it still maintains the high performance that comes baked in with HAProxy. Think of Elastic HAProxy as a bolt-on to HAProxy, running in it's own isolated process.

*Currently in Alpha/Proof of Concept Phase*

## Architecture

The main method of implementation is to wrap the HAProxy binary and execute it from a goroutine. We normalize the configuration as a bunch of structs and serialize the output as the HAProxy config.

Zero downtime reconfigurations are accomplished using the `-sf` trick and (if on linux) some iptables magic.
