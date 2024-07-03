package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookingQueryParams struct {
	RoomID     primitive.ObjectID `bson:"roomID,omitempty"`
	UserID     primitive.ObjectID `bson:"userID,omitempty"`
	NumPersons int                `bson:"numPersons,omitempty" validate:"omitempty,numeric,min=0"`
	From       time.Time          `bson:"from,omitempty"`
	Till       time.Time          `bson:"till,omitempty"`
	Canceled   bool               `bson:"canceled,omitempty"`
}

type BookRoomParams struct {
	From       time.Time `bson:"from" json:"from" validate:"required,valid_date"`
	Till       time.Time `bson:"till" json:"till" validate:"required,valid_date"`
	NumPersons int       `bson:"numPersons" json:"num_persons" validate:"required"`
}

type UpdateBookingParams struct {
	Canceled bool `bson:"canceled" json:"canceled" validate:"required"`
}

type Booking struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID     primitive.ObjectID `bson:"userID" json:"user_id,omitempty"`
	RoomID     primitive.ObjectID `bson:"roomID" json:"room_id,omitempty"`
	NumPersons int                `bson:"numPersons" json:"num_persons"`
	From       time.Time          `bson:"from" json:"from"`
	Till       time.Time          `bson:"till" json:"till"`
	Canceled   bool               `bson:"canceled" json:"canceled"`
}
