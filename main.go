package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/virhanali/user-management/config"
)

func main() {
	config.Database()

	app := fiber.New()

	app.Listen(":3000")
}
