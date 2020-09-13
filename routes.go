package main

import (
	"Teamwork-Golang/api"

	"github.com/gorilla/mux"
)

func initializeRoutes(route *mux.Router, s services) {

	route.HandleFunc("/teamwork/user", api.CreateUser(s.registering)).Methods("POST")

}
