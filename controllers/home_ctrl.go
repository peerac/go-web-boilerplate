package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func HomeController(c *fiber.Ctx) error {
	return c.Render("home/dashboard", fiber.Map{})
}

/*
	Common Function
*/
