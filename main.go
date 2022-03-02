package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/johandui/domain-automation/controller"
	"github.com/johandui/domain-automation/utils"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: controller.Error,
	})
	utils.LoadEnv()
	app.Use(recover.New())
	app.Post("/", controller.Create)
	app.Listen(":3000")
}
