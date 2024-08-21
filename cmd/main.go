package main

import (
	"book_management/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	routes.Init(app)

	app.Listen(":3000")
}
