package main

import (
	"./client"
	"./server"
	"flag"
)

var (
	ser = flag.Bool("s", false, "run server")
	cli = flag.Bool("c", false, "run client")
)

func main() {
	flag.PrintDefaults()
	flag.Parse() // Scans the arg list and sets up flags
	if *ser {
		server.Server()
	}
	if *cli {
		client.Client()
	}

}
