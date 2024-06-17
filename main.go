package main

import (
	"context"
	"flag"
	"log"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"hotel.com/app"
	"hotel.com/db"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(db.DBuri))
	if err != nil {
		log.Fatal(err)
	}

	database := db.InitDatabase(client, db.DBname)

	address := flag.String("serverPort", ":5000", "")
	flag.Parse()

	app := app.New(database)

	app.Listen(*address)
}
