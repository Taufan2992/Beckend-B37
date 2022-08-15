package main

import (
	"fmt"
	"golang/database"
	"golang/pkg/mysql"
	"golang/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// initial DB
	mysql.DatabaseInit()

	// run migration
	database.RunMigration()

	r := mux.NewRouter()

	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	//path file
	r.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads")))) // add this code

	fmt.Println("server running localhost:5000")
	http.ListenAndServe("localhost:5000", r)
}
