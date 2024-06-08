package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type RoomType int

const (
	SingleRoomType RoomType = iota
	DoubleRoomType
	SeaSideRoomType
	DeluxeRoomType
)

type CreateRoomParams struct {
	Type    RoomType             `bson:"type" json:"type" validate:"required,min=2"`
	Price   int                `bson:"price" json:"price" validate:"required,min=2"`
	HotelID primitive.ObjectID `bson:"hotelID" json:"hotel_id" validate:"nullable"`
}

type UpdateRoomParams struct {
	Type    RoomType             `bson:"type" json:"type" validate:"required,min=2"`
	Price   int                `bson:"price" json:"price" validate:"required"`
	HotelID primitive.ObjectID `bson:"hotelID" json:"hotel_id" validate:"required"`
}

type Room struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Type    RoomType           `bson:"type" json:"type"`
	Price   int                `bson:"price" json:"price"`
	HotelID primitive.ObjectID `bson:"hotelID" json:"hotel_id"`
}
