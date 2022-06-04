package services

import (
	"go-api/database"
	. "go-api/errors"
	. "go-api/inputs"
	. "go-api/models"
)

// findBookCategories
// find and append book categories by ids
func findBookCategories(book *Book, ids []int) ErrorResponse {
	var categories []Category

	if len(ids) == 0 {
		return ErrorResponse{}
	}

	res := database.GORM.Find(&categories, ids)
	if res.Error != nil {
		return ErrorResponse{Message: "Categories were not found", Status: 404}
	}

	book.Categories = categories
	return ErrorResponse{}
}

// FindBooks
// find books
func FindBooks() ([]Book, ErrorResponse) {
	var books []Book

	res := database.GORM.Preload("Categories").Find(&books)
	if res.Error != nil {
		return []Book{}, ErrorResponse{Message: "Books were not found", Status: 404}
	}

	return books, ErrorResponse{}
}

// FindBookById
// find book by id
func FindBookById(id int) (Book, ErrorResponse) {
	var book Book

	res := database.GORM.Preload("Categories").First(&book, id)
	if res.Error != nil {
		return Book{}, ErrorResponse{Message: "Book was not found", Status: 404}
	}

	return book, ErrorResponse{}
}

// CreateBook
// create book from input
func CreateBook(bookInput CreateBookInput) (Book, ErrorResponse) {
	var book Book

	book.Name = bookInput.Name

	err := findBookCategories(&book, bookInput.Categories)
	if err.Message != "" {
		return Book{}, err
	}

	res := database.GORM.Create(&book)
	return book, ErrorResponse{Message: res.Error.Error(), Status: 400}
}

// UpdateBook
// update book from input
func UpdateBook(id int, bookInput UpdateBookInput) (Book, ErrorResponse) {
	book, err := FindBookById(id)
	if err.Message != "" {
		return Book{}, err
	}

	if bookInput.Name != "" {
		book.Name = bookInput.Name
	}

	err = findBookCategories(&book, bookInput.Categories)
	if err.Message != "" {
		return Book{}, err
	}

	res := database.GORM.Save(book)
	return book, ErrorResponse{Message: res.Error.Error(), Status: 400}
}
