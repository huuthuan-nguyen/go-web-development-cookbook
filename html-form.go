package main

import (
	"html/template"
	"log"
	"net/http"
)

const (
	CONNECTION_HOST = "localhost"
	CONNECTION_PORT = "8080"
)

func login(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, _ := template.ParseFiles("templates/login-form.html")
	parsedTemplate.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", login)
	err := http.ListenAndServe(CONNECTION_HOST+":"+CONNECTION_PORT, nil)
	if err != nil {
		log.Fatal("Error starting http server: ", err)
		return
	}
}