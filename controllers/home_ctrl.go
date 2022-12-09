package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func HomeController(c *fiber.Ctx) error {
	return c.Render("home/dashboard", fiber.Map{})
}

/*
	API Controller
*/

func HomeAPIController(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"data": "dashboard page",
	})
}
