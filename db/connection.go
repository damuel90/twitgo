package db

import (
	"context"
	"log"

	"github.com/damuel90/twitgo/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DBConnection = connectionDB()

func connectionDB() *mongo.Client {
	URI := utils.Getenv("MONGO_URI", "mongodb://localhost:27017/twitgo")
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(URI))
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	if err = client.Ping(context.TODO(), nil); err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Database has connected successfully")
	return client
}

func CheckConnection() bool {
	if err := DBConnection.Ping(context.TODO(), nil); err != nil {
		log.Fatal(err.Error())
		return false
	}
	return true
}
