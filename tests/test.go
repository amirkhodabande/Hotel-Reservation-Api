package tests

import (
	"context"
	"log"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"hotel.com/db"
)

const (
	dburi  = "mongodb://localhost:27017"
	dbname = "test-hotel-reservation"
)

func setup(*testing.T) *db.Store {
	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}

	client.Database(dbname).Drop(ctx)

	return db.InitDatabase(client, dbname)
}
