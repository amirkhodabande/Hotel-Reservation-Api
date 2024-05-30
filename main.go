package main

import (
	"context"
	"flag"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"hotel.com/api"
	"hotel.com/api/validators"
	"hotel.com/db"
)

const dburi = "mongodb://localhost:27017"

var app = fiber.New(fiber.Config{
	ErrorHandler: func(ctx *fiber.Ctx, err error) error {
		return ctx.JSON(map[string]string{"error": err.Error()})
	},
})

func main() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}

	userHandler := api.NewUserHandler(db.NewMongoUserStore(client))

	address := flag.String("serverPort", ":5000", "")
	flag.Parse()

	app := fiber.New(app.Config())

	apiV1 := app.Group("api/v1")
	apiV1.Get("/", handleHome)

	apiV1.Get("/users", userHandler.HandleGetUsers)
	apiV1.Post("/users", validators.ValidateCreateUser, userHandler.HandleCreateUser)
	apiV1.Get("/users/:id", userHandler.HandleGetUser)
	apiV1.Put("/users/:id", validators.ValidateUpdateUser, userHandler.HandleUpdateUser)
	apiV1.Delete("/users/:id", userHandler.HandleDeleteUser)

	app.Listen(*address)
}

func handleHome(c *fiber.Ctx) error {

	return c.JSON(map[string]string{"msg": "working"})
}
