package main

import (
	"Teamwork-Golang/api"

	"github.com/gorilla/mux"
)

func initializeRoutes(route *mux.Router, s services) {

	route.HandleFunc("/auth/user", api.CreateUser(s.registering)).Methods("POST")
	route.HandleFunc("/auth/signin", api.UserSignIn(s.registering)).Methods("POST")

}
