package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func RootController(c *fiber.Ctx) error {
	return c.Render("index/index", fiber.Map{})
}
