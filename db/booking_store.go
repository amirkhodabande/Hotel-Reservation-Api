package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"hotel.com/types"
)

const bookingColl = "bookings"

type BookingStore interface {
	Get(ctx context.Context, filter *types.BookingQueryParams) ([]*types.Booking, error)
	Insert(ctx context.Context, booking *types.Booking) (*types.Booking, error)
	GetByID(ctx context.Context, id string) (*types.Booking, error)
	UpdateByID(ctx context.Context, id string, params *types.UpdateBookingParams) error
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

func (s *MongoBookingStore) Get(ctx context.Context, filter *types.BookingQueryParams) ([]*types.Booking, error) {
	opts := &options.FindOptions{}
	opts.SetSkip((filter.GetPage() - 1) * filter.GetLimit())
	opts.SetLimit(filter.GetLimit())

	query := bson.M{}

	if !filter.RoomID.IsZero() {
		query["roomID"] = filter.RoomID
	}
	if !filter.UserID.IsZero() {
		query["userID"] = filter.UserID
	}
	if filter.NumPersons != 0 {
		query["numPersons"] = filter.NumPersons
	}
	if !filter.From.IsZero() {
		query["from"] = bson.M{"$gte": filter.From}
	}
	if !filter.Till.IsZero() {
		query["till"] = bson.M{"$lte": filter.Till}
	}
	query["canceled"] = filter.Canceled

	cur, err := s.coll.Find(ctx, query, opts)
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

func (s *MongoBookingStore) GetByID(ctx context.Context, id string) (*types.Booking, error) {
	bid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var booking types.Booking

	if err := s.coll.FindOne(ctx, bson.M{"_id": bid}).Decode(&booking); err != nil {
		return nil, err
	}

	return &booking, nil
}

func (s *MongoBookingStore) UpdateByID(ctx context.Context, id string, params *types.UpdateBookingParams) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	if _, err := s.coll.UpdateByID(ctx, oid, bson.M{"$set": params}); err != nil {
		return err
	}

	return nil
}
