package auth_routes

import (
	auth_controller "marketing/src/controllers/Api/V1/auth"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func Auth(api fiber.Router, db *sqlx.DB) {
	// appPort := os.Getenv("JWT_SECRET")
	jwtSECRET := os.Getenv("JWT_SECRET")
	auth := auth_controller.NewAuthController(db, jwtSECRET)

	authGroup  := api.Group("/auth")

	authGroup.Post("/", auth.Login)

}