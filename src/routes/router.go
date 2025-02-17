package routes

import (
	"marketing/src/middleware"
	auth_routes "marketing/src/routes/auth"
	notifications_routes "marketing/src/routes/notifications"
	"marketing/src/routes/role_router.go"
	user_router "marketing/src/routes/user"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

/*
 *Author: Noch
 *SetupRoutes init all application routes
 *Params:
 *	- app: The fiber application instance
 *	- db: The database connection instance
 */

 func SetupRoutes(app *fiber.App, db *sqlx.DB) {
    prefix := app.Group("/api/v1")  //Create Api V1 route group
	auth_routes.Auth(prefix, db)
	// Use the protected middleware with configuration
	prefix.Use(middleware.JwtMiddleware())
	// prefix.Use(middleware.Protected())
   

    user_router.UserRoutes(prefix, db)
    role_router.RoleRoutes(prefix, db)
    notifications_routes.Notification(prefix, db)
}

