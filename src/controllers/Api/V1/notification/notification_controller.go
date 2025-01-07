package notification_controller

import (
	notification_model "marketing/src/models/notification"
	notification_repository "marketing/src/repositeries/notification"
	notifications_service "marketing/src/services/notifications"
	"marketing/src/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type NotificationController struct {
	notificationService notifications_service.NotificationService
}

// func NewNotificationController(service *notifications_service.NotificationService) *NotificationController {
// 	return &NotificationController{
// 		notificationService: *service, 
// 	}
// }
func NewNotificationController(db *sqlx.DB) *NotificationController {

	repo := notification_repository.NewNotificationRepository(db)
	service := notifications_service.NewNotificationService(repo)

	return &NotificationController{
		notificationService: service, 
	}
}

func (c *NotificationController) Create(ctx *fiber.Ctx) error {
	req := new(notification_model.CreateNotificationRequest)

	// Parse the body of the request
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			utils.ApiResponse(
			false,
			"Failed to parse request body",
			fiber.StatusBadRequest,
			err.Error(),
		))
	}

	// Call the service to create the notification
	response, err := c.notificationService.Create(req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			utils.ApiResponse(
			false,
			"Failed to create notification",
			fiber.StatusInternalServerError,
			err.Error(),
		))
	}

	// Return success response
	return ctx.Status(fiber.StatusCreated).JSON(
		utils.ApiResponse(
		true, 
		"Notification created successfully",
		fiber.StatusCreated,
		response,
	))
}

func (c *NotificationController) Update(ctx *fiber.Ctx) error {

	id , err := strconv.Atoi(ctx.Params("id"))
	if  err != nil  {
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			utils.ApiResponse(
				false,
				"Invalid request body",
				fiber.StatusInternalServerError,
				err,
			),
		)
	}

	req := new(notification_model.UpdateNotificationRequest)
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			utils.ApiResponse(
				false,
				"Failed to parse requese body",
				fiber.StatusBadRequest,
				err,
			),
		)
	}

	response, err := c.notificationService.Update(id, req)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			utils.ApiResponse(
				false,
				"Failed to update notification",
				fiber. StatusInternalServerError,
				err.Error(),
			),
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(
		utils.ApiResponse(
			true,
			"Notification updated successfully",
			fiber.StatusOK,
			response,
		),
	)
}


// func (c *NotificationController) Show(ctx *fiber.Ctx) error {
//     // Get pagination parameters with validation
//     page := ctx.QueryInt("page", 1)
//     perPage := ctx.QueryInt("per_page", 10)


//     // Get notifications with pagination
//     response, err := c.notificationService.Show(page, perPage)
//     if err != nil {
//         return ctx.Status(fiber.StatusInternalServerError).JSON(
//             utils.ApiResponse(
//                 false,
//                 "Failed to retrieve notifications",
//                 fiber.StatusInternalServerError,
//                 err.Error(),
//             ),
//         )
//     }

//     // Return successful response with pagination
//     return ctx.Status(fiber.StatusOK).JSON(
//         utils.ApiResponseWithPagination(
//             true,
//             "Notifications retrieved successfully",
//             fiber.StatusOK,
//             response.Notifications,
//             page,
//             perPage,
//             response.Total,
//         ),
//     )
// }
func (c *NotificationController) Show(ctx *fiber.Ctx) error {
	// Get pagination parameters
	page := ctx.QueryInt("page", 1)
	perPage := ctx.QueryInt("per_page", 2) //ctx.QueryInt is used to get the query params from url

	// Get users with pagination
	response, err := c.notificationService.Show(page, perPage)
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

func (c *NotificationController) Delete(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			utils.ApiResponse(
				false,
				"Invalid ID",
				fiber.StatusBadRequest,
				err.Error(),
			),
		)
	}

	deletedBy := 1 // Replace with dynamic user ID if available
	if err := c.notificationService.Delete(id, deletedBy); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			utils.ApiResponse(
				false,
				"Failed to delete notification",
				fiber.StatusInternalServerError,
				err.Error(),
			),
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(
		utils.ApiResponse(
			true,
			"Notification deleted successfully",
			fiber.StatusOK,
			nil,
		),
	)
}
