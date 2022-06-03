package errors

import "github.com/gofiber/fiber/v2"

type ErrorResponse struct {
	Status  int
	Message string
}

type ErrorResponseWithField struct {
	Status  int
	Message string
	Field   string
}

func CreateErrorMap(errorResponse ErrorResponse) interface{} {
	return fiber.Map{
		"error":  errorResponse.Message,
		"status": errorResponse.Status,
	}
}

func CreateErrorMapWithField(errorResponse ErrorResponseWithField) interface{} {
	return fiber.Map{
		"error":  errorResponse.Message,
		"field":  errorResponse.Field,
		"status": errorResponse.Status,
	}
}
