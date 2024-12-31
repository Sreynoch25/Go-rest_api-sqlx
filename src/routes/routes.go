package routes

import (
	user_router "marketing/src/routes/user"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

// SetupRoutes function is responsible for initializing and grouping
func SetupRoutes(app *fiber.App, db *sqlx.DB) {
	// User routes
	prefix := app.Group("/api/v1")
	
	user_router.UserRoutes(prefix, db)
}