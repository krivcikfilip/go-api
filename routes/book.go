package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-api/errors"
	"go-api/inputs"
	"go-api/services"
	"go-api/utils"
)

func CreateBookRoutes(api fiber.Router) {
	router := api.Group("/books")

	router.Get("/", findBooks)
	router.Get("/:id", findBook)
	router.Put("/:id", updateBook)
	router.Post("/", createBook)
}

func findBooks(c *fiber.Ctx) error {
	books, err := services.FindBooks()
	if err.Message != "" {
		return errors.CustomError(c, err)
	}

	return c.Status(200).JSON(books)
}

func findBook(c *fiber.Ctx) error {
	bookId, _ := c.ParamsInt("id")

	book, err := services.FindBook(bookId)
	if err.Message != "" {
		return errors.CustomError(c, err)
	}

	return c.Status(200).JSON(book)
}

func createBook(c *fiber.Ctx) error {
	var bookInput inputs.CreateBookInput

	err := utils.ParseInput(c, &bookInput)
	if err.Message != "" {
		return errors.CustomError(c, err)
	}

	errF := utils.ValidateInput(&bookInput)
	if errF.Message != "" {
		return errors.InvalidField(c, errF)
	}

	book, err := services.CreateBook(bookInput)
	if err.Message != "" {
		return errors.CustomError(c, err)
	}

	return c.Status(200).JSON(book)
}

func updateBook(c *fiber.Ctx) error {
	bookId, _ := c.ParamsInt("id")

	var bookInput inputs.UpdateBookInput

	err := utils.ParseInput(c, &bookInput)
	if err.Message != "" {
		return errors.CustomError(c, err)
	}

	errF := utils.ValidateInput(&bookInput)
	if errF.Message != "" {
		return errors.InvalidField(c, errF)
	}

	book, err := services.UpdateBook(bookId, bookInput)
	if err.Message != "" {
		return errors.CustomError(c, err)
	}

	return c.Status(200).JSON(book)
}
