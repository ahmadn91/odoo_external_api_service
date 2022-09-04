package main

import (
	"log"
	"os"

	"github.com/ahmadn91/odoo_external_api_service/config"
	"github.com/ahmadn91/odoo_external_api_service/handlers"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
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
		
	)
	boil.SetDB(config.Database)
	boil.DebugMode = true


	app.Get("/resPartner", handlers.GetContacts)
	log.Printf("server started")
	log.Fatal(app.Listen(":3000"))
}