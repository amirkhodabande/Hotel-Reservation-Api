package tests

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"hotel.com/api"
	"hotel.com/api/validators"
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

func TestCreateUserValidation(t *testing.T) {
	tdb := setup(t)
	defer tdb.teardown(t)

	app := fiber.New()
	userHandler := api.NewUserHandler(tdb.UserStore)
	app.Post("/", validators.ValidateCreateUser, userHandler.HandleCreateUser)

	tests := []struct {
		description string
		params      types.CreateUserParams
	}{
		{
			description: "first_name is required",
			params: types.CreateUserParams{
				Email:    "test@gmail.com",
				LastName: "Ltest",
				Password: "password",
			},
		},
		{
			description: "first_name min is 2",
			params: types.CreateUserParams{
				Email:     "test@gmail.com",
				FirstName: "a",
				LastName:  "Ltest",
				Password:  "password",
			},
		},
		{
			description: "last_name is required",
			params: types.CreateUserParams{
				Email:     "test@gmail.com",
				FirstName: "Ltest",
				Password:  "password",
			},
		},
		{
			description: "last_name min is 2",
			params: types.CreateUserParams{
				Email:     "test@gmail.com",
				FirstName: "test",
				LastName:  "L",
				Password:  "password",
			},
		},
	}

	for _, test := range tests {
		b, _ := json.Marshal(test.params)
		req := httptest.NewRequest("POST", "/", bytes.NewReader(b))
		req.Header.Add("Content-Type", "application-json")

		res, _ := app.Test(req)

		assert.Equalf(t, 422, res.StatusCode, test.description)
	}
}
