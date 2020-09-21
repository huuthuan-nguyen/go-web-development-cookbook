package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
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

func deleteRecord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]

	if ok {
		log.Print("going to delete record in database for id ::", id)
		stmt, err := db.Prepare("DELETE FROM employees WHERE id=?")
		if err != nil {
			log.Print("error occurred while preparing query ::", err)
			return
		}

		result, err := stmt.Exec(id)
		if err != nil {
			log.Print("error occurred while executing query ::", err)
			return
		}

		rowsAffected, err := result.RowsAffected()
		fmt.Fprintf(w, "Number of rows deleted in database are :: %d", rowsAffected)
	} else {
		fmt.Fprintf(w, "Error occurred while deleting record in database for id %s", id)
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/employees/{id}", deleteRecord).Methods("DELETE")
	defer db.Close()
	err := http.ListenAndServe(CONNECTION_HOST+":"+CONNECTION_PORT, router)
	if err != nil {
		log.Fatal("error starting http server ::", err)
		return
	}
}