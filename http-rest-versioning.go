package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
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
}

type Employee struct {
	Id string `json:"id"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}

type Employees []Employee
var employees Employees
var employeesV1 Employees
var employeesV2 Employees

func init() {
	employees = Employees{
		Employee{
			"1",
			"Foo",
			"Bar",
		},
	}
	employeesV1 = Employees{
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
	employeesV2 = Employees{
		Employee{
			"1",
			"Baz",
			"Qux",
		},
		Employee{
			"2",
			"Quux",
			"Quuz",
		},
	}
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	if strings.HasPrefix(r.URL.Path, "/v1") {
		json.NewEncoder(w).Encode(employeesV1)
	} else if strings.HasPrefix(r.URL.Path, "/v2") {
		json.NewEncoder(w).Encode(employeesV2)
	} else {
		json.NewEncoder(w).Encode(employees)
	}
}

func AddRoutes(router *mux.Router) *mux.Router {
	for _, route := range routes {
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(route.HandlerFunc)
	}
	return router
}

func main() {
	muxRouter := mux.NewRouter().StrictSlash(true)
	router := AddRoutes(muxRouter)

	// v1
	AddRoutes(muxRouter.PathPrefix("/v1").Subrouter())
	// v2
	AddRoutes(muxRouter.PathPrefix("/v2").Subrouter())
	err := http.ListenAndServe(CONNECTION_HOST+":"+CONNECTION_PORT, router)
	
	if err != nil {
		log.Fatal("error starting http server :: ", err)
		return
	}
}