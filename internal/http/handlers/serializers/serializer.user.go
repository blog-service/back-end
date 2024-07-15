package serializers

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserGetUserByIDRequest struct {
	Id primitive.ObjectID `query:"id"`
}
