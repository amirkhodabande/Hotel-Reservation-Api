package app

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"hotel.com/db"
	"hotel.com/routes"
)

var app = fiber.New(fiber.Config{
	ErrorHandler: func(ctx *fiber.Ctx, err error) error {
		if err.Error() == "unauthorized" || err.Error() == "invalid credentials" {
			ctx.Status(fiber.StatusUnauthorized)
		}

		if errors.Is(err, mongo.ErrNoDocuments) {
			ctx.Status(fiber.StatusNotFound)
		}

		return ctx.JSON(map[string]string{"error": err.Error()})
	},
})

func New(database *db.Store) *fiber.App {
	app := fiber.New(app.Config())

	routes.RegisterRoutes(database, app)

	return app
}
