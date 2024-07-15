package repositories

import (
	"back-end/internal/datasource/database"
	"back-end/internal/datasource/models"
	"back-end/pkg/logger"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var consoleLog = logger.ConsoleLog()

type service struct {
	coll *mongo.Collection
	ctx  context.Context
}

type Service interface {
	FindOneByID(id primitive.ObjectID) (user *models.User, err error)
}

func NewUser(ctx context.Context) Service {
	return &service{
		coll: database.GetUserCollection(),
		ctx:  ctx,
	}
}

func (s *service) FindOneByID(id primitive.ObjectID) (user *models.User, err error) {
	if err = s.coll.FindOne(s.ctx, bson.D{{"_id", id}}).Decode(&user); err != nil {
		consoleLog.Error().Err(err).Str("func", "FindOne.Decode").Msg("userRepo")
		return nil, err
	}
	return user, nil
}
