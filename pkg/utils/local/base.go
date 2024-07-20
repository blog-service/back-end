package local

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type service struct {
	ctx *gin.Context
}

type Service interface {
	SetUserId(userId primitive.ObjectID)
	GetUserId() primitive.ObjectID
}

func New(ctx *gin.Context) Service {
	return &service{
		ctx: ctx,
	}
}

const (
	userIdKey = "userId"
)
