package main

import (
	"flag"

	"github.com/kuche1/browser-chess/client"
	"github.com/kuche1/browser-chess/server"
)

func main() {
	runServer := flag.Bool("run-server", false, "Run server")
	flag.Parse()

	if *runServer {
		server.Main()
	} else {
		client.Main()
	}
}
