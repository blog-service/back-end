package v1

import (
	"back-end/internal/datasource/models"
	"back-end/internal/datasource/repositories"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService interface {
	GetInfoById(ctx *gin.Context, userId primitive.ObjectID) (*models.User, error)
	Create(ctx *gin.Context, user *models.User) error
}

type userService struct {
}

func NewUserService() UserService {
	return &userService{}
}

func (s *userService) GetInfoById(ctx *gin.Context, userId primitive.ObjectID) (*models.User, error) {
	userInfo, err := repositories.NewUser(ctx).FindOneByID(userId)
	if err != nil {
		return nil, err
	}
	return userInfo, nil
}

func (s *userService) Create(ctx *gin.Context, user *models.User) error {
	if _, err := repositories.NewUser(ctx).InsertOne(user); err != nil {
		return err
	}
	return nil
}
