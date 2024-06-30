package middlewares

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"hotel.com/api/custom_errors"
	"hotel.com/db"
)

func Authenticate(db *db.Store) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		token, ok := ctx.GetReqHeaders()["Authorization"]

		if !ok {
			return custom_errors.Unauthorized()
		}

		claims, err := validateToken(token[0])
		if err != nil {
			return custom_errors.Unauthorized()
		}

		expires := claims["expires"].(string)

		if time.Now().String() > expires {
			return custom_errors.Unauthorized()
		}

		userID := claims["id"].(string)
		user, err := db.UserStore.GetByID(ctx.Context(), userID)

		if err != nil {
			return custom_errors.Unauthorized()
		}

		ctx.Context().SetUserValue("user", user)
		return ctx.Next()
	}
}

func validateToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {

			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		secret := os.Getenv("JWT_SECRET")

		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("unauthorized")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("unauthorized")
	}

	return claims, nil
}
