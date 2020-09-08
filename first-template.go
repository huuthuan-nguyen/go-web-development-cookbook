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
	http.HandleFunc("/", renderTemplate)
	err := http.ListenAndServe(CONNECTION_HOST+":"+CONNECTION_PORT, nil)
	if err != nil {
		log.Fatal("Error starting http server: ", err)
		return
	}
}