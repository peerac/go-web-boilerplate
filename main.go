package main

import (
	"fmt"
	"log"
	"os"

	"jtik-pnl/pepeta/initilizers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
)

func init() {
	initilizers.LoadENV()
	initilizers.ConnectToDatabase()
	initilizers.SyncDatabase()
}

func main() {
	//load template
	engine := html.New("./views", ".html")

	// configure app
	app := fiber.New(fiber.Config{
		AppName: "jtik-pnl-pepeta",
		Views:   engine,
	})
	app.Use(logger.New())
	app.Static("/", "./public")

	//route
	Routes(app)

	//config
	PORT := os.Getenv("APP_PORT")

	err := app.Listen(fmt.Sprintf(":%v", PORT))
	if err != nil {
		log.Fatalf("SERVER:: failed starting the server -> err: %v", err)
	}
}
