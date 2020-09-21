package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	CONNECTION_HOST = "localhost"
	CONNECTION_PORT = "8080"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	http.HandleFunc("/", helloWorld)
	err := http.ListenAndServe(CONNECTION_HOST+":"+CONNECTION_PORT, nil)
	if err != nil {
		log.Fatal("error starting http server :", err)
		return
	}
}