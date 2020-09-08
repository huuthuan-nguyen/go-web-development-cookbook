package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"github.com/gorilla/schema"
)

const (
	CONNECTION_HOST = "localhost"
	CONNECTION_PORT = "8080"
)

type User struct {
	Username string
	Password string
}

func readForm(r *http.Request) *User {
	r.ParseForm()
	user := new(User)
	decoder := schema.NewDecoder()
	decodeErr := decoder.Decode(user, r.PostForm)
	if decodeErr != nil {
		log.Printf("Error mapping parsed form data to struct: ", decodeErr)
	}
	return user
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		parsedTemplate, _ := template.ParseFiles("templates/login-form.html")
		parsedTemplate.Execute(w, nil)
	} else {
		user := readForm(r)
		fmt.Fprintf(w, "Hello " + user.Username + "!")
	}
}

func main() {
	http.HandleFunc("/", login)
	err := http.ListenAndServe(CONNECTION_HOST+":"+CONNECTION_PORT, nil)
	if err != nil {
		log.Fatal("Error starting http server: ", err)
		return
	}
}