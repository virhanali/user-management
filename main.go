package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/virhanali/user-management/config"
	"github.com/virhanali/user-management/handler"
	"github.com/virhanali/user-management/repository"
	"github.com/virhanali/user-management/routes"
	"github.com/virhanali/user-management/services"
)

func main() {
	config.Database()

	app := fiber.New()

	userRepository := repository.NewUserRepository(config.DB)
	userService := services.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	routes.Routes(app, userHandler)

	app.Listen(":3000")
}
