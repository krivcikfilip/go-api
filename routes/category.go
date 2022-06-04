package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-api/errors"
	"go-api/inputs"
	"go-api/services"
	"go-api/utils"
)

// CreateCategoriesRoutes
// create categories routes
func CreateCategoriesRoutes(api fiber.Router) {
	router := api.Group("/categories")

	router.Get("/", findCategories)
	router.Get("/:id", findCategory)
	router.Put("/:id", updateCategory)
	router.Post("/", createCategory)
}

// findCategories
// find categories route
func findCategories(c *fiber.Ctx) error {
	categories, err := services.FindCategories()
	if err.Message != "" {
		return errors.CustomError(c, err)
	}

	return c.Status(200).JSON(categories)
}

// findCategory
// find category route
func findCategory(c *fiber.Ctx) error {
	categoryId, _ := c.ParamsInt("id")

	category, err := services.FindCategoryById(categoryId)
	if err.Message != "" {
		return errors.CustomError(c, err)
	}

	return c.Status(200).JSON(category)
}

// createCategory
// create category route
func createCategory(c *fiber.Ctx) error {
	var categoryInput inputs.CreateCategoryInput

	err := utils.ParseInput(c, &categoryInput)
	if err.Message != "" {
		return errors.CustomError(c, err)
	}

	errF := utils.ValidateInput(&categoryInput)
	if errF.Message != "" {
		return errors.InvalidField(c, errF)
	}

	category, err := services.CreateCategory(categoryInput)
	if err.Message != "" {
		return errors.CustomError(c, err)
	}

	return c.Status(200).JSON(category)
}

// updateCategory
// update category route
func updateCategory(c *fiber.Ctx) error {
	categoryId, _ := c.ParamsInt("id")

	var categoryInput inputs.UpdateCategoryInput

	err := utils.ParseInput(c, &categoryInput)
	if err.Message != "" {
		return errors.CustomError(c, err)
	}

	errF := utils.ValidateInput(&categoryInput)
	if errF.Message != "" {
		return errors.InvalidField(c, errF)
	}

	category, err := services.UpdateCategory(categoryId, categoryInput)
	if err.Message != "" {
		return errors.CustomError(c, err)
	}

	return c.Status(200).JSON(category)
}
