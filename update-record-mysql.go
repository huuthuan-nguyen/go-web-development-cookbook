package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"strconv"
)

const (
	CONNECTION_HOST = "localhost"
	CONNECTION_PORT = "8080"
	DRIVER_NAME = "mysql"
	DATA_SOURCE_NAME = "root:@/golangdb"
)

var db *sql.DB
var connectionError error

func init() {
	db, connectionError = sql.Open(DRIVER_NAME, DATA_SOURCE_NAME)
	if connectionError != nil {
		log.Fatal("error connecting to database ::", connectionError)
	}
}

type Employee struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

func updateRecord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	vals := r.URL.Query()
	name, ok := vals["name"]
	if ok {
		log.Print("going to update record in database for id ::", id)
		stmt, err := db.Prepare("UPDATE employees SET name=? where id=?")
		if err != nil {
			log.Print("error occurred while executing query ::", err)
			return
		}
		result, err := stmt.Exec(name[0], id)
		if err != nil {
			log.Print("error occurred while executing query ::", err)
			return
		}
		rowsAffected, err := result.RowsAffected()
		fmt.Fprintf(w, "Number of rows updated in database are :: %s", strconv.FormatInt(rowsAffected, 10))
	} else {
		fmt.Fprintf(w, "Error occurred while updating record in database for id :: %s", id)
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/employees/{id}", updateRecord).Methods("PUT")
	defer db.Close()
	err := http.ListenAndServe(CONNECTION_HOST+":"+CONNECTION_PORT, router)
	if err != nil {
		log.Fatal("error starting http server ::", err)
		return
	}
}