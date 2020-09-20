package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

const (
	CONNECTION_HOST = "localhost"
	CONNECTION_PORT = "8080"
)

type Route struct {
	Name string
	Method string
	Pattern string
	HandlerFunc http.HandlerFunc
}

type Routes []Route
var routes = Routes{
	Route{
		"getEmployees",
		"GET"
		"/employees"
		getEmployees,
	},
	Route{
		"addEmployee",
		"POST",
		"/employees",
		addEmployee,
	},
	Route{
		"deleteEmployee",
		"DELETE",
		"/employees/{id}",
		deleteEmployee,
	},
}

type Employee struct {
	Id string `json:"id"`
	FirstName string
	LastName string
}
type Employees []Employee
var employees Employees

func init() {
	employees = Employees{
		Employee{
			"1",
			"Foo",
			"Bar",
		},
		Employee{
			"2",
			"Baz",
			"Qux",
		},
	}
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(employees)
}

func deleteEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for idx, _ := range employees {
		if idx == id {
			log.Printf("deleting employee id :: %s with firstName as :: %s and lastName as :: %s", idx, employees[idx].FirstName, employees[idx].LastName)
			employees = append(employees[:idx], employees[idx+1:]...)
			w.WriteHeader(http.StatusNoContent)
			break
		}
	}
}

func addEmployee(w http.ResponseWriter, r *http.Request) {
	employee := Employee{}
	err := json.NewDecoder(r.Body).Decode(&employee)
}
