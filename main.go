package main

import (
	"github.com/Arismonx/task-manager/config"
	"github.com/Arismonx/task-manager/routers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	//func LoadENV from .env file
	config.LoadENV()

	//Connecting Database
	config.ConnectDB()

	app := fiber.New()

	//Test path '/'
	routers.SetAPI(app)
	app.Listen(":8080")
}
