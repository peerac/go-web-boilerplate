package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

/*
	View Controller
*/

func LoginViewController(c *fiber.Ctx) error {
	return c.Render("auth/login", fiber.Map{})
}

func RegisterViewController(c *fiber.Ctx) error {
	return c.Render("auth/register", fiber.Map{})
}

/*
	API Controller
*/

func LoginAPIController(c *fiber.Ctx) error {
	// Get user form value
	// Check username password
	// If correct, generate JWT, and save login cookies, and redirect to dashboard.
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"data": "Login API route",
	})
}

func RegisterAPIController(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"data": " Register API route",
	})
}

/*
	Common Controller
*/
