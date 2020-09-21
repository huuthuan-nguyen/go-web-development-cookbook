package main

import (
	"fmt"
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
	Id bson.ObjectId `bson:"_id:omitempty" json:"id"`
	Name string `json:"name"`
}

func init() {
	session, connectionError = mgo.Dial(MONGO_DB_HOST)
	if connectionError != nil {
		log.Fatal("error connecting to databse ::", connectionError)
	}
	session.SetMode(mgo.Monotonic, true)
}

func updateDocument(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]

	employee := Employee{}
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		log.Printf("error while decoding employee data ::", err)
		return
	}

	if ok {
		log.Print("going to update document in database for id ::", id)
		collection := session.DB("golangdb").C("employees")
		var changeInfo *mgo.ChangeInfo
		changeInfo, err = collection.UpsertId(bson.ObjectIdHex(id), bson.M{"$set": bson.M{"name": employee.Name}})

		if err != nil {
			log.Print("error occurred while updating record in database ::", err)
			return
		}
		fmt.Fprintf(w, "Number of documents updated in database are :: %d", changeInfo.Updated)
	} else {
		fmt.Fprintf(w, "Error occurred while updating document in database for id :: %s", id)
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/employees/{id}", updateDocument).Methods("PUT")
	defer session.Close()
	err := http.ListenAndServe(CONNECTION_HOST+":"+CONNECTION_PORT, router)
	if err != nil {
		log.Fatal("error starting http server ::", err)
		return
	}
}