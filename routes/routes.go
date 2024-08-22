package routes

import "github.com/gofiber/fiber/v2"

func Init(app *fiber.App) {
	api := app.Group("/api")
	initUserRouter(api)
	initBookRouter(api)
}
