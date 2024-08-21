package routes

import (
	"book_management/dto"

	"github.com/gofiber/fiber/v2"
)

func initUserRouter(router fiber.Router) {
	userGroup := router.Group("/user")
	userGroup.Post("/register", register)
}

func register(ctx *fiber.Ctx) error{
	user := &dto.User{}
	if err := ctx.BodyParser(user); err != nil {
		return err
	}
	return ctx.JSON(fiber.Map{"code": 0, "message": "success", "data": user.Username + " register success"})
}
