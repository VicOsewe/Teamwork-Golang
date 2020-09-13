package main

import (
	"log"
	"net/http"

	"Teamwork-Golang/data"
	"Teamwork-Golang/registering"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	db, err := gorm.Open("postgres", "host=db port=5432 user=postgres dbname=postgres sslmode=disable password=postgres")

	if err != nil {

		panic("failed to connect database")

	}

	defer db.Close()
	route := mux.NewRouter()
	services := initializeServices(db)
	initializeRoutes(route, services)
	log.Fatal(http.ListenAndServe(":2000", route).Error(), nil)

}

type services struct {
	registering registering.RegisterService
}

func initializeServices(db *gorm.DB) services {
	dbrepo := data.NewUserRepository(db)
	s := services{}
	s.registering = registering.NewRegisteringService(dbrepo)
	return s
}
