package main

import (
	"flag"

	"github.com/ralpheichelberger/udp/client"
	"github.com/ralpheichelberger/udp/server"
)

func main() {
	var startClient bool
	flag.BoolVar(&startClient, "client", false, "start the client")
	flag.Parse()
	if startClient {
		client.Start()
	} else {
		server.Start()
	}
}
