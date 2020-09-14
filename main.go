package main

import (
	"Teamwork-Golang/getting"
	"log"
	"net/http"

	"Teamwork-Golang/data"
	"Teamwork-Golang/registering"

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
	registering registering.RegisterService
	getting     getting.GettingService
}

func initializeServices(db *gorm.DB) services {
	dbrepo := data.NewUserRepository(db)
	s := services{}
	s.registering = registering.NewRegisteringService(dbrepo)
	s.getting = getting.NewGettingService(dbrepo)
	return s
}
