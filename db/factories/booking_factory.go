package factories

import (
	"context"
	"maps"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"hotel.com/types"
)

func (f *Factory) CreateBooking(data map[string]any) *types.Booking {
	var userID, roomID primitive.ObjectID

	if value, exists := data["user_id"]; exists {
		userID = value.(primitive.ObjectID)
	} else {
		user := f.CreateUser(map[string]any{})
		userID = user.ID
	}

	if value, exists := data["room_id"]; exists {
		roomID = value.(primitive.ObjectID)
	} else {
		room := f.CreateRoom(map[string]any{})
		roomID = room.ID
	}

	sample := map[string]any{
		"user_id":     userID,
		"room_id":     roomID,
		"num_persons": 2,
		"from":        time.Now().Add(time.Hour * 1),
		"till":        time.Now().Add(time.Hour * 48),
		"canceled":    false,
	}
	maps.Copy(sample, data)

	booking := &types.Booking{}
	transcode(sample, booking)

	ctx := context.Background()
	f.BookingStore.Insert(ctx, booking)

	return booking
}
