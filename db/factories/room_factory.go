package factories

import (
	"context"
	"maps"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"hotel.com/types"
)

func (f *Factory) CreateRoom(data map[string]any) *types.Room {
	var hotelID primitive.ObjectID

	if value, exists := data["hotel_id"]; exists {
		hotelID = value.(primitive.ObjectID)
	} else {
		hotel := f.CreateHotel(map[string]any{})
		hotelID = hotel.ID
	}

	sample := map[string]any{
		"type":     types.SingleRoomType,
		"price":    70,
		"hotel_id": hotelID,
	}
	maps.Copy(sample, data)

	room := &types.Room{}
	transcode(sample, room)

	ctx := context.Background()
	f.RoomStore.Insert(ctx, room)

	return room
}
