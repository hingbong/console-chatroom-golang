package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func Client() {
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("Error dialing", err.Error())
		return
	}
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("What is your name?")
	name, _ := inputReader.ReadString('\n')
	trimmedName := strings.Trim(name, "\r\n")
	readError := make(chan bool)
	writeExit := make(chan bool)
	go func(connection net.Conn, sem chan<- bool) {
		for {
			data := make([]byte, 2048)
			n, err := connection.Read(data)
			if err != nil {
				fmt.Println(err)
				sem <- true
				return
			}
			fmt.Println(strings.Trim(string(data[:n]), "\r\n"))
		}
	}(conn, readError)
	go func(connection net.Conn, name string, reader *bufio.Reader, sem chan<- bool) {
		for {
			fmt.Println("What to send to the server? Type exit to quit.")
			input, _ := reader.ReadString('\n')
			trimmedInput := strings.Trim(input, "\r\n")
			if trimmedInput == "exit" {
				sem <- true
				return
			}
			_, err = connection.Write([]byte(name + " says: " + trimmedInput))
		}
	}(conn, trimmedName, inputReader, writeExit)
	select {
	case <-readError:
	case <-writeExit:
	}
}
