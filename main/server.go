package main

import (
	"fmt"
	"net"
	"strings"
)

var clients = map[string]net.Conn{}

func Server() {
	conn, _ := net.Listen("tcp", "127.0.0.1:8888")
	for {
		accept, e := conn.Accept()
		if e != nil {
			fmt.Println(e)
			return
		}
		addr := accept.RemoteAddr().String()
		clients[addr] = accept
		go run(addr, accept)
	}
}

func run(addr string, conn net.Conn) {
	for {
		buf := make([]byte, 2048)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error :", err.Error())
			delete(clients, addr)
			return
		}
		fmt.Printf("Get: %v\n", strings.Trim(string(buf[:n]), "\r\n"))
		for key, value := range clients {
			if key == addr {
				continue
			}
			_, err := value.Write(buf[:n])
			if err != nil {
				fmt.Println("Error :", err.Error())
				delete(clients, addr)
				return
			}
		}
	}
}
