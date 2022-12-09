package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func LoginController(c *fiber.Ctx) error {
	return c.Render("auth/login", fiber.Map{})
}

func RegisterController(c *fiber.Ctx) error {
	return c.Render("auth/register", fiber.Map{})
}

/*
	API Controller
*/

func LoginAPIController(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"data": "Login API route",
	})
}

func RegisterAPIController(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"data": " Register API route",
	})
}
