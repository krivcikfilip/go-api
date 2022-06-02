package services

import (
	"go-api/database"
	. "go-api/models"
)

func FindCategories() ([]Category, error) {
	var categories []Category

	res := database.GORM.Find(&categories)
	if res.Error != nil {
		return categories, res.Error
	}

	return categories, nil
}
