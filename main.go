package main

import (
	"consoleChatRoomInGolang/client"
	"consoleChatRoomInGolang/server"

	"flag"
	"fmt"
)

var (
	ser = flag.Bool("s", false, "run server")
	cli = flag.Bool("c", false, "run client")
)

func main() {
	flag.PrintDefaults()
	flag.Parse() // Scans the arg list and sets up flags
	if *ser {
		listenPort := new(string)
		fmt.Println("Please input the port you want to listen")
		_, err := fmt.Scanln(listenPort)
		if err != nil {
			panic(err.Error())
		}
		server.Server(listenPort)
		return
	}
	if *cli {
		serverAddr := new(string)
		fmt.Println("Please input the server addr and port like ip:port, for example 127.0.0.1:8080")
		_, err := fmt.Scanln(serverAddr)
		if err != nil {
			panic(err.Error())
		}
		client.Client(serverAddr)
	}

}
