package types

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type LoginParams struct {
	Email    string `json:"email" validate:"required,min=5,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type CreateUserParams struct {
	Email     string `json:"email" validate:"required,min=5,email"`
	FirstName string `json:"first_name" validate:"required,min=2"`
	LastName  string `json:"last_name" validate:"required,min=2"`
	Password  string `json:"password" validate:"required,min=6"`
}

type UpdateUserParams struct {
	FirstName string `bson:"firstName" json:"first_name" validate:"required,min=2"`
	LastName  string `bson:"lastName" json:"last_name" validate:"required,min=2"`
}

type User struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName         string             `bson:"firstName" json:"first_name"`
	LastName          string             `bson:"lastName" json:"last_name"`
	Email             string             `bson:"email" json:"email"`
	EncryptedPassword string             `bson:"encryptedPassword" json:"-"`
}

func (user *User) CreateToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":      user.ID,
		"email":   user.Email,
		"expires": time.Now().Add(time.Hour * 4),
	})
	secret := os.Getenv("JWT_SECRET")
	tokenStr, _ := token.SignedString([]byte(secret))

	return tokenStr
}

func NewUserFromParams(params *CreateUserParams) (*User, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.MinCost)

	if err != nil {
		return nil, err
	}

	return &User{
		FirstName:         params.FirstName,
		LastName:          params.LastName,
		Email:             params.Email,
		EncryptedPassword: string(encpw),
	}, nil
}

func IsValidPassword(encpw, pw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(encpw), []byte(pw)) == nil
}
