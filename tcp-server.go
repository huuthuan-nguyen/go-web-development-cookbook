package main

import (
	"log"
	"net"
)

const (
	CONNECTION_HOST = "localhost"
	CONNECTION_PORT = "8080"
	CONNECTION_TYPE = "tcp"
)

func main() {
	listener, err := net.Listen(CONNECTION_TYPE, CONNECTION_HOST+":"+CONNECTION_PORT)
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
		log.Println(conn)
	}
}