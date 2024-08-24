package middleware

import (
	"book_management/constants"

	fiberJwt "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)


func JWtMiddleware() func(ctx *fiber.Ctx) error {
	return fiberJwt.New(fiberJwt.Config{
		SigningKey: fiberJwt.SigningKey{Key: []byte(constants.JWT_SECRET)},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return &fiber.Error{Code: fiber.StatusUnauthorized, Message: "Unauthorized"}
		},
		Filter: func(c *fiber.Ctx) bool {
			for _, route := range constants.PUBLIC_ROUTES {
				if route == c.Path() {
					return true
				}
			}
			return false
		},
	})
}
