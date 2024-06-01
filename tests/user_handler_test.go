package tests

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"hotel.com/api"
	"hotel.com/types"
)

func TestCreateUser(t *testing.T) {
	tdb := setup(t)
	defer tdb.teardown(t)

	app := fiber.New()
	userHandler := api.NewUserHandler(tdb.UserStore)
	app.Post("/", userHandler.HandleCreateUser)

	params := types.CreateUserParams{
		Email:     "test@gmail.com",
		FirstName: "test",
		LastName:  "Ltest",
		Password:  "password",
	}

	b, _ := json.Marshal(params)
	req := httptest.NewRequest("POST", "/", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application-json")

	res, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}

	var user types.User
	json.NewDecoder(res.Body).Decode(&user)

	assert.NotEmpty(t, user.ID)
	assert.Equal(t, params.FirstName, user.FirstName)
	assert.Equal(t, params.LastName, user.LastName)
	assert.Empty(t, user.EncryptedPassword)

	// b, _ = io.ReadAll(res.Body)
	// fmt.Println(string(b))
}
