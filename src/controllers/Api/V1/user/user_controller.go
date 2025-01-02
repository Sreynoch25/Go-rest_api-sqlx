// controllers/user/user_controller.go
package user_controller

import (
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"

	user_model "marketing/src/models/user"
	user_repository "marketing/src/repositeries/user"
	user_service "marketing/src/services/user"
	"marketing/src/utils"
)

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

func (c *UserController) Create(ctx *fiber.Ctx) error {
    userReq := new(user_model.UserRequest)
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

    user, err := c.userService.Create(userReq)
    if err != nil {
        if validationErr, ok := err.(validator.ValidationErrors); ok {
            return ctx.Status(fiber.StatusBadRequest).JSON(
                utils.ApiResponse(
                    false,
                    "Validation failed",
                    fiber.StatusBadRequest,
                    formatValidationErrors(validationErr),
                ),
            )
        }
        
        return ctx.Status(fiber.StatusInternalServerError).JSON(
            utils.ApiResponse(
                false,
                "Failed to create user",
                fiber.StatusInternalServerError,
                err.Error(),
            ),
        )
    }

    return ctx.Status(fiber.StatusCreated).JSON(
        utils.ApiResponse(
            true,
            "User created successfully",
            fiber.StatusCreated,
            user,
        ),
    )
}

func (c *UserController) Update(ctx *fiber.Ctx) error {
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

    userReq := new(user_model.UserRequest)
    if err := ctx.BodyParser(userReq); err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(
            utils.ApiResponse(
                false,
                "Invalid request body",
                fiber.StatusBadRequest,
                err,
            ),
        )
    }

    if err := c.userService.Update(id, userReq); err != nil {
        if validationErr, ok := err.(validator.ValidationErrors); ok {
            return ctx.Status(fiber.StatusBadRequest).JSON(
                utils.ApiResponse(
                    false,
                    "Validation failed",
                    fiber.StatusBadRequest,
                    // formatValidationErrors(validationErr),
					formatValidationErrors(validationErr), //formatValidationErrors is a function that formats the validation errors
                ),
            )
        }

        return ctx.Status(fiber.StatusInternalServerError).JSON(
            utils.ApiResponse(
                false,
                "Failed to update user",
                fiber.StatusInternalServerError,
                err,
            ),
        )
    }

    return ctx.Status(fiber.StatusOK).JSON(
        utils.ApiResponse(
            true,
            "User updated successfully",
            fiber.StatusOK,
            nil,
        ),
    )
}

func formatValidationErrors(validationErr validator.ValidationErrors) interface{} {
	panic("unimplemented")
}

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

func (c *UserController) ShowOne(ctx *fiber.Ctx) error {
	// Parse ID parameter
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

	// Handle not found
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

	// Return successful response
	return ctx.Status(fiber.StatusOK).JSON(
		utils.ApiResponse(
			true,
			"User retrieved successfully",
			fiber.StatusOK,
			user,
		),
	)
}

func (c *UserController) Delete(ctx *fiber.Ctx) error {
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

    deletedBy := ctx.QueryInt("deleted_by", 0)
    if err := c.userService.Delete(id, deletedBy); err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(
            utils.ApiResponse(
                false,
                "Failed to delete user",
                fiber.StatusInternalServerError,
                err,
            ),
        )
    }

    return ctx.Status(fiber.StatusOK).JSON(
        utils.ApiResponse(
            true,
            "User deleted successfully",
            fiber.StatusOK,
            nil,
        ),
    )
}