package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/virhanali/user-management/domain/models"
	"github.com/virhanali/user-management/domain/response"
	"github.com/virhanali/user-management/services"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{userService: service}
}

func (handler UserHandler) Create(ctx *fiber.Ctx) error {
	userRequest := models.CreateUserRequest{}
	if err := ctx.BodyParser(&userRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}
	newUser, err := handler.userService.Store(userRequest)

	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Create user failed",
			Error:   err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(response.SuccessResponse{
		Code:    fiber.StatusOK,
		Message: "User created successfully",
		Data:    newUser,
	})
}