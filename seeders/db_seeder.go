package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"hotel.com/db"
	"hotel.com/types"
)

var client *mongo.Client

func main() {
	var err error
	client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(db.DBuri))
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	seedUsersTable(ctx)
	seedHotelsTable(ctx)
}

func seedUsersTable(ctx context.Context) {
	fmt.Println("Seeding users table...")

	userStore := db.NewMongoUserStore(client, db.DBname)

	user, _ := types.NewUserFromParams(types.CreateUserParams{
		Email:     "test1@gmail.com",
		FirstName: "test1",
		LastName:  "Ltest1",
		Password:  "password",
	})
	anotherUser, _ := types.NewUserFromParams(types.CreateUserParams{
		Email:     "test2@gmail.com",
		FirstName: "test2",
		LastName:  "Ltest2",
		Password:  "password",
	})

	userStore.Insert(ctx, user)
	userStore.Insert(ctx, anotherUser)

	fmt.Println("users Done!")
}

func seedHotelsTable(ctx context.Context) {
	fmt.Println("Seeding hotels table...")

	hotelStore := db.NewMongoHotelStore(client, db.DBname)

	hotel, err := hotelStore.Insert(ctx, &types.Hotel{
		Name:     "TestHotel",
		Location: "Iran",
		Rating:   3,
		Rooms:    []primitive.ObjectID{},
	})
	if err != nil {
		log.Fatal(err)
	}

	roomStore := db.NewMongoRoomStore(client, db.DBname, hotelStore)

	roomStore.Insert(ctx, &types.Room{
		Type:    types.SingleRoomType,
		Price:   70,
		HotelID: hotel.ID,
	})
	roomStore.Insert(ctx, &types.Room{
		Type:    types.DeluxeRoomType,
		Price:   120,
		HotelID: hotel.ID,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("hotels Done!")
}
