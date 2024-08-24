package services

import (
	"book_management/db"
	"book_management/dto"
	"book_management/utils"
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
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
	username := utils.ParseJwt(ctx.Locals("user").(*jwt.Token))

	fmt.Printf("username: %v\n", username)

	book := new(dto.Book)
	if err := ctx.BodyParser(book); err != nil {
		return &fiber.Error{Code: 400, Message: err.Error()}
	}

	if book.Name == "" {
		return &fiber.Error{Code: 400, Message: "书名不能为空"}
	}

	if book.Author == "" {
		return &fiber.Error{Code: 400, Message: "作者不能为空"}
	}

	if book.Cover == "" {
		return &fiber.Error{Code: 400, Message: "封面不能为空"}
	}

	if book.Description == "" {
		book.Description = "暂无描述"
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

func Upload(ctx *fiber.Ctx) error {
	file, err := ctx.FormFile("file")
	if err != nil {
		return &fiber.Error{Code: 400, Message: err.Error()}
	}
	// 只能上传图片
	fmt.Printf("file.Filename: %v\n", file.Filename)
	typeList := []string{".jpg", ".png", ".jpeg"}
	flag := false
	for _, v := range typeList {
		if strings.Contains(file.Filename, v) {
			flag = true
			break
		}
	}
	if !flag {
		return &fiber.Error{Code: 400, Message: "只能上传图片"}
	}
	
	if err := ctx.SaveFile(file, "storage/"+file.Filename); err != nil {
		return &fiber.Error{Code: 400, Message: err.Error()}
	}
	return ctx.JSON(utils.SuccessResponse("/static/" + file.Filename))
}
