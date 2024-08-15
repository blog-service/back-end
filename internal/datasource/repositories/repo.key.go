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

type keyRepo struct {
	coll *mongo.Collection
	ctx  context.Context
}

type KeyRepo interface {
	FindOneByUserIdAndTokenId(userId primitive.ObjectID, tokenId string, opts ...OptionsQuery) (key *models.Key, errCode int, err error)
	FindOneByUserId(userId primitive.ObjectID, opts ...OptionsQuery) (key *models.Key, errCode int, err error)
	InsertOne(key models.Key) (newKey *models.Key, errCode int, err error)
	DeleteOneByUserId(userId primitive.ObjectID) (errCode int, err error)
	DeleteOneById(id primitive.ObjectID) (errCode int, err error)
	FindOneByTokenId(tokenId string, opts ...OptionsQuery) (key *models.Key, errCode int, err error)
}

func NewKey(ctx context.Context) KeyRepo {
	return &keyRepo{
		coll: database.GetKeyCollection(),
		ctx:  ctx,
	}
}

func (s *keyRepo) FindOneByUserIdAndTokenId(userId primitive.ObjectID, tokenId string, opts ...OptionsQuery) (key *models.Key, errCode int, err error) {
	opt := NewOptions()
	if len(opts) > 0 {
		opt = opts[0]
	}
	findOneOptions := options.FindOne()
	findOneOptions.SetProjection(opt.QueryOnlyField())
	if err = s.coll.FindOne(s.ctx, bson.M{"user_id": userId, "token_id": tokenId}, findOneOptions).Decode(&key); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, constants.ErrCodeUserKeyNotFound, err
		}
		consoleLog.Error().Err(err).Str("func", "FindOneByUserIdAndTokenId-FindOne.Decode").Msg("keyRepo")
		return nil, constants.ErrCodeUnknown, err
	}
	return key, constants.ErrCodeNoErr, nil
}

func (s *keyRepo) FindOneByUserId(userId primitive.ObjectID, opts ...OptionsQuery) (key *models.Key, errCode int, err error) {
	opt := NewOptions()
	if len(opts) > 0 {
		opt = opts[0]
	}
	findOneOptions := options.FindOne()
	findOneOptions.SetProjection(opt.QueryOnlyField())
	if err = s.coll.FindOne(s.ctx, bson.M{"user_id": userId}, findOneOptions).Decode(&key); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, constants.ErrCodeUserKeyNotFound, err
		}
		consoleLog.Error().Err(err).Str("func", "FindOneByUserId-FindOne.Decode").Msg("keyRepo")
		return nil, constants.ErrCodeUnknown, err
	}
	return key, constants.ErrCodeNoErr, nil
}

func (s *keyRepo) InsertOne(key models.Key) (newKey *models.Key, errCode int, err error) {
	result, err := s.coll.InsertOne(s.ctx, key)
	if err != nil {
		consoleLog.Error().Err(err).Str("func", "InsertOne-InsertOne").Msg("keyRepo")
		return nil, constants.ErrCodeUnknown, err
	}
	key.Id = result.InsertedID.(primitive.ObjectID)
	return &key, constants.ErrCodeNoErr, nil
}

func (s *keyRepo) DeleteOneByUserId(userId primitive.ObjectID) (errCode int, err error) {
	if _, err = s.coll.DeleteOne(s.ctx, bson.M{"user_id": userId}); err != nil {
		consoleLog.Error().Err(err).Str("func", "DeleteOneByUserId-DeleteOne").Msg("keyRepo")
		return constants.ErrCodeUnknown, err
	}
	return constants.ErrCodeNoErr, nil
}

func (s *keyRepo) DeleteOneById(id primitive.ObjectID) (errCode int, err error) {
	if _, err = s.coll.DeleteOne(s.ctx, bson.M{"_id": id}); err != nil {
		consoleLog.Error().Err(err).Str("func", "DeleteOneById-DeleteOne").Msg("keyRepo")
		return constants.ErrCodeUnknown, err
	}
	return constants.ErrCodeNoErr, nil
}

func (s *keyRepo) FindOneByTokenId(tokenId string, opts ...OptionsQuery) (key *models.Key, errCode int, err error) {
	opt := NewOptions()
	if len(opts) > 0 {
		opt = opts[0]
	}
	findOneOptions := options.FindOne()
	findOneOptions.SetProjection(opt.QueryOnlyField())
	if err = s.coll.FindOne(s.ctx, bson.M{"token_id": tokenId}, findOneOptions).Decode(&key); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, constants.ErrCodeUserKeyNotFound, err
		}
		consoleLog.Error().Err(err).Str("func", "FindOneByTokenId-FindOne.Decode").Msg("keyRepo")
		return nil, constants.ErrCodeUnknown, err
	}
	return key, constants.ErrCodeNoErr, nil
}
