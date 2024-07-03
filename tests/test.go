package tests

import (
	"context"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"hotel.com/app"
	"hotel.com/db"
	"hotel.com/db/factories"
	"hotel.com/types"
)

var factory *factories.Factory

func setup(*testing.T) (*fiber.App, *db.Store) {
	err := godotenv.Load("../.env.testing")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("DB_URI")))
	if err != nil {
		log.Fatal(err)
	}

	client.Database(os.Getenv("DB_NAME")).Drop(ctx)

	tdb := db.InitDatabase(client, os.Getenv("DB_NAME"))

	app := app.New(tdb)

	factory = factories.New(tdb)

	return app, tdb
}

func loginAs(user *types.User, req *http.Request) {
	token := user.CreateToken()
	req.Header.Add("Authorization", token)
}
