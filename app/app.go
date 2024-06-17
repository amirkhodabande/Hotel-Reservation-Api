package app

import (
	"github.com/gofiber/fiber/v2"
	"hotel.com/db"
	"hotel.com/routes"
)

var app = fiber.New(fiber.Config{
	ErrorHandler: func(ctx *fiber.Ctx, err error) error {
		if err.Error() == "unauthorized" || err.Error() == "invalid credentials" {
			ctx.Status(fiber.StatusUnauthorized)
		}

		return ctx.JSON(map[string]string{"error": err.Error()})
	},
})

func New(database *db.Store) *fiber.App {
	app := fiber.New(app.Config())

	routes.RegisterRoutes(database, app)

	return app
}
