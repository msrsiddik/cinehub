package main

import (
	"entities-module/database"
	"graphql-module/server"
	"log"
	"os"
	"restapi-module/restapi"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize the database connection
	db, q := database.DbInit()
	//database.GenModelQuery(db)

	// Set the port for the server, defaulting to 3000 if not specified in environment variables
	port := "3000" // Default port
	if envPort := os.Getenv("SERVER_PORT"); envPort != "" {
		port = envPort
	}

	app := fiber.New()

	// Set up a health check endpoint
	app.Get("/healthz", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	server.GraphServer(app, db, q)
	restapi.RestApiServer(app)

	log.Fatal(app.Listen(":" + port))
}
