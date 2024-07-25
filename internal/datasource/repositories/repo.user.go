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

type userRepo struct {
	coll *mongo.Collection
	ctx  context.Context
}

type UserRepo interface {
	FindOneByID(id primitive.ObjectID, opts ...OptionsQuery) (user *models.User, err error)
	InsertOne(user *models.User) (newUser *models.User, err error)
}

func NewUser(ctx context.Context) UserRepo {
	return &userRepo{
		coll: database.GetUserCollection(),
		ctx:  ctx,
	}
}

func (s *userRepo) FindOneByID(id primitive.ObjectID, opts ...OptionsQuery) (user *models.User, err error) {
	opt := NewOptions()
	if len(opts) > 0 {
		opt = opts[0]
	}
	findOneOptions := options.FindOne()
	findOneOptions.SetProjection(opt.QueryOnlyField())
	if err = s.coll.FindOne(s.ctx, bson.M{"_id": id}, findOneOptions).Decode(&user); err != nil {
		consoleLog.Error().Err(err).Str("func", "FindOneByID-FindOne.Decode").Msg("userRepo")
		return nil, err
	}
	return user, nil
}

func (s *userRepo) InsertOne(user *models.User) (newUser *models.User, err error) {
	result, err := s.coll.InsertOne(s.ctx, user)
	if err != nil {
		consoleLog.Error().Err(err).Str("func", "InsertOne-InsertOne").Msg("userRepo")
		return nil, err
	}
	user.Id = result.InsertedID.(primitive.ObjectID)
	return user, nil
}
