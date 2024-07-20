package repositories

import (
	"back-end/internal/datasource/database"
	"back-end/internal/datasource/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type keyRepo struct {
	coll *mongo.Collection
	ctx  context.Context
}

type KeyRepo interface {
	FindOneByID(id primitive.ObjectID) (user *models.User, err error)
}

func NewKey(ctx context.Context) KeyRepo {
	return &keyRepo{
		coll: database.GetUserCollection(),
		ctx:  ctx,
	}
}

func (s *keyRepo) FindOneByID(id primitive.ObjectID) (user *models.User, err error) {
	if err = s.coll.FindOne(s.ctx, bson.M{"_id": id}).Decode(&user); err != nil {
		consoleLog.Error().Err(err).Str("func", "FindOne.Decode").Msg("userRepo")
		return nil, err
	}
	return user, nil
}
