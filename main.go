package main

import (
	"flag"

	"github.com/gofiber/fiber/v2"
	"hotel.com/api"
)

func main() {
	address := flag.String("serverPort", ":5000", "")
	flag.Parse()

	app := fiber.New()

	apiV1 := app.Group("api/v1")
	apiV1.Get("/", handleHome)

	apiV1.Get("/users", api.HandleGetUsers)
	apiV1.Get("/users/:id", api.HandleGetUser)

	app.Listen(*address)
}

func handleHome(c *fiber.Ctx) error {
	
	return c.JSON(map[string]string{"msg": "working"})
}
