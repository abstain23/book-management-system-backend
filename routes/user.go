package routes

import (
	"book_management/services"

	"github.com/gofiber/fiber/v2"
)

func initUserRouter(router fiber.Router) {
	userGroup := router.Group("/user")
	userGroup.Post("/register", register)
	userGroup.Post("/login", login)
}

func register(ctx *fiber.Ctx) error {
	return services.Register(ctx)
}

func login(ctx *fiber.Ctx) error {
	return services.Login(ctx)
}
