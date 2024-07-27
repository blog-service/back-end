package v1

import (
	"time"

	"back-end/internal/constants"
	"back-end/internal/datasource/models"
	"back-end/internal/datasource/repositories"
	"back-end/internal/http/datatransfers/requests"
	"back-end/pkg/hash"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService interface {
	GetInfoById(ctx *gin.Context, userId primitive.ObjectID) (user *models.User, errCode int, err error)
	Create(ctx *gin.Context, user *requests.UserSignUpRequest) (errCode int, err error)
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

func (s *userService) Create(ctx *gin.Context, user *requests.UserSignUpRequest) (errCode int, err error) {
	passHash, err := hash.New().HashPassword(user.Password)
	if err != nil {
		return constants.ErrCodeHashPassFailed, err
	}
	user.Password = passHash

	currentTime := time.Now()
	if _, errCode, err = repositories.NewUser(ctx).InsertOne(models.User{
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Password:  passHash,
		Phone:     user.Phone,
		Status:    constants.UserStatusRegistered,
		RoleId:    constants.UserRoleReader,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}); err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return constants.ErrCodeDuplicateData, err
		}
		return errCode, err
	}
	return constants.ErrCodeNoErr, nil
}
