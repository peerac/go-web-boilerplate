package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func LoginController(c *fiber.Ctx) error {
	return c.Render("auth/login", fiber.Map{})
}

func RegisterController(c *fiber.Ctx) error {
	return c.Render("auth/register", fiber.Map{})
}
