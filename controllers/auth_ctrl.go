package controllers

import (
	"fmt"
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
	Business Controller
*/
func LoginHandler(c *fiber.Ctx) error {
	return nil
}

func RegisterHandler(c *fiber.Ctx) error {
	fullName := c.FormValue("fullName")
	emailAddr := c.FormValue("emailAddr")
	password := c.FormValue("password")
	confirmPass := c.FormValue("confirmPassword")

	fmt.Println("FORM VALUE: ", fullName, emailAddr, password, confirmPass)

	return nil
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
