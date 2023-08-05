package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DB struct {
	client *mongo.Client
}

func Connect() (DB, error) {
	ctx := context.TODO()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return DB{}, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return DB{}, err
	}

	return DB{client: client}, nil

}

func (db DB) GetUserCollection() *mongo.Collection {
	userCollection := db.client.Database("REST-API-Golang").Collection("user")

	return userCollection
}
