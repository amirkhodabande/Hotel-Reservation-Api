package app

import (
	"github.com/gofiber/fiber/v2"
	"hotel.com/api/custom_errors"
	"hotel.com/db"
	"hotel.com/routes"
)

var app = fiber.New(fiber.Config{
	ErrorHandler: func(ctx *fiber.Ctx, err error) error {
		if customError, ok := err.(custom_errors.Error); ok {
			return ctx.Status(customError.Status()).JSON(customError)
		}

		return ctx.JSON(map[string]string{"error": err.Error()})
	},
})

func New(database *db.Store) *fiber.App {
	app := fiber.New(app.Config())

	routes.RegisterRoutes(database, app)

	return app
}
