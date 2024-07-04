package user

import (
	userRepo "back-end/internal/datasource/repositories/users"
)

type service struct {
	repo *userRepo.Service
}

type Service interface {
}

func New() Service {
	return &service{}
}
