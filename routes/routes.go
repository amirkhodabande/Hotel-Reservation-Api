package routes

import (
	"github.com/gofiber/fiber/v2"
	"hotel.com/api"
	"hotel.com/api/validators"
	"hotel.com/db"
)

func RegisterRoutes(database *db.Store, app *fiber.App) {
	userHandler := api.NewUserHandler(database)
	hotelHandler := api.NewHotelHandler(database)
	roomHandler := api.NewRoomHandler(database)

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
}

func handleHome(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "working"})
}
