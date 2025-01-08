package auth_controller

import (
	"fmt"
	auth_model "marketing/src/models/auth"
	auth_repository "marketing/src/repositeries/auth"
	auth_service "marketing/src/services/auth"
	"marketing/src/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type AuthController struct {
	authService auth_service.AuthService
}

// func NewAuthControllerb (*sqlx.DB) *AuthController{
// 	repo := auth_service.NewAuthService(db, jwtSecret)
// 	authService := auth_service.NewAuthService(repo)
// 	return &AuthController{
// 		authService: authService,
// 	}
// }

func NewAuthController(db *sqlx.DB, jwtSecret string) *AuthController {
	return &AuthController{
		authService: auth_service.NewAuthService(auth_repository.NewAuthRepository(db), jwtSecret),
	}
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	loginRequest := auth_model.UserLogin{}
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
	fmt.Println("ddd", err)
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
