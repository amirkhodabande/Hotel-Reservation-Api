package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"hotel.com/types"
)

const hotelColl = "hotels"

type HotelStore interface {
	Get(ctx context.Context, queryParams *types.HotelQueryParams) ([]*types.Hotel, error)
	Insert(ctx context.Context, hotel *types.Hotel) (*types.Hotel, error)
	GetByID(ctx context.Context, id string) (*types.Hotel, error)
	UpdateByID(ctx context.Context, id string, params types.UpdateHotelParams) error
}

type MongoHotelStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoHotelStore(client *mongo.Client, dbname string) *MongoHotelStore {
	return &MongoHotelStore{
		client: client,
		coll:   client.Database(dbname).Collection(hotelColl),
	}
}

func (s *MongoHotelStore) Get(ctx context.Context, queryParams *types.HotelQueryParams) ([]*types.Hotel, error) {
	opts := &options.FindOptions{}
	opts.SetSkip((queryParams.GetPage() - 1) * queryParams.GetLimit())
	opts.SetLimit(queryParams.GetLimit())

	cur, err := s.coll.Find(ctx, queryParams, opts)

	if err != nil {
		return nil, err
	}

	var hotels []*types.Hotel

	if err := cur.All(ctx, &hotels); err != nil {
		return []*types.Hotel{}, nil
	}

	return hotels, nil
}

func (s *MongoHotelStore) Insert(ctx context.Context, hotel *types.Hotel) (*types.Hotel, error) {
	res, err := s.coll.InsertOne(ctx, hotel)

	if err != nil {
		return nil, err
	}

	hotel.ID = res.InsertedID.(primitive.ObjectID)
	return hotel, nil
}

func (s *MongoHotelStore) GetByID(ctx context.Context, id string) (*types.Hotel, error) {
	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	var hotel types.Hotel

	if err := s.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&hotel); err != nil {
		return nil, err
	}

	return &hotel, nil
}

func (s *MongoHotelStore) UpdateByID(ctx context.Context, id string, params types.UpdateHotelParams) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	if _, err = s.coll.UpdateByID(ctx, oid, bson.D{{"$set", params}}); err != nil {
		return err
	}

	return nil
}
