package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

const (
	CONNECTION_HOST = "localhost"
	CONNECTION_PORT = "8080"
	CONNECTION_TYPE = "tcp"
)

func handleRequest(conn net.Conn) {
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading: ", err.Error())
	}
	fmt.Print("Message Received: ", string(message))
	conn.Write([]byte(message + "\n"))
	conn.Close()
}

func main() {
	listener, err := net.Listen(CONNECTION_TYPE, CONNECTION_HOST + ":" + CONNECTION_PORT)
	if err != nil {
		log.Fatal("Error starting tcp server: ", err)
	}
	defer listener.Close()
	log.Println("Listening on " + CONNECTION_HOST + ":" + CONNECTION_PORT)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Error accepting: ", err.Error())
		}
		go handleRequest(conn)
	}
}