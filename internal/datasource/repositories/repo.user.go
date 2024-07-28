package repositories

import (
	"context"
	"errors"

	"back-end/internal/constants"
	"back-end/internal/datasource/database"
	"back-end/internal/datasource/models"
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
	InsertOne(user models.User) (newUser *models.User, errCode int, err error)
	FindOneByID(id primitive.ObjectID, opts ...OptionsQuery) (user *models.User, errCode int, err error)
	FindOneByUsername(username string, opts ...OptionsQuery) (user *models.User, errCode int, err error)
}

func NewUser(ctx context.Context) UserRepo {
	return &userRepo{
		coll: database.GetUserCollection(),
		ctx:  ctx,
	}
}

func (s *userRepo) FindOneByID(id primitive.ObjectID, opts ...OptionsQuery) (user *models.User, errCode int, err error) {
	opt := NewOptions()
	if len(opts) > 0 {
		opt = opts[0]
	}
	findOneOptions := options.FindOne()
	findOneOptions.SetProjection(opt.QueryOnlyField())
	if err = s.coll.FindOne(s.ctx, bson.M{"_id": id}, findOneOptions).Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, constants.ErrCodeUserNotFound, err
		}
		consoleLog.Error().Err(err).Str("func", "FindOneByID-FindOne.Decode").Msg("userRepo")
		return nil, constants.ErrCodeUnknown, err
	}
	return user, constants.ErrCodeNoErr, nil
}

func (s *userRepo) InsertOne(user models.User) (newUser *models.User, errCode int, err error) {
	result, err := s.coll.InsertOne(s.ctx, user)
	if err != nil {
		consoleLog.Error().Err(err).Str("func", "InsertOne-InsertOne").Msg("userRepo")
		return nil, constants.ErrCodeUnknown, err
	}
	user.Id = result.InsertedID.(primitive.ObjectID)
	return &user, constants.ErrCodeNoErr, nil
}

func (s *userRepo) FindOneByUsername(username string, opts ...OptionsQuery) (user *models.User, errCode int, err error) {
	opt := NewOptions()
	if len(opts) > 0 {
		opt = opts[0]
	}
	findOneOptions := options.FindOne()
	findOneOptions.SetProjection(opt.QueryOnlyField())
	if err = s.coll.FindOne(s.ctx, bson.M{"username": username}, findOneOptions).Decode(&user); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, constants.ErrCodeUserNotFound, err
		}
		consoleLog.Error().Err(err).Str("func", "FindOneByUsername-FindOne.Decode").Msg("userRepo")
		return nil, constants.ErrCodeUnknown, err
	}
	return user, constants.ErrCodeNoErr, nil
}
