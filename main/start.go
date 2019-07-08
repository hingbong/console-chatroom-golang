package main

import (
	"flag"
)

var (
	serv = flag.Bool("s", false, "run server")
	cli  = flag.Bool("c", false, "run client")
)

func main() {
	flag.PrintDefaults()
	flag.Parse() // Scans the arg list and sets up flags
	if *serv {
		Server()
	}
	if *cli {
		Client()
	}

}
