package database

import (
	"back-end/internal/config"
	"back-end/internal/constants"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func ConnectToDB() error {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(config.Cfg.MongodbUrl))
	if err != nil {
		return err
	}
	mongoClient = client

	return nil
}

func GetUserCollection() *mongo.Collection {
	return mongoClient.Database(config.Cfg.DatabaseName).Collection(constants.UserCollection)
}
