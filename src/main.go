package main

import (
	"log"
	"os"

	config "marketing/src/configs"
	"marketing/src/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)
/*
 *Author: Noch
 *init this called before main, sets up the environment variables
*/
// func init() {
// 	
// 	if err := godotenv.Load(); err != nil {
// 		log.Println("No .env file found, proceeding without it")
// 	}
// }


/*
 *Author: Noch
 *main is the application entry point
*/
func main() {

	// Load environment variables from .env file , if no .env file exists
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, proceeding without it")
	}
	
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
		appPort = "3000" // Default port
	}
	log.Fatal(app.Listen(":" + appPort), err)
	// log.Fatal(app.Listen(":3000"))
}
