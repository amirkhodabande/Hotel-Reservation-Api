package tests

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"hotel.com/api"
	"hotel.com/types"
)

func TestGetRoomsList(t *testing.T) {
	app, _ := setup(t, services())

	user := factory.CreateUser(map[string]any{})
	hotel := factory.CreateHotel(map[string]any{})
	room := factory.CreateRoom(map[string]any{
		"hotel_id": hotel.ID,
	})

	req := httptest.NewRequest("GET", fmt.Sprintf("/api/v1/hotels/%s/rooms", hotel.ID.Hex()), nil)
	req.Header.Add("Content-Type", "application-json")
	loginAs(user, req)

	res, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}

	encodedRes, _ := io.ReadAll(res.Body)
	exceptedRes, _ := json.Marshal(api.SuccessResponse([]types.Room{*room}))

	assert.Equal(t, 200, res.StatusCode)
	assert.JSONEq(t, string(exceptedRes), string(encodedRes))
}
