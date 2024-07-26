package v1

import (
	"back-end/internal/constants"
	"back-end/internal/datasource/models"
	"back-end/internal/datasource/repositories"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService interface {
	GetInfoById(ctx *gin.Context, userId primitive.ObjectID) (user *models.User, errCode int, err error)
	Create(ctx *gin.Context, user *models.User) (errCode int, err error)
}

type userService struct {
}

func NewUserService() UserService {
	return &userService{}
}

func (s *userService) GetInfoById(ctx *gin.Context, userId primitive.ObjectID) (user *models.User, errCode int, err error) {
	userInfo, errCode, err := repositories.NewUser(ctx).FindOneByID(userId)
	if err != nil {
		return nil, errCode, err
	}
	return userInfo, constants.ErrCodeNoErr, nil
}

func (s *userService) Create(ctx *gin.Context, user *models.User) (errCode int, err error) {
	if _, errCode, err = repositories.NewUser(ctx).InsertOne(user); err != nil {
		return errCode, err
	}
	return constants.ErrCodeNoErr, nil
}
