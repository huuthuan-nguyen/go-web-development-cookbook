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

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Login!")
}

func logout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Logout!")
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	err := http.ListenAndServe(CONNECTION_HOST+":"+CONNECTION_PORT, nil)
	if err != nil {
		log.Fatal("Error starting http server: ", err)
		return
	}
}