package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"hotel.com/types"
)

const bookingColl = "bookings"

type BookingStore interface {
	Get(ctx context.Context, filter bson.M) ([]*types.Booking, error)
	Insert(ctx context.Context, booking *types.Booking) (*types.Booking, error)
}

type MongoBookingStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoBookingStore(client *mongo.Client, dbname string) *MongoBookingStore {
	return &MongoBookingStore{
		client: client,
		coll:   client.Database(dbname).Collection(bookingColl),
	}
}

func (s *MongoBookingStore) Get(ctx context.Context, filter bson.M) ([]*types.Booking, error) {
	cur, err := s.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var bookings []*types.Booking

	if err := cur.All(ctx, &bookings); err != nil {
		return []*types.Booking{}, nil
	}

	return bookings, nil
}

func (s *MongoBookingStore) Insert(ctx context.Context, booking *types.Booking) (*types.Booking, error) {
	res, err := s.coll.InsertOne(ctx, booking)

	if err != nil {
		return nil, err
	}

	booking.ID = res.InsertedID.(primitive.ObjectID)
	return booking, nil
}
