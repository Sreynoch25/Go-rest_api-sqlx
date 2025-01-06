package notificationz_controller

import (
	notification_model "marketing/src/models/notification"
	notifications_service "marketing/src/services/notifications"
	"marketing/src/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type NotificationController struct {
	notificationService notifications_service.NotificationService
}

func NewNotificationController(service *notifications_service.NotificationService) *NotificationController {
	return &NotificationController{
		notificationService: *service, 
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


func (c *NotificationController) Show(ctx *fiber.Ctx) error {
	page, err := strconv.Atoi(ctx.Query("page", "1"))
	if err != nil {
		page = 1
	}

	pageSize, err := strconv.Atoi(ctx.Query("pageSize", "10"))
	if err != nil {
		pageSize = 10
	}

	if page < 1 {
		page = 1
	}

	if pageSize < 1 {
		pageSize = 10
	}

	response, err := c.notificationService.Show(page, pageSize)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			utils.ApiResponse(
				false,
				"Failed to retrieve notifications",
				fiber.StatusInternalServerError,
				err.Error(),
			),
		)
	}

	return ctx.Status(fiber.StatusOK).JSON(
		utils.ApiResponse(
			true,
			"Notifications retrieved successfully",
			fiber.StatusOK,
			response,
		),
	)
}
