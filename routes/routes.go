package routes

import (
	"github.com/gofiber/fiber/v2"
	"hotel.com/api"
	"hotel.com/api/middlewares"
	"hotel.com/api/validators"
	"hotel.com/db"
)

func RegisterRoutes(database *db.Store, app *fiber.App) {
	authHandler := api.NewAuthHandler(database)
	userHandler := api.NewUserHandler(database)
	hotelHandler := api.NewHotelHandler(database)
	roomHandler := api.NewRoomHandler(database)

	apiV1 := app.Group("api/v1")

	apiV1.Get("/", handleHome)

	// authentication
	apiV1.Post("/login", validators.ValidateLogin, authHandler.HandleLogin)

	// user routes
	userRoutes := apiV1.Group("users")
	userRoutes.Get("/", middlewares.Authenticate, userHandler.HandleGetUsers)
	userRoutes.Post("/", validators.ValidateCreateUser, userHandler.HandleCreateUser)
	userRoutes.Get("/:id", middlewares.Authenticate, userHandler.HandleGetUser)
	userRoutes.Put("/:id", middlewares.Authenticate, validators.ValidateUpdateUser, userHandler.HandleUpdateUser)
	userRoutes.Delete("/:id", middlewares.Authenticate, userHandler.HandleDeleteUser)

	// hotel routes
	hotelRoutes := apiV1.Group("hotels", middlewares.Authenticate)
	hotelRoutes.Get("/", hotelHandler.HandleGetHotels)

	// room routes
	hotelRoutes.Get("/:id/rooms", roomHandler.HandleGetRooms)
}

func handleHome(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "working"})
}
