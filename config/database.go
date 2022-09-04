package config

import (
	"fmt"
	"log"

	"github.com/ahmadn91/odoo_external_api_service/entities"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *boil.Executor

var DATABASE_URI string = "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable"

func Connect(host, user, password, dbName, port string) error {
	var err error
		connectionString := fmt.Sprintf(DATABASE_URI, host, user, password, dbName, port)


	Database, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	Database.AutoMigrate(&entities.Contact{})
	return nil
}