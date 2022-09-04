package main

import (
	"log"
	"os"

	"github.com/Integrated-Path/sas4_external_api/config"
	"github.com/Integrated-Path/sas4_external_api/handlers"
	"github.com/gofiber/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/goccy/go-json"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func main() {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	log.Printf("Connect to the database")
	config.Connect(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)
	boil.SetDB(config.Database)
	boil.DebugMode = true


	app.Get("/resPartner", handlers.GetContact)
	log.Printf("server started")
	log.Fatal(app.Listen(":3000"))
}