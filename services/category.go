package services

import (
	"go-api/database"
	. "go-api/errors"
	. "go-api/inputs"
	. "go-api/models"
)

// FindCategories
// find categories
func FindCategories() ([]Category, ErrorResponse) {
	var categories []Category

	res := database.GORM.Find(&categories)
	if res.Error != nil {
		return []Category{}, ErrorResponse{Message: "Categories were not found", Status: 404}
	}

	return categories, ErrorResponse{}
}

// FindCategoryById
// find category by id
func FindCategoryById(id int) (Category, ErrorResponse) {
	var category Category

	res := database.GORM.First(&category, id)
	if res.Error != nil {
		return Category{}, ErrorResponse{Message: "Category was not found", Status: 404}
	}

	return category, ErrorResponse{}
}

// CreateCategory
// create category from input
func CreateCategory(categoryInput CreateCategoryInput) (Category, ErrorResponse) {
	var category Category

	category.Name = categoryInput.Name

	res := database.GORM.Create(&category)
	if res.Error != nil {
		return Category{}, ErrorResponse{Message: res.Error.Error(), Status: 400}
	}

	return category, ErrorResponse{}
}

// UpdateCategory
// update category from input
func UpdateCategory(id int, categoryInput UpdateCategoryInput) (Category, ErrorResponse) {
	category, err := FindCategoryById(id)
	if err.Message != "" {
		return Category{}, err
	}

	if categoryInput.Name != "" {
		category.Name = categoryInput.Name
	}

	res := database.GORM.Save(&category)
	if res.Error != nil {
		return Category{}, ErrorResponse{Message: res.Error.Error(), Status: 400}
	}

	return category, ErrorResponse{}
}
