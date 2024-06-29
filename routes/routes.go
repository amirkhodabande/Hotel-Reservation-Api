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
	bookingHandler := api.NewBookingHandler(database)

	apiV1 := app.Group("api/v1")

	apiV1.Get("/", handleHome)

	// authentication
	apiV1.Post("/login", validators.ValidateLogin, authHandler.HandleLogin)

	// user routes
	userRoutes := apiV1.Group("users")
	userRoutes.Get("/", middlewares.Authenticate(database), userHandler.HandleGetUsers)
	userRoutes.Post("/", validators.ValidateCreateUser, userHandler.HandleCreateUser)
	userRoutes.Get("/:id", middlewares.Authenticate(database), userHandler.HandleGetUser)
	userRoutes.Put("/:id", middlewares.Authenticate(database), validators.ValidateUpdateUser, userHandler.HandleUpdateUser)
	userRoutes.Delete("/:id", middlewares.Authenticate(database), userHandler.HandleDeleteUser)

	// hotel routes
	hotelRoutes := apiV1.Group("hotels", middlewares.Authenticate(database))
	hotelRoutes.Get("/", hotelHandler.HandleGetHotels)

	// room routes
	hotelRoutes.Get("/:id/rooms", roomHandler.HandleGetRooms)

	// booking routes
	bookingRoutes := apiV1.Group("/bookings", middlewares.Authenticate(database))
	bookingRoutes.Get("/", bookingHandler.HandleGetBookings)
	bookingRoutes.Post("/:id", validators.ValidateBookingRoom, bookingHandler.HandleBookRoom)
	bookingRoutes.Delete("/:id", bookingHandler.HandleCancelBooking)
}

func handleHome(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "working"})
}
