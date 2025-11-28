package routers

import (
	"simple-api-with-go-fiber/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupRouters(app *fiber.App) {
	app.Get("/", controller.Homepage)
	app.Post("/create", controller.CreateUser)
	app.Get("/apiusers", controller.GetUsers)
}
