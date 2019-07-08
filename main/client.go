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
	go func() {
		for {
			data := make([]byte, 2048)
			n, err := conn.Read(data)
			if err != nil {
				fmt.Println(err)
				readError <- true
				return
			}
			fmt.Println(strings.Trim(string(data[:n]), "\r\n"))
		}
	}()
	go func() {
		for {
			fmt.Println("What to send to the server? Type exit to quit.")
			input, _ := inputReader.ReadString('\n')
			trimmedInput := strings.Trim(input, "\r\n")
			if trimmedInput == "exit" {
				writeExit <- true
				return
			}
			_, err = conn.Write([]byte(trimmedName + " says: " + trimmedInput))
		}
	}()
	select {
	case <-readError:
	case <-writeExit:
	}
}
