package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/virhanali/user-management/handler"
)

func Routes(app fiber.Router, userHandler *handler.UserHandler) {
	r := app.Group("/api/v1")
	r.Get("/users", userHandler.GetAllUser)
	r.Get("/users/:id", userHandler.GetUserByID)

	r.Get("/admin", userHandler.GetAllUser)
	r.Get("/admin/:id", userHandler.GetUserByID)
	r.Post("/admin", userHandler.CreateUser)
	r.Put("/admin/:id", userHandler.UpdateUser)
	r.Delete("/admin/:id", userHandler.DeleteUser)
}
