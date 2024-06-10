package tests

import (
	"context"
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"hotel.com/api"
	"hotel.com/types"
)

func TestGetHotelList(t *testing.T) {
	tdb := setup(t)

	app := fiber.New()
	hotelHandler := api.NewHotelHandler(tdb.HotelStore)
	app.Get("/", hotelHandler.HandleGetHotels)

	hotel, _ := tdb.HotelStore.Insert(context.Background(), &types.Hotel{
		Name:     "TestHotel",
		Location: "Iran",
		Rating:   3,
		Rooms:    []primitive.ObjectID{},
	})

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Add("Content-Type", "application-json")

	res, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}

	encodedRes, _ := io.ReadAll(res.Body)
	exceptedRes, _ := json.Marshal([]types.Hotel{*hotel})

	assert.Equal(t, 200, res.StatusCode)
	assert.JSONEq(t, string(encodedRes), string(exceptedRes))
}
