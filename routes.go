package main

import (
	"jtik-pnl/pepeta/controllers"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	// index route
	app.Get("/", controllers.HomeController)

	// Auth routes
	app.Get("/login", controllers.LoginViewController)
	app.Post("/login", controllers.LoginHandler)
	app.Get("/register", controllers.RegisterViewController)
	app.Post("/register", controllers.RegisterHandler)

}
