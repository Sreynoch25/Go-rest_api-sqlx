package role_router

import (
	role_controller "marketing/src/controllers/Api/V1/role"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

/*
 * Author: Noch
 * RoleRoutes sets up all role-related routes
 * Params:
 *   - api: The router group to attach routes to
 *   - db: Database connection
 */
func RoleRoutes(api fiber.Router, db *sqlx.DB) {
	rc := role_controller.NewRoleController(db) // Initialize a new role controller
	roleRouter := api.Group("/roles") // Create a new sub-router for role routes

	roleRouter.Get("/", rc.Show) // GET /api/v1/roles
	roleRouter.Get("/:id", rc.ShowOne) // GET /api/v1/roles/:id
	roleRouter.Post("/", rc.Create) // POST /api/v1/roles
	roleRouter.Put("/:id", rc.Update) // PUT /api/v1/roles/:id
	roleRouter.Delete("/:id", rc.Delete) // DELETE /api/v1/roles/:id
}


