package factories

import (
	"context"
	"maps"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"hotel.com/types"
)

func (f *Factory) CreateHotel(data map[string]any) *types.Hotel {
	sample := map[string]any{
		"name":     "TestHotel",
		"location": "Iran",
		"rating":   3,
		"rooms":    []primitive.ObjectID{},
	}
	maps.Copy(sample, data)

	hotel := &types.Hotel{}
	transcode(sample, hotel)

	ctx := context.Background()
	f.HotelStore.Insert(ctx, hotel)

	return hotel
}
