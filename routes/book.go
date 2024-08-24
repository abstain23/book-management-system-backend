package routes

import (
	"book_management/services"

	"github.com/gofiber/fiber/v2"
)


func initBookRouter(router fiber.Router) {
	bookGroup := router.Group("/book")
	bookGroup.Get("/list", list)
	bookGroup.Get("/:id", findById)
	bookGroup.Post("/create", create)
	bookGroup.Put("/update", update)
	bookGroup.Delete("/delete/:id", delete)
	bookGroup.Post("/upload", upload)
}

func list(ctx *fiber.Ctx) error {
	return services.GetBookList(ctx)
}

func findById(ctx *fiber.Ctx) error {
	return services.GetBookById(ctx)
}

func create(ctx *fiber.Ctx) error {
	return services.AddBook(ctx)
}

func update(ctx *fiber.Ctx) error {
	return services.UpdateBook(ctx)
}

func delete(ctx *fiber.Ctx) error {
	return services.DeleteBook(ctx)
}

func upload(ctx *fiber.Ctx) error {
	return services.Upload(ctx)
}
