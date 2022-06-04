package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	. "go-api/errors"
	"strings"
)

// createErrorMessage
// create error message by tag and value
func createErrorMessage(tag string, value string) string {
	switch tag {
	case "required":
		return "Required field"
	case "min":
		return "Minimum length: " + value
	}
	return ""
}

// ParseInput
// parse and transform request body
func ParseInput(c *fiber.Ctx, input interface{}) ErrorResponse {
	err := c.BodyParser(&input)
	if err != nil {
		return ErrorResponse{Message: err.Error(), Status: 400}
	}

	return ErrorResponse{}
}

// ValidateInput
// validate input body
func ValidateInput(input interface{}) ErrorResponseWithField {
	validate := validator.New()
	var errorField ErrorResponseWithField

	err := validate.Struct(input)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errorField.Field = strings.ToLower(err.Field())
			errorField.Message = createErrorMessage(err.Tag(), err.Param())
		}
	}

	errorField.Status = 400

	return errorField
}
