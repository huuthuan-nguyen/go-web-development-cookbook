package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
)

const (
	CONNECTION_HOST = "localhost"
	CONNECTION_PORT = "8080"
	MONGO_DB_HOST = "127.0.0.1"
)

var session *mgo.Session
var connectionError error

type Employee struct {
	Id bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Name string `json:"name"`
}

func init() {
	session, connectionError = mgo.Dial(MONGO_DB_HOST)
	if connectionError != nil {
		log.Fatal("error connecting to database ::", connectionError)
	}
	session.SetMode(mgo.Monotonic, true)
}

func createDocument(w http.ResponseWriter, r *http.Request) {
	employee := Employee{}
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		log.Print("error occurred while decoding employee data ::", err)
		return
	}

	log.Printf("adding employee id :: %s with name as :: %s", employee.Id, employee.Name)
	collection := session.DB("golangdb").C("employees")

	bsonEmployee := Employee{bson.NewObjectId(), employee.Name}

	err = collection.Insert(&bsonEmployee)
	if err != nil {
		log.Print("error occurred while inserting document in database ::", err)
		return
	}
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	json.NewEncoder(w).Encode(bsonEmployee)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/employees", createDocument).Methods("POST")
	defer session.Close()
	err := http.ListenAndServe(CONNECTION_HOST+":"+CONNECTION_PORT, router)
	if err != nil {
		log.Fatal("error starting http server ::", err)
		return
	}
}