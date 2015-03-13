# Elastic HAProxy

Amazon ELB is awesome. It's cheap and easy to get started with, but can be limiting for advanced or high performance architectures. Sometimes you just need a little bit more control.

**Elastic HAProxy is a modern wrapper for HAProxy.**

* Mostly compatible ELB HTTP Api
* Reports key HAproxy metrics to Statsd
* Dynamically update frontends and backends (with zero downtime reloads)
* Multi-node coordination, scale out to N load balancers
* Etcd, Route53, and Docker support
* Angular dashboard with full control
* Zero-downtime configuration

All of this is done outside of the HAProxy code path, so it still maintains the high performance that comes baked in with HAProxy. Think of Elastic HAProxy as a bolt-on to HAProxy, running in it's own isolated process.

**Currently in Alpha/Proof of Concept Phase. Do not use in production!**

## Architecture

The main method of implementation is to wrap the HAProxy binary and execute it from a goroutine. We normalize the configuration as a bunch of structs and serialize the output as the HAProxy config.

It is completely out-of-band from the request cycle and uses the stock HAProxy build.

Zero downtime reconfigurations are accomplished using the `-sf` trick and (if on linux) some iptables magic.

## License

    The MIT License (MIT)

    Copyright (c) 2014 Steve Corona

    Permission is hereby granted, free of charge, to any person obtaining a copy
    of this software and associated documentation files (the "Software"), to deal
    in the Software without restriction, including without limitation the rights
    to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
    copies of the Software, and to permit persons to whom the Software is
    furnished to do so, subject to the following conditions:

    The above copyright notice and this permission notice shall be included in all
    copies or substantial portions of the Software.

    THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
    IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
    FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
    AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
    LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
    OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
    SOFTWARE.
