package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"hotel.com/types"
)

const roomColl = "rooms"

type RoomStore interface {
	Get(ctx context.Context, filter bson.M) ([]*types.Room, error)
	Insert(ctx context.Context, user *types.Room) (*types.Room, error)
}

type MongoRoomStore struct {
	client *mongo.Client
	coll   *mongo.Collection

	*MongoHotelStore
}

func NewMongoRoomStore(client *mongo.Client, dbname string, hotelStore *MongoHotelStore) *MongoRoomStore {
	return &MongoRoomStore{
		client:          client,
		coll:            client.Database(dbname).Collection(roomColl),
		MongoHotelStore: hotelStore,
	}
}

func (s *MongoRoomStore) Get(ctx context.Context, filter bson.M) ([]*types.Room, error) {
	cur, err := s.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var rooms []*types.Room

	if err := cur.All(ctx, &rooms); err != nil {
		return []*types.Room{}, nil
	}

	return rooms, nil
}

func (s *MongoRoomStore) Insert(ctx context.Context, room *types.Room) (*types.Room, error) {
	res, err := s.coll.InsertOne(ctx, room)
	if err != nil {
		return nil, err
	}

	room.ID = res.InsertedID.(primitive.ObjectID)

	hotel, err := s.MongoHotelStore.GetByID(ctx, room.HotelID.Hex())
	if err != nil {
		return nil, err
	}

	if err := s.MongoHotelStore.UpdateByID(
		ctx,
		room.HotelID.Hex(),
		types.UpdateHotelParams{
			Name:     hotel.Name,
			Location: hotel.Location,
			Rooms:    append(hotel.Rooms, room.ID),
		},
	); err != nil {
		return nil, err
	}

	return room, nil
}
