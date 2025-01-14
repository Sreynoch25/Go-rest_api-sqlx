package role_controller

import (
	"fmt"
	role_model "marketing/src/models/role"
	role_service "marketing/src/services/roles"
	"marketing/src/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type RoleController interface {
	Show(ctx *fiber.Ctx) error
	ShowOne(ctx *fiber.Ctx) error
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type roleController struct {
	roleService role_service.RoleService
}

func NewRoleController(db *sqlx.DB) RoleController {
	return &roleController{
		roleService: role_service.NewRoleService(db),
	}
}

func (c *roleController) Show(ctx *fiber.Ctx) error {
    // Get pagination parameters with better defaults
    page := ctx.QueryInt("page", 1)
    perPage := ctx.QueryInt("per_page", 10) // Changed default to 10 for better usability

    // Validate pagination parameters
    if page < 1 {
        return ctx.Status(fiber.StatusBadRequest).JSON(
            utils.ApiResponse(
                false,
                "Page number must be greater than 0",
                fiber.StatusBadRequest,
                nil,
            ),
        )
    }

    if perPage < 1 {
        return ctx.Status(fiber.StatusBadRequest).JSON(
            utils.ApiResponse(
                false,
                "Items per page must be greater than 0",
                fiber.StatusBadRequest,
                nil,
            ),
        )
    }

    response, err := c.roleService.Show(page, perPage)
    if err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(
            utils.ApiResponse(
                false,
                "Failed to fetch roles",
                fiber.StatusInternalServerError,
                err,
            ),
        )
    }

    return ctx.Status(fiber.StatusOK).JSON(
        utils.ApiResponseWithPagination(
            true,
            "Roles retrieved successfully",
            fiber.StatusOK,
            response.Roles,  
            page,
            perPage,
            response.Total,  
        ),
    )
}

func (c *roleController) ShowOne(ctx *fiber.Ctx) error {
	// Extracts role ID from URL parameters
	id, err := strconv.Atoi(ctx.Params("id")) // strconv is used to convert string to int
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			utils.ApiResponse(
				false,
				"Invalid role ID format",
				fiber.StatusBadRequest,
				err,
			),
		)
	}

	// Get role
	role, err := c.roleService.ShowOne(id)
	fmt.Println(err)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			utils.ApiResponse(
				false,
				"Failed to fetch role",
				fiber.StatusInternalServerError,
				err,
			),
		)
	}
	

	// Handle not found case
	if role == nil {
		return ctx.Status(fiber.StatusNotFound).JSON(
			utils.ApiResponse(
				false,
				"Role not found",
				fiber.StatusNotFound,
				nil,
			),
		)
	}

	// Returns role data if found
	return ctx.Status(fiber.StatusOK).JSON(
		utils.ApiResponse(
			true,
			"Role retrieved successfully",
			fiber.StatusOK,
			role,
		),
	)
}

func (c *roleController) Create(ctx *fiber.Ctx) error {
	// Parse request body into CreateRoleRequest struct
	roleReq := new(role_model.CreateRoleRequest)
	if err := ctx.BodyParser(roleReq); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			utils.ApiResponse(
				false,
				"Failed to parse request body",
				fiber.StatusBadRequest,
				err,
			),
		)
	}

	// Calls the service layer to create the role
	role, err := c.roleService.Create(roleReq)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			utils.ApiResponse(
				false,
				"Failed to create role",
				fiber.StatusInternalServerError,
				err.Error(),
			),
		)
	}

	return ctx.Status(fiber.StatusCreated).JSON(
		utils.ApiResponse(
			true,
			"Role created successfully",
			fiber.StatusCreated,
			role,
		),
	)
}

func (c *roleController) Update(ctx *fiber.Ctx) error {
	// Parse role ID from URL parameters
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			utils.ApiResponse(
				false,
				"Invalid role ID",
				fiber.StatusBadRequest,
				err,
			),
		)
	}

	// Parse request body into UpdateRoleRequest struct
	roleReq := new(role_model.UpdateRoleRequest)
	if err := ctx.BodyParser(roleReq); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			utils.ApiResponse(
				false,
				"Failed to parse request body",
				fiber.StatusBadRequest,
				err,
			),
		)
	}

	// Call service layer to update role
	updatedRole, err := c.roleService.Update(id, roleReq)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			utils.ApiResponse(
				false,
				"Failed to update role",
				fiber.StatusInternalServerError,
				err.Error(),
			),
		)
	}

	// Return success response with updated role data
	return ctx.Status(fiber.StatusOK).JSON(
		utils.ApiResponse(
			true,
			"Role updated successfully",
			fiber.StatusOK,
			updatedRole,
		),
	)
}

func (c *roleController) Delete(ctx *fiber.Ctx) error {
 id, err := ctx.ParamsInt("id")
    if err != nil {
        return ctx.Status(fiber.StatusBadRequest).JSON(
            utils.ApiResponse(
                false,
                "Invalid role ID",
                fiber.StatusBadRequest,
                err,
            ),
        )
    }

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

    if err := c.roleService.Delete(id, deletedBy); err != nil {
        return ctx.Status(fiber.StatusInternalServerError).JSON(
            utils.ApiResponse(
                false,
                "Failed to delete role",
                fiber.StatusInternalServerError,
                err,
            ),
        )
    }

	// Return success response
	return ctx.Status(fiber.StatusOK).JSON(
		utils.ApiResponse(
			true,
			"Role deleted successfully",
			fiber.StatusOK,
			nil,
		),
	)
}
