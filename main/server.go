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
	defer func() {
		msg := fmt.Sprintf("one client offline, there are %d clients online", len(clients))
		fmt.Println(msg)
		sendMsg(addr, []byte(msg), len(msg))
	}()

	for {
		msg := fmt.Sprintf("one client oneline, there are %d clients online", len(clients))
		fmt.Println(msg)
		sendMsg(addr, []byte(msg), len(msg))
		buf := make([]byte, 2048)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error :", err.Error())
			delete(clients, addr)
			return
		}
		fmt.Printf("Get: %v\n", strings.Trim(string(buf[:n]), "\r\n"))
		sendMsg(addr, buf, n)
	}
}

func sendMsg(addr string, data []byte, length int) {
	for key, value := range clients {
		if key == addr {
			continue
		}
		_, err := value.Write(data[:length])
		if err != nil {
			fmt.Println("Error :", err.Error())
			delete(clients, addr)
			return
		}
	}
}
