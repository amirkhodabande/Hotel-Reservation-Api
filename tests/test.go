package tests

import (
	"context"
	"log"
	"testing"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"hotel.com/app"
	"hotel.com/db"
)

const (
	dburi  = "mongodb://localhost:27017"
	dbname = "test-hotel-reservation"
)

func setup(*testing.T) (*fiber.App, *db.Store) {
	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}

	client.Database(dbname).Drop(ctx)

	tdb := db.InitDatabase(client, dbname)

	app := app.New(tdb)

	return app, tdb
}
