package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/virhanali/user-management/handler"
)

func Routes(app fiber.Router, userHandler *handler.UserHandler) {
	r := app.Group("/api/v1")
	r.Post("/", userHandler.Create)
}
