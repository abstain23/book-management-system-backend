package main

import (
	"book_management/routes"
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := c.Response().StatusCode()
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}
			return c.Status(200).JSON(fiber.Map{"msg": err.Error(), "status": code, "code": 1})
		},
	})

	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Asia/Shanghai",
	}))

	app.Use(cors.New())

	app.Static("/static", "storage")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	routes.Init(app)

	log.Fatal(app.Listen(":3000"))
}
