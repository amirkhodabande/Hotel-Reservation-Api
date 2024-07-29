package tests

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"hotel.com/api"
	"hotel.com/types"
)

func TestGetHotelList(t *testing.T) {
	app, _ := setup(t, services())

	user := factory.CreateUser(map[string]any{})
	hotel := factory.CreateHotel(map[string]any{})

	req := httptest.NewRequest("GET", "/api/v1/hotels?page=1", nil)
	req.Header.Add("Content-Type", "application-json")
	loginAs(user, req)

	res, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}

	encodedRes, _ := io.ReadAll(res.Body)

	exceptedRes, _ := json.Marshal(
		api.SuccessResponse([]types.Hotel{*hotel}).WithPagination(1, 1),
	)

	assert.Equal(t, 200, res.StatusCode)
	assert.JSONEq(t, string(exceptedRes), string(encodedRes))
}
