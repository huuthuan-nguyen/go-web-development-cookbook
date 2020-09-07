package main

import (
	"io"
	"net/http"
	"github.com/gorilla/handlers"
	"log"
)

const (
	CONNECTION_HOST = "localhost"
	CONNECTION_PORT = "8080"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World!")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloWorld)
	err := http.ListenAndServe(CONNECTION_HOST+":"+CONNECTION_PORT, handlers.CompressHandler(mux))
	if err != nil {
		log.Fatal("Error starting http server: ", err)
		return
	}
}