package validators

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"hotel.com/types"
)

func ValidateBookingRoom(c *fiber.Ctx) error {
	params := new(types.BookRoomParams)
	c.BodyParser(&params)

	validate.RegisterValidation("valid_date", validDate)

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

	c.Context().SetUserValue("params", params)
	
	return c.Next()
}

func validDate(fl validator.FieldLevel) bool {
	date := fl.Field().Interface().(time.Time)
	return date.After(time.Now())
}
