package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"hotel.com/db"
	"hotel.com/db/factories"
	"hotel.com/types"
)

var client *mongo.Client

func main() {
	var err error
	client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(db.DBuri))
	if err != nil {
		log.Fatal(err)
	}

	database := db.InitDatabase(client, db.DBname)

	factories := factories.New(database)

	seedUsersTable(factories)
	seedHotelsTable(factories)
}

func seedUsersTable(f *factories.Factory) {
	fmt.Println("Seeding users table...")

	f.CreateUser(map[string]any{
		"email": "test1@gmail.com",
	})
	f.CreateUser(map[string]any{
		"email": "test2@gmail.com",
	})

	fmt.Println("users Done!")
}

func seedHotelsTable(f *factories.Factory) {
	fmt.Println("Seeding hotels table...")

	hotel := f.CreateHotel(map[string]any{})

	f.CreateRoom(map[string]any{
		"type":     types.SingleRoomType,
		"price":    70,
		"hotel_id": hotel.ID,
	})
	f.CreateRoom(map[string]any{
		"type":     types.DeluxeRoomType,
		"price":    500,
		"hotel_id": hotel.ID,
	})

	fmt.Println("hotels Done!")
}
