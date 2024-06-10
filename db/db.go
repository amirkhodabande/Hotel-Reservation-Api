package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DBuri  string = "mongodb://localhost:27017"
	DBname string = "hotel-reservation"
)

type Store struct {
	UserStore
	HotelStore
	RoomStore
}

func InitDatabase(client *mongo.Client, dbName string) *Store {
	hotelStore := NewMongoHotelStore(client, dbName)

	return &Store{
		UserStore:  NewMongoUserStore(client, dbName),
		HotelStore: hotelStore,
		RoomStore:  NewMongoRoomStore(client, dbName, hotelStore),
	}
}

func ToObjectID(id string) primitive.ObjectID {
	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		panic(err)
	}
	return oid
}
