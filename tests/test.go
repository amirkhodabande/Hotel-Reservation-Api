package tests

import (
	"context"
	"log"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"hotel.com/app"
	"hotel.com/db"
	"hotel.com/types"
)

const (
	dburi  = "mongodb://localhost:27017"
	dbname = "test-hotel-reservation"
)

func setup(*testing.T) (*fiber.App, *db.Store) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

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

func loginAs(user *types.User, req *http.Request) {
	token := user.CreateToken()
	req.Header.Add("Authorization", token)
}
