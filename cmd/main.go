package main

import (
	"drivehino-trivia-restapi/database"

	"github.com/gofiber/fiber/v2"
)

func main()  {
	database.ConnectDb()

	app := fiber.New()

	setupRoutes(app)

	// move below to routes.go
	// app.Get("/", func (c *fiber.Ctx) error {
	// 	return c.SendString("Hello Alfin Umi Kulsum")
	// })

	app.Listen(":3000")
}