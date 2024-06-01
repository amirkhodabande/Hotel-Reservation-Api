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

type testdb struct {
	db.UserStore
}

func (tdb *testdb) teardown(t *testing.T) {
	if err := tdb.UserStore.Drop(context.TODO()); err != nil {
		t.Fatal(err)
	}
}

func setup(*testing.T) *testdb {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}

	return &testdb{
		UserStore: db.NewMongoUserStore(client, dbname),
	}
}
