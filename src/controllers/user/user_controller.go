package user_controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"

	user_repository "marketing/src/repositeries/user"
	user_service "marketing/src/services/user"
)

// UserFactory is an interface defining methods for user-related HTTP handlers
type UserFactory interface {
	ShowAllUser(ctx *fiber.Ctx) error //hadles fetching all users
	ShowUserByID(ctx *fiber.Ctx) error //handles fetching a user by ID
}

type UserController struct {
	userService user_service.UserService
}

/*
    *author:Noch
     * NewUserController initializes a new UserController instance
*/
func NewUserController(db *sqlx.DB) *UserController {
	repo := user_repository.NewUserRepository(db)
	service := user_service.NewUserService(repo)
	return &UserController{
		userService: service,
	}
}

/*
    *author:Noch
    * ShowAllUser is a handler function for fetching all users.
*/
func (c *UserController) ShowAllUser(ctx *fiber.Ctx) error {
	users, err := c.userService.GetAll() 
    // Respond with HTTP 500 (Internal Server Error) if retrieval fails
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{ // Return a 500 Internal Server Error
			"error": err.Error(),
		})
	}

      // Respond with the user data in JSON format
	return ctx.JSON(fiber.Map{
		"data": users, // Include the retrieved user data in the response
	})
}


/*
    *author:Noch
    *ShowUserByID is a handler function for fetching a single user by their ID.
*/
func (c *UserController) ShowUserByID(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id")) 	// Extract the user ID from the URL params
    // Respond with HTTP 400 (Bad Request) if the ID is invalid
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{ // Return a 400 Bad Request
			"error": "Invalid ID format",
		})
	}

    // Retrieve the user by ID from the service layer
	user, err := c.userService.GetByID(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(), // Return the error message
		})
	}

    // Respond with the user data in JSON format
	return ctx.JSON(fiber.Map{ 
		"data": user, // Include the retrieved user data in the response
	})
}
