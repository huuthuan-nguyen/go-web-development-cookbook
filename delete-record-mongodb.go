package main

import (
	"fmt"
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

func deleteDocument(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]

	if ok {
		log.Print("going to delete document in database for id ::", id)
		collection := session.DB("golangdb").C("employees")
		removeErr := collection.Remove(bson.M{"_id": bson.ObjectIdHex(id)})

		if removeErr != nil {
			log.Print("error removing document from database ::", removeErr)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	} else {
		fmt.Fprintf(w, "Error occurred while deleting document in database for id :: %s", id)
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/employees/{id}", deleteDocument).Methods("DELETE")
	defer session.Close()
	err := http.ListenAndServe(CONNECTION_HOST+":"+CONNECTION_PORT, router)
	if err != nil {
		log.Fatal("error starting http server ::", err)
		return
	}
}