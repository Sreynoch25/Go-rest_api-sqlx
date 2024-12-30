package user_router

import (
	user_controller "marketing/src/controllers/user"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

// UserRoutes function
func UserRoutes(api fiber.Router, db *sqlx.DB) {
	uc := user_controller.NewUserController(db)  // Initialize a new user controller
	userRouter := api.Group("/users") // Create a new sub-router for user routes

	userRouter.Get("/", uc.ShowAllUser) // Fetch all users
	userRouter.Get("/:id", uc.ShowUserByID) // Fetch a user by ID
}
