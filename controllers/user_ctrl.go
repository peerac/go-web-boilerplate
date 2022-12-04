package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func UserController(c *fiber.Ctx) error {
	return c.Render("users/user", fiber.Map{})
}
