package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"hotel.com/types"
)

func TestGetRoomsList(t *testing.T) {
	app, tdb := setup(t)

	user := &types.User{
		Email:             "test@gmail.com",
		FirstName:         "test",
		LastName:          "Ltest",
		EncryptedPassword: "testEncrypted",
	}
	tdb.UserStore.Insert(context.Background(), user)
	hotel, _ := tdb.HotelStore.Insert(context.Background(), &types.Hotel{
		Name:     "TestHotel",
		Location: "Iran",
		Rating:   3,
		Rooms:    []primitive.ObjectID{},
	})
	room, _ := tdb.RoomStore.Insert(context.Background(), &types.Room{
		Type:    types.DoubleRoomType,
		Price:   900,
		HotelID: hotel.ID,
	})

	req := httptest.NewRequest("GET", fmt.Sprintf("/api/v1/hotels/%s/rooms", hotel.ID.Hex()), nil)
	req.Header.Add("Content-Type", "application-json")
	loginAs(user, req)

	res, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}

	encodedRes, _ := io.ReadAll(res.Body)
	exceptedRes, _ := json.Marshal([]types.Room{*room})

	assert.Equal(t, 200, res.StatusCode)
	assert.JSONEq(t, string(exceptedRes), string(encodedRes))
}
