package database

import (
	"context"

	"back-end/internal/config"
	"back-end/internal/constants"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var cfg = config.GetConfig()

var mongoClient *mongo.Client

func ConnectToDB() error {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(cfg.MongodbUrl))
	if err != nil {
		return err
	}
	mongoClient = client

	return nil
}

func GetUserCollection() *mongo.Collection {
	return mongoClient.Database(cfg.DatabaseName).Collection(constants.UserCollection)
}
