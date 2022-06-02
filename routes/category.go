package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-api/services"
)

func FindCategories(c *fiber.Ctx) error {
	categories, err := services.FindCategories()

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "Categories were not found",
		})
	}

	return c.Status(200).JSON(categories)
}
