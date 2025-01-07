package notifications_routes

import (
	notification_controller "marketing/src/controllers/Api/V1/notification"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Notification(api fiber.Router, db *sqlx.DB) {
	nc := notification_controller.NewNotificationController(db)
	notifications := api.Group("/notifications")

	notifications.Post("/", nc.Create)
	notifications.Put("/:id", nc.Update)
	notifications.Get("/", nc.Show)
	notifications.Delete("/:id", nc.Delete)
}
