package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BookRoomParams struct {
	From       time.Time `bson:"from" json:"from" validate:"required,valid_date"`
	Till       time.Time `bson:"till" json:"till" validate:"required,valid_date"`
	NumPersons int       `bson:"numPersons" json:"num_persons" validate:"required"`
}

type Booking struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID     primitive.ObjectID `bson:"userID" json:"user_id,omitempty"`
	RoomID     primitive.ObjectID `bson:"roomID" json:"room_id,omitempty"`
	NumPersons int                `bson:"numPersons" json:"num_persons"`
	From       time.Time          `bson:"from,omitempty" json:"from"`
	Till       time.Time          `bson:"till,omitempty" json:"till"`
}
