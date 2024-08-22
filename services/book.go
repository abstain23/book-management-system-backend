package services

import (
	"book_management/db"
	"book_management/dto"
	"book_management/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetBookList(ctx *fiber.Ctx) error {
	books := db.GetBookList()
	return ctx.JSON(utils.SuccessResponse(books))
}

func GetBookById(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return &fiber.Error{Code: 400, Message: "ID不合法"}
	}
	book, exist := db.GetBook(id)
	if !exist {
		return &fiber.Error{Code: 400, Message: "ID不存在"}
	}
	return ctx.JSON(utils.SuccessResponse(book))
}

func AddBook(ctx *fiber.Ctx) error {
	book := new(dto.Book)
	if err := ctx.BodyParser(book); err != nil {
		return &fiber.Error{Code: 400, Message: err.Error()}
	}
	err := db.AddBook(book)

	if err != nil {
		return &fiber.Error{Code: 400, Message: err.Error()}
	}
	return ctx.JSON(utils.SuccessResponse(book))

}

func UpdateBook(ctx *fiber.Ctx) error {
	book := new(dto.Book)
	if err := ctx.BodyParser(book); err != nil {
		return &fiber.Error{Code: 400, Message: err.Error()}
	}
	if book.Id <= 0 {
		return &fiber.Error{Code: 400, Message: "ID不合法"}
	}
	err := db.UpdateBook(*book)

	if err != nil {
		return &fiber.Error{Code: 400, Message: err.Error()}
	}

	return ctx.JSON(utils.SuccessResponse(book))

}

func DeleteBook(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return &fiber.Error{Code: 400, Message: "ID不合法"}
	}
	err, exist := db.DeleteBook(id)
	if err != nil {
		return &fiber.Error{Code: 400, Message: err.Error()}
	}
	if !exist {
		return &fiber.Error{Code: 400, Message: "ID不存在"}
	}
	return ctx.JSON(utils.SuccessResponse(nil))
}
