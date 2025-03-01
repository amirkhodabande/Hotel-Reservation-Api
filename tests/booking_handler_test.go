package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"hotel.com/api"
	"hotel.com/types"
)

func TestGetBookedRoomsList(t *testing.T) {
	app, tdb := setup(t, services())

	user := factory.CreateUser(map[string]any{})
	factory.CreateBooking(map[string]any{
		"user_id": user.ID,
	})

	req := httptest.NewRequest("GET", "/api/v1/bookings", nil)
	req.Header.Add("Content-Type", "application-json")
	loginAs(user, req)

	res, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}

	encodedRes, _ := io.ReadAll(res.Body)

	filter := &types.BookingQueryParams{UserID: user.ID}
	bookings, _ := tdb.BookingStore.Get(context.Background(), filter)
	exceptedRes, _ := json.Marshal(
		api.SuccessResponse(bookings).WithPagination(int64(len(bookings)), filter.GetPage()),
	)

	assert.Equal(t, 200, res.StatusCode)
	assert.JSONEq(t, string(exceptedRes), string(encodedRes))
}

func TestBookingRoom(t *testing.T) {
	app, tdb := setup(t, services())

	user := factory.CreateUser(map[string]any{})
	hotel := factory.CreateHotel(map[string]any{})
	room := factory.CreateRoom(map[string]any{
		"hotel_id": hotel.ID,
	})

	bookingParams, _ := json.Marshal(types.BookRoomParams{
		From:       time.Now().Add(time.Hour * 1),
		Till:       time.Now().Add(time.Hour * 48),
		NumPersons: 2,
	})

	req := httptest.NewRequest("POST", fmt.Sprintf("/api/v1/bookings/%s", room.ID.Hex()), bytes.NewReader(bookingParams))
	req.Header.Add("Content-Type", "application-json")
	loginAs(user, req)

	res, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}

	filter := &types.BookingQueryParams{UserID: user.ID}
	bookings, _ := tdb.BookingStore.Get(context.Background(), filter)

	assert.Equal(t, 201, res.StatusCode)
	assert.Equal(t, bookings[0].UserID, user.ID)
	assert.Equal(t, bookings[0].RoomID, room.ID)
}

func TestCancelBookedRoom(t *testing.T) {
	app, tdb := setup(t, services())

	user := factory.CreateUser(map[string]any{})
	booking := factory.CreateBooking(map[string]any{
		"user_id": user.ID,
	})

	req := httptest.NewRequest("DELETE", fmt.Sprintf("/api/v1/bookings/%s", booking.ID.Hex()), nil)
	req.Header.Add("Content-Type", "application-json")
	loginAs(user, req)

	res, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}

	fromDatabase, _ := tdb.BookingStore.GetByID(context.Background(), booking.ID.Hex())

	assert.Equal(t, 200, res.StatusCode)
	assert.True(t, fromDatabase.Canceled)
}
