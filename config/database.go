package config

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
)

var Database *sql.DB

var DATABASE_URI string = "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable"

func Connect(host, user, password, dbName, port string) error {
	var err error
		connectionString := fmt.Sprintf(DATABASE_URI, host, user, password, dbName, port)


	Database, err = sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err)
	}

	
	return nil
}