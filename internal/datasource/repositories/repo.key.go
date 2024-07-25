package repositories

import (
	"back-end/internal/datasource/database"
	"back-end/internal/datasource/models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type keyRepo struct {
	coll *mongo.Collection
	ctx  context.Context
}

type KeyRepo interface {
	FindOneByID(id primitive.ObjectID, opts ...OptionsQuery) (user *models.User, err error)
}

func NewKey(ctx context.Context) KeyRepo {
	return &keyRepo{
		coll: database.GetUserCollection(),
		ctx:  ctx,
	}
}

func (s *keyRepo) FindOneByID(id primitive.ObjectID, opts ...OptionsQuery) (user *models.User, err error) {
	opt := NewOptions()
	if len(opts) > 0 {
		opt = opts[0]
	}
	findOneOptions := options.FindOne()
	findOneOptions.SetProjection(opt.QueryOnlyField())
	if err = s.coll.FindOne(s.ctx, bson.M{"_id": id}, findOneOptions).Decode(&user); err != nil {
		consoleLog.Error().Err(err).Str("func", "FindOneByID-FindOne.Decode").Msg("keyRepo")
		return nil, err
	}
	return user, nil
}
