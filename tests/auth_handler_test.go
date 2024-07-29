package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"hotel.com/types"
)

func TestCanLoginSuccessfully(t *testing.T) {
	app, _ := setup(t, services())

	user := factory.CreateUser(map[string]any{})
	factory.CreateUser(map[string]any{
		"email": "test2@gmail.com",
	})

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
	var loginRes map[string]map[string]types.User
	json.Unmarshal(encodedRes, &loginRes)

	assert.Equal(t, 200, res.StatusCode)
	assert.Equal(t, user.Email, loginRes["data"]["user"].Email)
}

func TestCannotLoginWithWrongCredentials(t *testing.T) {
	app, _ := setup(t, services())

	factory.CreateUser(map[string]any{})

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
	app, _ := setup(t, services())

	tests := []struct {
		description string
		params      types.LoginParams
	}{
		{
			description: "email is required",
			params: types.LoginParams{
				Password: "password",
			},
		},
		{
			description: "email should be valid",
			params: types.LoginParams{
				Email:    "invalid",
				Password: "password",
			},
		},
		{
			description: "password is required",
			params: types.LoginParams{
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
