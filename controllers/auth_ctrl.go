package controllers

import (
	"fmt"

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
	emailAddr := c.FormValue("emailAddr")
	password := c.FormValue("password")

	fmt.Println("FORM VALUE: ", emailAddr, password)

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
