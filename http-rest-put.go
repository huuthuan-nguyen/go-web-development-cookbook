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
		"GET",
		"/employees",
		getEmployees,
	},
	Route{
		"addEmployee",
		"POST",
		"/employees",
		addEmployee,
	},
	Route{
		"updateEmployee",
		"PUT",
		"/employees/{id}",
		updateEmployee,
	},
}

type Employee struct {
	Id string `json:"id"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
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
			"Quz",
		},
	}
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	json.NewEncoder(w).Encode(employees)
}

func updateEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	employee := Employee{}
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		log.Printf("error occurred while decoding employee data :: ", err)
		return
	}

	for idx, employeex := range employees {
		if employeex.Id == id {
			log.Printf("updating employee id :: %s with firstName as :: %s and lastName as :: %s", id, employee.FirstName, employee.LastName)
			employees[idx].FirstName = employee.FirstName
			employees[idx].LastName = employee.LastName
			employee.Id = id
			w.Header().Set("Content-Type", "application/json;charset=utf-8")
			json.NewEncoder(w).Encode(employee)
			break
		}
	}
}

func addEmployee(w http.ResponseWriter, r *http.Request) {
	employee := Employee{}
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		log.Printf("error while decoding employee data :: ", err)
		return
	}

	log.Printf("adding employee id :: %s with firstName as :: %s and lastName as :: %s", employee.Id, employee.FirstName, employee.LastName)
	employees = append(employees, employee)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	json.NewEncoder(w).Encode(employee)
}

func AddRoutes(router *mux.Router) *mux.Router {
	for _, route := range routes {
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).HandlerFunc(route.HandlerFunc)
	}
	return router
}

func main() {
	muxRouter := mux.NewRouter().StrictSlash(true)
	router := AddRoutes(muxRouter)
	err := http.ListenAndServe(CONNECTION_HOST+":"+CONNECTION_PORT, router)
	if err != nil {
		log.Fatal("error starting http server :: ", err)
		return
	}
}