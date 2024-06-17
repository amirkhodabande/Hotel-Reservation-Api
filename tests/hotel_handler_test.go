package tests

import (
	"context"
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"hotel.com/types"
)

func TestGetHotelList(t *testing.T) {
	app, tdb := setup(t)

	hotel, _ := tdb.HotelStore.Insert(context.Background(), &types.Hotel{
		Name:     "TestHotel",
		Location: "Iran",
		Rating:   3,
		Rooms:    []primitive.ObjectID{},
	})

	req := httptest.NewRequest("GET", "/api/v1/hotels", nil)
	req.Header.Add("Content-Type", "application-json")

	res, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}

	encodedRes, _ := io.ReadAll(res.Body)
	exceptedRes, _ := json.Marshal([]types.Hotel{*hotel})

	assert.Equal(t, 200, res.StatusCode)
	assert.JSONEq(t, string(exceptedRes), string(encodedRes))
}
