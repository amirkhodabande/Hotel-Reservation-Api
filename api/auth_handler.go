package api

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"hotel.com/db"
	"hotel.com/types"
)

type AuthHandler struct {
	*db.Store
}

func NewAuthHandler(store *db.Store) *AuthHandler {
	return &AuthHandler{
		store,
	}
}

func (h *AuthHandler) HandleLogin(c *fiber.Ctx) error {
	var data types.LoginParams
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	user, err := h.UserStore.GetByEmail(c.Context(), data.Email)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return fmt.Errorf("invalid credentials")
		}
		return err
	}

	if !types.IsValidPassword(user.EncryptedPassword, data.Password) {
		return fmt.Errorf("invalid credentials")
	}

	tokenStr := user.CreateToken()

	return c.JSON(map[string]any{
		"user":  user,
		"token": tokenStr,
	})
}
