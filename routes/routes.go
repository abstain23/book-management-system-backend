package routes

import "github.com/gofiber/fiber/v2"

func Init(app *fiber.App) {
	 v1 := app.Group("/api")
	initUserRouter(v1)
}
