package api

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"hotel.com/api/custom_errors"
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
	params := c.Context().UserValue("params").(*types.LoginParams)

	user, err := h.UserStore.GetByEmail(c.Context(), params.Email)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return custom_errors.InvalidCredentials()
		}
		return custom_errors.Internal()
	}

	if !types.IsValidPassword(user.EncryptedPassword, params.Password) {
		return fmt.Errorf("invalid credentials")
	}

	tokenStr := user.CreateToken()

	resp := SuccessResponse(map[string]any{
		"user":  user,
		"token": tokenStr,
	})
	return c.JSON(resp)
}
