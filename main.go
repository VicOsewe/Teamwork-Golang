package main

import (
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	db, err := gorm.Open("postgres", "host=db port=5432 user=postgres dbname=postgres sslmode=disable password=postgres")

	if err != nil {

		panic("failed to connect database")

	}

	defer db.Close()

}
