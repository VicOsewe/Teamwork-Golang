package main

import (
	"Teamwork-Golang/deleting"
	"Teamwork-Golang/updating"
	"log"
	"net/http"

	"Teamwork-Golang/creating"
	"Teamwork-Golang/data"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=db sslmode=disable password=password")

	if err != nil {
		log.Fatal(err.Error(), nil)
	}

	defer db.Close()
	route := mux.NewRouter()
	services := initializeServices(db)
	initializeRoutes(route, services)
	log.Fatal(http.ListenAndServe(":2000", route).Error(), nil)

}

type services struct {
	creating creating.CreatingService
	updating updating.UpdateService
	deleting deleting.DeleteService
	// getting     getting.GettingService
}

func initializeServices(db *gorm.DB) services {
	dbrepo := data.NewUserRepository(db)
	s := services{}
	s.creating = creating.NewcreatingService(dbrepo)
	s.updating = updating.NewUpdatingService(dbrepo)
	s.deleting = deleting.NewDeletingService(dbrepo)
	return s
}
