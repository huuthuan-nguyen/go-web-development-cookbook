package main

import (
	"html/template"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

const (
	CONNECTION_HOST = "localhost"
	CONNECTION_PORT = "8080"
)

type Person struct {
	Id string
	Name string
}

func renderTemplate(w http.ResponseWriter, r *http.Request) {
	person := Person{Id: "1", Name: "Foo"}
	parsedTemplate, _ := template.ParseFiles("templates/first-template.html")
	err := parsedTemplate.Execute(w, person)
	if err != nil {
		log.Printf("Error occurred while executing the template or writing its output: ", err)
		return
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", renderTemplate).Methods("GET")
	router.PathPrefix("/").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("static/"))))
	err := http.ListenAndServe(CONNECTION_HOST+":"+CONNECTION_PORT, router)
	if err != nil {
		log.Fatal("Error starting http server: ", err)
		return
	}
}