package v1

import (
	"back-end/internal/datasource/models"
	"back-end/internal/datasource/repositories"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService interface {
	GetInfoUserById(ctx *gin.Context, userId primitive.ObjectID) (*models.User, error)
}

type userService struct {
}

func NewUserService() UserService {
	return &userService{}
}

func (s *userService) GetInfoUserById(ctx *gin.Context, userId primitive.ObjectID) (*models.User, error) {
	userRepo := repositories.NewUser(ctx)
	userInfo, err := userRepo.FindOneByID(userId)
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}
