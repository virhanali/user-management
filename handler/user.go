package handler

import (
	"strconv"

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

func (handler UserHandler) CreateUser(ctx *fiber.Ctx) error {
	userRequest := models.CreateUserRequest{}
	if err := ctx.BodyParser(&userRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}
	newUser, err := handler.userService.CreateUser(userRequest)

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

func (handler UserHandler) GetAllUser(ctx *fiber.Ctx) error {
	users, err := handler.userService.GetAllUser()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Get all user failed",
			Error:   err.Error(),
		})
	}

	if len(users) < 1 {
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
			Code:    fiber.StatusNotFound,
			Message: "No user found",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(response.SuccessResponse{
		Code:    fiber.StatusOK,
		Message: "Get all user successfully",
		Data:    users,
	})
}

func (handler UserHandler) GetUserByID(ctx *fiber.Ctx) error {
	userId, _ := strconv.Atoi(ctx.Params("id"))

	user, err := handler.userService.GetUserById(userId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Get user by id failed",
			Error:   err.Error(),
		})
	}
	if user.ID == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorResponse{
			Code:    fiber.StatusNotFound,
			Message: "No user found",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(response.SuccessResponse{
		Code:    fiber.StatusOK,
		Message: "Get user by id successfully",
		Data:    user,
	})
}

func (handler UserHandler) UpdateUser(ctx *fiber.Ctx) error {
	userId, _ := strconv.Atoi(ctx.Params("id"))
	userRequest := models.UpdateUserRequest{}
	if err := ctx.BodyParser(&userRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
	}
	user, err := handler.userService.UpdateUser(userRequest, userId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Update user failed",
			Error:   err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(response.SuccessResponse{
		Code:    fiber.StatusOK,
		Message: "Update user successfully",
		Data:    user,
	})
}

func (handler UserHandler) DeleteUser(ctx *fiber.Ctx) error {
	userId, _ := strconv.Atoi(ctx.Params("id"))
	err := handler.userService.DeleteUser(userId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Delete user failed",
			Error:   err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(response.SuccessResponse{
		Code:    fiber.StatusOK,
		Message: "Delete user successfully",
	})
}
