package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"hotel.com/api"
	"hotel.com/api/validators"
	"hotel.com/types"
)

func TestGetUserList(t *testing.T) {
	tdb := setup(t)
	defer tdb.teardown(t)

	app := fiber.New()
	userHandler := api.NewUserHandler(tdb.UserStore)
	app.Get("/", userHandler.HandleGetUsers)

	user := &types.User{
		Email:             "test@gmail.com",
		FirstName:         "test",
		LastName:          "Ltest",
		EncryptedPassword: "testEncrypted",
	}
	anotherUser := &types.User{
		Email:             "test@gmail.com",
		FirstName:         "test",
		LastName:          "Ltest",
		EncryptedPassword: "testEncrypted",
	}
	tdb.UserStore.InsertUser(context.Background(), user)
	tdb.UserStore.InsertUser(context.Background(), anotherUser)

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Add("Content-Type", "application-json")

	res, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}

	// TODO: search to find better way
	bodyByte, _ := io.ReadAll(res.Body)
	userList := []types.User{*user, *anotherUser}
	resByte, _ := json.Marshal(userList)

	assert.Equal(t, 200, res.StatusCode)
	assert.JSONEq(t, string(resByte), string(bodyByte))
}

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

	assert.Equal(t, 201, res.StatusCode)
	assert.NotEmpty(t, user.ID)
	assert.Equal(t, params.FirstName, user.FirstName)
	assert.Equal(t, params.LastName, user.LastName)
	assert.Empty(t, user.EncryptedPassword)
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

func TestGetUserBy(t *testing.T) {
	tdb := setup(t)
	defer tdb.teardown(t)

	app := fiber.New()
	userHandler := api.NewUserHandler(tdb.UserStore)
	app.Get("/:id", userHandler.HandleGetUser)

	user := &types.User{
		Email:             "test@gmail.com",
		FirstName:         "test",
		LastName:          "Ltest",
		EncryptedPassword: "testEncrypted",
	}
	tdb.UserStore.InsertUser(context.Background(), user)

	req := httptest.NewRequest("GET", fmt.Sprint("/", user.ID.Hex()), nil)
	req.Header.Add("Content-Type", "application-json")

	res, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}

	bodyByte, _ := io.ReadAll(res.Body)
	resByte, _ := json.Marshal(user)

	assert.Equal(t, 200, res.StatusCode)
	assert.JSONEq(t, string(resByte), string(bodyByte))
}

func TestUpdateUser(t *testing.T) {
	tdb := setup(t)
	defer tdb.teardown(t)

	app := fiber.New()
	userHandler := api.NewUserHandler(tdb.UserStore)
	app.Put("/:id", userHandler.HandleUpdateUser)

	user := &types.User{
		Email:             "test@gmail.com",
		FirstName:         "test",
		LastName:          "Ltest",
		EncryptedPassword: "testEncrypted",
	}
	tdb.UserStore.InsertUser(context.Background(), user)

	updateParams := types.UpdateUserParams{
		FirstName: "updatedName",
		LastName:  "updatedLastName",
	}
	b, _ := json.Marshal(updateParams)

	req := httptest.NewRequest("PUT", fmt.Sprint("/", user.ID.Hex()), bytes.NewReader(b))
	req.Header.Add("Content-Type", "application-json")

	res, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}

	newUser, _ := tdb.UserStore.GetUserByID(context.Background(), user.ID.Hex())

	assert.Equal(t, 200, res.StatusCode)
	assert.Equal(t, updateParams.FirstName, newUser.FirstName)
	assert.Equal(t, updateParams.LastName, newUser.LastName)
}

func TestDeleteUser(t *testing.T) {
	tdb := setup(t)
	defer tdb.teardown(t)

	app := fiber.New()
	userHandler := api.NewUserHandler(tdb.UserStore)
	app.Delete("/:id", userHandler.HandleDeleteUser)

	user := &types.User{
		Email:             "test@gmail.com",
		FirstName:         "test",
		LastName:          "Ltest",
		EncryptedPassword: "testEncrypted",
	}
	tdb.UserStore.InsertUser(context.Background(), user)

	req := httptest.NewRequest("DELETE", fmt.Sprint("/", user.ID.Hex()), nil)
	req.Header.Add("Content-Type", "application-json")

	res, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}

	u, _ := tdb.UserStore.GetUserByID(context.Background(), user.ID.Hex())

	assert.Equal(t, 200, res.StatusCode)
	assert.Empty(t, u)
}
