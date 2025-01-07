// controllers/user/user_controller.go
package user_controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"

	user_model "marketing/src/models/user"
	user_repository "marketing/src/repositeries/user"
	user_service "marketing/src/services/user"
	"marketing/src/utils"
)


type UserFactory interface {
    Create(ctx *fiber.Ctx) error
    Update(ctx *fiber.Ctx) error
    Show(ctx *fiber.Ctx) error
    ShowOne(ctx *fiber.Ctx) error
    Delete(ctx *fiber.Ctx) error
}

type UserController struct {
	userService user_service.UserService
}

func NewUserController(db *sqlx.DB) *UserController {
	repo := user_repository.NewUserRepository(db)
	service := user_service.NewUserService(repo)
	return &UserController{
		userService: service,
	}
}

/*
 * Author: Noch
 * Create handles POST requests to create a new user
 */
func (c *UserController) Create(ctx *fiber.Ctx) error {

    // Parse request body into CreateUserRequest struct
    userReq := new(user_model.CreateUserRequest)
    if err := ctx.BodyParser(userReq); err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(
            utils.ApiResponse(
                false,
                "Failed to parse request body",
                fiber.StatusBadRequest,
                err,
            ),
        )
    }

    // Calls the service layer to create the user
    user, err := c.userService.Create(userReq)
    if err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(
            utils.ApiResponse(
                false,
                "Failed to create user",
                fiber.StatusInternalServerError,
                err.Error(),
            ),
        )
    }

    //Return appropriate response with the status code
    return ctx.Status(fiber.StatusCreated).JSON(
        utils.ApiResponse(
            true,
            "User created successfully",
            fiber.StatusCreated,
            user,
        ),
    )
}

/*
 * Author: Noch
 *  Update handles PUT requests to update an existing user
 */
 func (c *UserController) Update(ctx *fiber.Ctx) error {
    // Parse user ID from URL parameters
    id, err := strconv.Atoi(ctx.Params("id"))
    if err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(
            utils.ApiResponse(
                false,
                "Invalid user ID",
                fiber.StatusBadRequest,
                err,
            ),
        )
    }

    // Parse request body into UpdateUserRequest struct
    userReq := new(user_model.UpdateUserRequest)
    if err := ctx.BodyParser(userReq); err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(
            utils.ApiResponse(
                false,
                "Failed to parse request body",
                fiber.StatusBadRequest,
                err,
            ),
        )
    }

    // Call service layer to update user
    updatedUser, err := c.userService.Update(id, userReq)
    if err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(
            utils.ApiResponse(
                false,
                "Failed to update user",
                fiber.StatusInternalServerError,
                err.Error(),
            ),
        )
    }

    // Return success response with updated user data
    return ctx.Status(fiber.StatusOK).JSON(
        utils.ApiResponse(
            true,
            "User updated successfully",
            fiber.StatusOK,
            updatedUser,
        ),
    )
}

/*
 * Author: Noch
 * Show handles GET requests to retrieve a paginated list of users
 */


func (c *UserController) Show(ctx *fiber.Ctx) error {
	// Get pagination parameters
	page := ctx.QueryInt("page", 1)
	perPage := ctx.QueryInt("per_page", 2) //ctx.QueryInt is used to get the query params from url

	// Get users with pagination
	response, err := c.userService.Show(page, perPage)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			utils.ApiResponse(
				false,
				"Failed to fetch users",
				fiber.StatusInternalServerError,
				err,
			),
		)
	}



	// Return successful response
	return ctx.Status(fiber.StatusOK).JSON(
		utils.ApiResponseWithPagination(
			true,
			"Users retrieved successfully",
			6000,
			response,
			page,
			perPage,
			response.Total,
		),
	)
}

/*
 * Author: Noch
 * ShowOne handles GET requests to retrieve a single user by ID
 */
func (c *UserController) ShowOne(ctx *fiber.Ctx) error {

	// Extracts user ID from URL parameters
	id, err := strconv.Atoi(ctx.Params("id")) //strconv is used to convert string to int
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			utils.ApiResponse(
				false,
				"Invalid user ID format",
				fiber.StatusBadRequest,
				err,
			),
		)
	}

	// Get user
	user, err := c.userService.ShowOne(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			utils.ApiResponse(
				false,
				"Failed to fetch user",
				fiber.StatusInternalServerError,
				err,
			),
		)
	}

	// Handle not found case
	if user == nil {
		return ctx.Status(fiber.StatusNotFound).JSON(
			utils.ApiResponse(
				false,
				"User not found",
				fiber.StatusNotFound,
				nil,
			),
		)
	}

	// Returns user data if found
	return ctx.Status(fiber.StatusOK).JSON(
		utils.ApiResponse(
			true,
			"User retrieved successfully",
			fiber.StatusOK,
			user,
		),
	)
}

/*
 * Author: Noch
 * Delete handles DELETE requests to soft delete a user
  * Required parameters:
 *	-  id: user ID (from URL)
 *	-   deleted_by: ID of user performing the deletion (from query string)
 */
func (c *UserController) Delete(ctx *fiber.Ctx) error {
    // Parse user ID from URL parameters
    id, err := strconv.Atoi(ctx.Params("id"))
    if err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(
            utils.ApiResponse(
                false,
                "Invalid user ID",
                fiber.StatusBadRequest,
                err,
            ),
        )
    }

    // Get deleted_by from query parameters
    deletedBy := ctx.QueryInt("deleted_by", 0)
    if deletedBy == 0 {
        return ctx.Status(fiber.StatusBadRequest).JSON(
            utils.ApiResponse(
                false,
                "deleted_by parameter is required",
                fiber.StatusBadRequest,
                nil,
            ),
        )
    }

    // Call service layer to delete user
    if err := c.userService.Delete(id, deletedBy); err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(
            utils.ApiResponse(
                false,
                "Failed to delete user",
                fiber.StatusInternalServerError,
                err.Error(),
            ),
        )
    }

    // Return success response
    return ctx.Status(fiber.StatusOK).JSON(
        utils.ApiResponse(
            true,
            "User deleted successfully",
            fiber.StatusOK,
            nil,
        ),
    )
}