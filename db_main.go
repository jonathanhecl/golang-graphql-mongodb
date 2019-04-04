package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoDB struct {
	session *mongo.Client
	users   *mongo.Collection
	recipes *mongo.Collection
}

// Connect to MongoDB
func connectDB() mongoDB {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(config.mongoUri))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	return mongoDB{
		session: client,
		recipes: client.Database(config.mongoDb).Collection("recipes"),
		users:   client.Database(config.mongoDb).Collection("users"),
	}
}

// Disconnect to MongoDB
func (db mongoDB) closeDB() {
	err := db.session.Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
