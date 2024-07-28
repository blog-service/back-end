package database

import (
	"context"

	"back-end/internal/config"
	"back-end/internal/constants"
	"back-end/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
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

	userIndex()
	keyIndex()

	return nil
}

func GetUserCollection() *mongo.Collection {
	return mongoClient.Database(cfg.DatabaseName).Collection(constants.UserCollection)
}

func GetKeyCollection() *mongo.Collection {
	return mongoClient.Database(cfg.DatabaseName).Collection(constants.KeyCollection)
}

func userIndex() {
	collIndex := mongoClient.Database(cfg.DatabaseName).Collection(constants.UserCollection).Indexes()
	ctxDrop, cancelDrop := context.WithTimeout(context.Background(), cfg.MongodbTimeout)
	defer cancelDrop()
	_, _ = collIndex.DropAll(ctxDrop)
	ctx, cancel := context.WithTimeout(context.Background(), cfg.MongodbTimeout)
	defer cancel()
	if _, err := collIndex.CreateMany(ctx, []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "username", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys:    bson.D{{Key: "email", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys:    bson.D{{Key: "phone", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys:    bson.D{{Key: "username", Value: 1}, {Key: "role", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys:    bson.D{{Key: "username", Value: 1}, {Key: "status", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	}); err != nil {
		logger.ConsoleLog().Fatal().Err(err).Msg("userIndex")
	}
}

func keyIndex() {
	collIndex := mongoClient.Database(cfg.DatabaseName).Collection(constants.KeyCollection).Indexes()
	ctxDrop, cancelDrop := context.WithTimeout(context.Background(), cfg.MongodbTimeout)
	defer cancelDrop()
	_, _ = collIndex.DropAll(ctxDrop)
	ctx, cancel := context.WithTimeout(context.Background(), cfg.MongodbTimeout)
	defer cancel()
	if _, err := collIndex.CreateMany(ctx, []mongo.IndexModel{
		{
			Keys:    bson.D{{Key: "user_id", Value: 1}, {Key: "token_id", Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	}); err != nil {
		logger.ConsoleLog().Fatal().Err(err).Msg("keyIndex")
	}
}
