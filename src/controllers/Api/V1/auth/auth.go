package auth_controller

import (
	auth_model "marketing/src/models/auth"
	auth_service "marketing/src/services/auth"
	"marketing/src/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type AuthController struct {
	authService auth_service.AuthService
}

func NewAuthController(db *sqlx.DB) *AuthController {
	return &AuthController{
		// 
		authService:  auth_service.NewAuthService(db),
	}
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	loginRequest := auth_model.LoginRequest{}
	if err := ctx.BodyParser(&loginRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			utils.ApiResponse(
				false,
				"Invalid request Body",
				fiber.StatusBadRequest,
				err,
			),
		)
	}

	response, err := c.authService.UserLogin(loginRequest)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(
			utils.ApiResponse(
				false,
				"Invalid credentials",
				fiber.StatusUnauthorized,
				err,
			), 

			
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(
		utils.ApiResponse(
			true,
			"Login successful",
			fiber.StatusOK,
			fiber.Map{
				"token": response.Token,
				
			},
		),
	)
}
