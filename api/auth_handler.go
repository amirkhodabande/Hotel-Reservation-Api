package api

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
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
		fmt.Println("p")

		return fmt.Errorf("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":      user.ID,
		"email":   user.Email,
		"expires": time.Now().Add(time.Hour * 4),
	})
	secret := os.Getenv("JWT_SECRET")
	fmt.Println(secret)
	tokenStr, _ := token.SignedString([]byte("secret"))

	return c.JSON(map[string]any{
		"user":  user,
		"token": tokenStr,
	})
}
