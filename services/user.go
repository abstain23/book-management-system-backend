package services

import (
	"book_management/constants"
	"book_management/db"
	"book_management/dto"
	"book_management/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Register(ctx *fiber.Ctx) error {
	user := &dto.User{}
	if err := ctx.BodyParser(user); err != nil {
		return &fiber.Error{Code: 400, Message: err.Error()}
	}

	if user.Username == "" {
		return &fiber.Error{Code: 400, Message: "用户名不能为空"}
	}

	if user.Password == "" {
		return &fiber.Error{Code: 400, Message: "密码不能为空"}
	}

	if len(user.Password) < 6 {
		return &fiber.Error{Code: 400, Message: "密码不能少于6位"}
	}

	users := db.GetUserList()
	for _, u := range users {
		if u.Username == user.Username {
			return &fiber.Error{Code: 400, Message: "用户名已存在"}
		}
	}

	if err := db.AddUser(*user); err != nil {
		return &fiber.Error{Code: 400, Message: err.Error()}
	}

	return ctx.JSON(utils.SuccessResponse(nil))
}

func Login(ctx *fiber.Ctx) error {
	users := db.GetUserList()
	loginUser := &dto.User{}
	if err := ctx.BodyParser(loginUser); err != nil {
		return &fiber.Error{Code: 400, Message: err.Error()}
	}
	for _, u := range users {
		if u.Username == loginUser.Username {
			if u.Password == loginUser.Password {
				claims := jwt.MapClaims{
					"username": u.Username,
					"exp": time.Now().Add(time.Hour * 24).Unix(),
				}

				token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
				tokenString, err := token.SignedString([]byte(constants.JWT_SECRET))
				if err != nil {
					return &fiber.Error{Code: fiber.StatusInternalServerError, Message: err.Error()}
				}

				return ctx.JSON(utils.SuccessResponse(tokenString))
			} else {
				return &fiber.Error{Code: 400, Message: "密码错误"}
			}
		}
	}

	return &fiber.Error{Code: 400, Message: "用户不存在"}
}
