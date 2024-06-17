package validators

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"hotel.com/types"
)

func ValidateLogin(c *fiber.Ctx) error {
	params := new(types.LoginParams)
	c.BodyParser(&params)

	customValidator := &customValidator{
		validator: validate,
	}

	if errs := customValidator.validate(params); len(errs) > 0 && errs[0].HasError {
		errorMessages := make(map[string]string)

		for _, err := range errs {
			errorMessages[err.Field] = fmt.Sprintf("%s field failed. Validation: '%s'", err.Field, err.Tag)
		}

		return c.Status(fiber.StatusUnprocessableEntity).JSON(map[string]map[string]string{
			"errors": errorMessages,
		})
	}

	return c.Next()
}
