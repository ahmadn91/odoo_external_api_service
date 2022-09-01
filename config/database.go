package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

var DATABASE_URI string = "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable"

func Connect(host, user, password, dbName string) error {
	var err error
		connectionString := fmt.Sprintf(DATABASE_URI, host, user, password, dbname, port)


	Database, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	Database.AutoMigrate(&entities.contact{})
	return nil
}