package main

import (
	"log"
	"net/http"

	"github.com/TheBusBoys/Tracker/pkg/routes"
	"github.com/gorilla/mux"
)

func main(){
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Print("Server is now running on: http://localhost:9010/")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}