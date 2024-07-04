package users

import "go.mongodb.org/mongo-driver/mongo"

type service struct {
	coll *mongo.Collection
}

type Service interface {
}

func New() Service {
	return &service{}
}
