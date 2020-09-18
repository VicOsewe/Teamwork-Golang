package main

import (
	"Teamwork-Golang/api"

	"github.com/gorilla/mux"
)

func initializeRoutes(route *mux.Router, s services) {

	route.HandleFunc("/auth/user", api.CreateUser(s.creating)).Methods("POST")
	route.HandleFunc("/auth/signin", api.UserSignIn(s.creating)).Methods("POST")
	route.HandleFunc("/articles", api.CreateArticle(s.creating)).Methods("POST")
	route.HandleFunc("/articles", api.UpdateArticle(s.updating))

}
