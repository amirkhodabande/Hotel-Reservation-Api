package main

import (
	"context"
	"flag"
	"log"
	"os"

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

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("DB_URI")))
	if err != nil {
		log.Fatal(err)
	}

	database := db.InitDatabase(client, os.Getenv("DB_NAME"))

	address := flag.String("serverPort", os.Getenv("APP_PORT"), "")
	flag.Parse()

	app := app.New(database)

	app.Listen(*address)
}
