package api

import (
	"github.com/gofiber/fiber/v2"
	"hotel.com/types"
)

func HandleGetUsers(c *fiber.Ctx) error {
	user := types.User{
		FirstName: "TestFirst",
		LastName:  "TestLast",
	}
	return c.JSON([]types.User{user})
}

func HandleGetUser(c *fiber.Ctx) error {
	return c.JSON("Single Useer")
}
