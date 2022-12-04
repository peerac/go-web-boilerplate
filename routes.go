package main

import (
	"jtik-pnl/pepeta/controllers"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Get("/", controllers.RootController)
	app.Get("/users", controllers.UserController)
}
