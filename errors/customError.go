package errors

import "github.com/gofiber/fiber/v2"

func CustomError(c *fiber.Ctx, response ErrorResponse) error {
	return c.Status(response.Status).JSON(CreateErrorMap(response))
}
