package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"hotel.com/types"
)

func TestGetUserList(t *testing.T) {
	app, tdb := setup(t)

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
	tdb.UserStore.Insert(context.Background(), user)
	tdb.UserStore.Insert(context.Background(), anotherUser)

	req := httptest.NewRequest("GET", "/api/v1/users", nil)
	req.Header.Add("Content-Type", "application-json")
	loginAs(user, req)

	res, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}

	encodedRes, _ := io.ReadAll(res.Body)
	exceptedRes, _ := json.Marshal([]types.User{*user, *anotherUser})

	assert.Equal(t, 200, res.StatusCode)
	assert.JSONEq(t, string(exceptedRes), string(encodedRes))
}

func TestCreateUser(t *testing.T) {
	app, _ := setup(t)

	params := types.CreateUserParams{
		Email:     "test@gmail.com",
		FirstName: "test",
		LastName:  "Ltest",
		Password:  "password",
	}

	b, _ := json.Marshal(params)
	req := httptest.NewRequest("POST", "/api/v1/users", bytes.NewReader(b))
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
	app, _ := setup(t)

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
		req := httptest.NewRequest("POST", "/api/v1/users", bytes.NewReader(b))
		req.Header.Add("Content-Type", "application-json")

		res, _ := app.Test(req)

		assert.Equalf(t, 422, res.StatusCode, test.description)
	}
}

func TestGetUserBy(t *testing.T) {
	app, tdb := setup(t)

	user := &types.User{
		Email:             "test@gmail.com",
		FirstName:         "test",
		LastName:          "Ltest",
		EncryptedPassword: "testEncrypted",
	}
	tdb.UserStore.Insert(context.Background(), user)

	req := httptest.NewRequest("GET", fmt.Sprint("/api/v1/users/", user.ID.Hex()), nil)
	req.Header.Add("Content-Type", "application-json")
	loginAs(user, req)

	res, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}

	encodedRes, _ := io.ReadAll(res.Body)
	exceptedRes, _ := json.Marshal(user)

	assert.Equal(t, 200, res.StatusCode)
	assert.JSONEq(t, string(exceptedRes), string(encodedRes))
}

func TestUpdateUser(t *testing.T) {
	app, tdb := setup(t)

	user := &types.User{
		Email:             "test@gmail.com",
		FirstName:         "test",
		LastName:          "Ltest",
		EncryptedPassword: "testEncrypted",
	}
	tdb.UserStore.Insert(context.Background(), user)

	updateParams := types.UpdateUserParams{
		FirstName: "updatedName",
		LastName:  "updatedLastName",
	}
	b, _ := json.Marshal(updateParams)

	req := httptest.NewRequest("PUT", fmt.Sprint("/api/v1/users/", user.ID.Hex()), bytes.NewReader(b))
	req.Header.Add("Content-Type", "application-json")
	loginAs(user, req)

	res, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}

	newUser, _ := tdb.UserStore.GetByID(context.Background(), user.ID.Hex())

	assert.Equal(t, 200, res.StatusCode)
	assert.Equal(t, updateParams.FirstName, newUser.FirstName)
	assert.Equal(t, updateParams.LastName, newUser.LastName)
}

func TestDeleteUser(t *testing.T) {
	app, tdb := setup(t)

	user := &types.User{
		Email:             "test@gmail.com",
		FirstName:         "test",
		LastName:          "Ltest",
		EncryptedPassword: "testEncrypted",
	}
	tdb.UserStore.Insert(context.Background(), user)

	req := httptest.NewRequest("DELETE", fmt.Sprint("/api/v1/users/", user.ID.Hex()), nil)
	req.Header.Add("Content-Type", "application-json")
	loginAs(user, req)

	res, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}

	u, _ := tdb.UserStore.GetByID(context.Background(), user.ID.Hex())

	assert.Equal(t, 200, res.StatusCode)
	assert.Empty(t, u)
}
