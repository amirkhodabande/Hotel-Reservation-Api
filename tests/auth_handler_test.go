package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"hotel.com/types"
)

func TestCanLoginSuccessfully(t *testing.T) {
	app, tdb := setup(t)

	user, _ := types.NewUserFromParams(types.CreateUserParams{
		Email:     "test1@gmail.com",
		FirstName: "test1",
		LastName:  "Ltest1",
		Password:  "password",
	})
	anotherUser, _ := types.NewUserFromParams(types.CreateUserParams{
		Email:     "test2@gmail.com",
		FirstName: "test2",
		LastName:  "Ltest2",
		Password:  "password2",
	})
	tdb.UserStore.Insert(context.Background(), user)
	tdb.UserStore.Insert(context.Background(), anotherUser)

	loginParams, _ := json.Marshal(types.LoginParams{
		Email:    "test1@gmail.com",
		Password: "password",
	})

	req := httptest.NewRequest("POST", "/api/v1/login", bytes.NewReader(loginParams))
	req.Header.Add("Content-Type", "application-json")
	loginAs(user, req)

	res, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}

	encodedRes, _ := io.ReadAll(res.Body)
	var loginRes map[string]types.User
	json.Unmarshal(encodedRes, &loginRes)

	assert.Equal(t, 200, res.StatusCode)
	assert.Equal(t, user.Email, loginRes["user"].Email)
}

func TestCannotLoginWithWrongCredentials(t *testing.T) {
	app, tdb := setup(t)

	user, _ := types.NewUserFromParams(types.CreateUserParams{
		Email:     "test1@gmail.com",
		FirstName: "test1",
		LastName:  "Ltest1",
		Password:  "password",
	})
	tdb.UserStore.Insert(context.Background(), user)

	loginParams, _ := json.Marshal(types.LoginParams{
		Email:    "wrong@gmail.com",
		Password: "wrong-password",
	})

	req := httptest.NewRequest("POST", "/api/v1/login", bytes.NewReader(loginParams))
	req.Header.Add("Content-Type", "application-json")

	res, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}

	encodedRes, _ := io.ReadAll(res.Body)
	var loginRes map[string]types.User
	json.Unmarshal(encodedRes, &loginRes)

	assert.Equal(t, 401, res.StatusCode)
}

func TestLoginValidation(t *testing.T) {
	app, _ := setup(t)

	tests := []struct {
		description string
		params      types.CreateUserParams
	}{
		{
			description: "email is required",
			params: types.CreateUserParams{
				Password: "password",
			},
		},
		{
			description: "email should be valid",
			params: types.CreateUserParams{
				Email:    "invalid",
				Password: "password",
			},
		},
		{
			description: "password is required",
			params: types.CreateUserParams{
				Email: "test@gmail.com",
			},
		},
	}

	for _, test := range tests {
		b, _ := json.Marshal(test.params)
		req := httptest.NewRequest("POST", "/api/v1/login", bytes.NewReader(b))
		req.Header.Add("Content-Type", "application-json")

		res, _ := app.Test(req)

		assert.Equalf(t, 422, res.StatusCode, test.description)
	}
}
