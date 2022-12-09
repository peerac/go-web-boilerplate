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
	app.Get("/api", controllers.HomeAPIController)

}
