package main

import (
	"fmt"
	"net"
	"strings"

	"go-micro.dev/v4/util/addr"
	mnet "go-micro.dev/v4/util/net"
)

func main() {
	advt := "[::]:5002"
	var err error
	var host, port string
	var cacheService bool
	if cnt := strings.Count(advt, ":"); cnt >= 1 {
		// ipv6 address in format [host]:port or ipv4 host:port
		host, port, err = net.SplitHostPort(advt)
		if err != nil {
			return
		}
	} else {
		host = advt
	}

	fmt.Println(host)
	fmt.Println(port)

	if ip := net.ParseIP(host); ip != nil {
		cacheService = true
	}

	addr, err := addr.Extract(host)
	if err != nil {
		return
	}

	// mq-rpc(eg. nats) doesn't need the port. its addr is queue name.
	if port != "" {
		addr = mnet.HostPort(addr, port)
	}

	fmt.Println(addr)
	fmt.Println(cacheService)
}
