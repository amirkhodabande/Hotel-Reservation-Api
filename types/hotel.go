package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HotelQueryParams struct {
	Name   string `bson:"name,omitempty" validate:"omitempty,min=2"`
	Rating int    `bson:"rating,omitempty" validate:"omitempty,numeric,min=0"`
	paginationQueryParam
}

type CreateHotelParams struct {
	Name     string               `bson:"name" json:"name" validate:"required,min=2"`
	Location string               `bson:"location" json:"location" validate:"required,min=2"`
	Rooms    []primitive.ObjectID `bson:"rooms" json:"rooms" validate:"nullable"`
}

type UpdateHotelParams struct {
	Name     string               `bson:"name" json:"name" validate:"required,min=2"`
	Location string               `bson:"location" json:"location" validate:"required,min=2"`
	Rooms    []primitive.ObjectID `bson:"rooms" json:"rooms" validate:"required"`
}

type Hotel struct {
	ID       primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string               `bson:"name" json:"name"`
	Location string               `bson:"location" json:"location"`
	Rating   int                  `bson:"rating" json:"rating"`
	Rooms    []primitive.ObjectID `bson:"rooms" json:"rooms"`
}
