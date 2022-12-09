package main

import (
	"jtik-pnl/pepeta/controllers"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	// index route
	app.Get("/", controllers.RootController)

	// Auth routes
	app.Get("/login", controllers.LoginController)
	app.Get("/register", controllers.RegisterController)

	// Home routes
	app.Get("/dashboard", controllers.HomeController)

	// Rest API
	api := app.Group("api")
	api.Get("/dashboard", controllers.HomeAPIController)
	api.Get("/login", controllers.LoginAPIController)
	api.Get("/register", controllers.RegisterAPIController)

}
