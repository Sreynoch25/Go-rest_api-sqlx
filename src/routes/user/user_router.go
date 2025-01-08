package user_router

import (
	user_controller "marketing/src/controllers/Api/V1/user"
	"marketing/src/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

/*
 *Author: Noch
 *UserRoutes sets up all user-related routes
 *Params:
 *	-  api: The router group to attach routes
 *	-  db: database connection
 */
func UserRoutes(api fiber.Router, db *sqlx.DB) {
	uc := user_controller.NewUserController(db)          // Initialize a new user controller
	userRouter := api.Group("/users")                    // Create a new sub-router for user routes
	userRouter.Get("/", middleware.Protected(), uc.Show) // GET /api/v1/users
	userRouter.Get("/:id", uc.ShowOne)                   // GET /api/v1/users/:id
	userRouter.Post("/", uc.Create)                      // POST /api/v1/users
	userRouter.Put("/:id", uc.Update)                    // PUT /api/v1/users/:id
	userRouter.Delete("/:id", uc.Delete)                 // DELETE /api/v1/users/:id
}
