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

var app = fiber.New(fiber.Config{
	ErrorHandler: func(ctx *fiber.Ctx, err error) error {
		return ctx.JSON(map[string]string{"error": err.Error()})
	},
})

func main() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(db.DBuri))
	if err != nil {
		log.Fatal(err)
	}

	userHandler := api.NewUserHandler(db.NewMongoUserStore(client, db.DBname))

	hotelStore := db.NewMongoHotelStore(client, db.DBname)
	hotelHandler := api.NewHotelHandler(hotelStore)
	roomHandler := api.NewRoomHandler(db.NewMongoRoomStore(client, db.DBname, hotelStore))

	address := flag.String("serverPort", ":5000", "")
	flag.Parse()

	app := fiber.New(app.Config())

	apiV1 := app.Group("api/v1")
	apiV1.Get("/", handleHome)

	// user routes
	apiV1.Get("/users", userHandler.HandleGetUsers)
	apiV1.Post("/users", validators.ValidateCreateUser, userHandler.HandleCreateUser)
	apiV1.Get("/users/:id", userHandler.HandleGetUser)
	apiV1.Put("/users/:id", validators.ValidateUpdateUser, userHandler.HandleUpdateUser)
	apiV1.Delete("/users/:id", userHandler.HandleDeleteUser)

	// hotel routes
	apiV1.Get("/hotels", hotelHandler.HandleGetHotels)

	// room routes
	apiV1.Get("/hotels/:id/rooms", roomHandler.HandleGetRooms)

	app.Listen(*address)
}

func handleHome(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "working"})
}
