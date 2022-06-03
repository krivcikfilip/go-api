package errors

import "github.com/gofiber/fiber/v2"

func InvalidField(c *fiber.Ctx, errorField ErrorResponseWithField) error {
	return c.Status(400).JSON(CreateErrorMapWithField(errorField))
}
