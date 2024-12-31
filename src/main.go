package main

import (
	"log"
	"os"

	config "marketing/src/configs"
	"marketing/src/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func init() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, proceeding without it")
	}
}

func main() {
	// Database connection
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Fiber app setup
	app := fiber.New()

	// Setup routes
	routes.SetupRoutes(app, db)

	// Start server
	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "3000" // Default fallback
	}
	log.Fatal(app.Listen(":" + appPort))
	// log.Fatal(app.Listen(":3000"))
}
