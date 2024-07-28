package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Key struct {
	Id      primitive.ObjectID `bson:"_id,omitempty"`
	UserId  primitive.ObjectID `bson:"user_id"`
	TokenId string             `bson:"token_id"`
}
