package main

import (
	"github.com/Arismonx/task-manager/config"
	"github.com/gofiber/fiber/v2"
)

func main() {
	//func LoadENV from .env file
	config.LoadENV()

	//Connecting Database
	config.ConnectDB()

	app := fiber.New()

	//Test path '/'
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})
	app.Listen(":8080")
}
