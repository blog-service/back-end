package local

import "go.mongodb.org/mongo-driver/bson/primitive"

func (s *service) SetUserId(userId primitive.ObjectID) {
	s.ctx.Set(userIdKey, userId)
}

func (s *service) GetUserId() primitive.ObjectID {
	return s.ctx.Value(userIdKey).(primitive.ObjectID)
}
