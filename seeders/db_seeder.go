package main

import (
	"context"
	"fmt"
	"log"

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

	seedUsersTable()
}

func seedUsersTable() {
	fmt.Println("Seeding users table...")

	userStore := db.NewMongoUserStore(client, db.DBname)

	user := &types.User{
		Email:             "test1@gmail.com",
		FirstName:         "test1",
		LastName:          "Ltest1",
		EncryptedPassword: "testEncrypted",
	}
	anotherUser := &types.User{
		Email:             "test2@gmail.com",
		FirstName:         "test2",
		LastName:          "Ltest2",
		EncryptedPassword: "testEncrypted",
	}

	userStore.InsertUser(context.Background(), user)
	userStore.InsertUser(context.Background(), anotherUser)

	fmt.Println("users Done!")
}
