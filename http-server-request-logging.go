package main

import (
	"net/http"
	"os"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
)

const (
	CONNECTION_HOST = "localhost"
	CONNECTION_PORT = "8080"
)

var GetRequestHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	},
)

var PostRequestHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("It's a Post Request!"))
	},
)

var PathVariableHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]
		w.Write([]byte("Hi " + name))
	},
)

func main() {
	router := mux.NewRouter()
	router.Handle("/", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(GetRequestHandler))).Methods("GET")
	logFile, err := os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Error starting http server: ", err)
		return
	}

	router.Handle("/post", handlers.LoggingHandler(logFile, PostRequestHandler)).Methods("POST")
	router.Handle("/hello/{name}", handlers.LoggingHandler(logFile, PathVariableHandler)).Methods("GET", "PUT")
	http.ListenAndServe(CONNECTION_HOST+":"+CONNECTION_PORT, router)
}