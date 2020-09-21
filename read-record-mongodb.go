package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	CONNECTION_HOST = "localhost"
	CONNECTION_PORT = "8080"
	MONGO_DB_HOST = "127.0.0.1"
)

var session *mgo.Session
var connectionError error

func init() {
	session, connectionError = mgo.Dial(MONGO_DB_HOST)
	if connectionError != nil {
		log.Fatal("error connecting to database ::", connectionError)
	}
	session.SetMode(mgo.Monotonic, true)
}

type Employee struct {
	Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name string `json:"name"`
}

func readDocuments(w http.ResponseWriter, r *http.Request) {
	log.Print("reading documents from database")
	var employees []Employee
	collection := session.DB("golangdb").C("employees")

	err := collection.Find(bson.M{}).All(&employees)
	if err != nil {
		log.Print("error occurre while reading documents from databse ::", err)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	json.NewEncoder(w).Encode(employees)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/employees", readDocuments).Methods("GET")
	defer session.Close()

	err := http.ListenAndServe(CONNECTION_HOST+":"+CONNECTION_PORT, router)
	if err != nil {
		log.Fatal("error starting http server ::", err)
		return
	}
}