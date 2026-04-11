package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() (*gorm.DB, error) {
	dataSrcName := "host=db user=postgres password=devpassword dbname=database port=5432 sslmode=disable"

	var err error

	db, err = gorm.Open(postgres.Open(dataSrcName), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	return db, nil
}
