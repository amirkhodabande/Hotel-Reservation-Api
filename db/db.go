package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	DBuri  string = "mongodb://localhost:27017"
	DBname string = "hotel-reservation"
)

type Dropper interface {
	Drop(ctx context.Context) error
}

func ToObjectID(id string) primitive.ObjectID {
	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		panic(err)
	}
	return oid
}
