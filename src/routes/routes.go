package routes

import (
	auth_routes "marketing/src/routes/auth"
	notifications_routes "marketing/src/routes/notifications"
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

	prefix := app.Group("/api/v1") //Create Api V1 route group\
	auth_routes.Auth(prefix, db)

	user_router.UserRoutes(prefix, db) //Init user routes under the Api/V1
	notifications_routes.Notification(prefix, db)
}
