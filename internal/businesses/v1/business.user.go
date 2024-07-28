package v1

import (
	"errors"
	"time"

	"back-end/internal/constants"
	"back-end/internal/datasource/models"
	"back-end/internal/datasource/repositories"
	"back-end/internal/http/datatransfers/requests"
	"back-end/pkg/hash"
	"back-end/pkg/jwt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService interface {
	GetInfoById(ctx *gin.Context, userId primitive.ObjectID) (user *models.User, errCode int, err error)
	Create(ctx *gin.Context, user *requests.UserSignUpRequest) (errCode int, err error)
	CheckUser(ctx *gin.Context, userReq *requests.UserSignInRequest) (userId primitive.ObjectID, errCode int, err error)
	RegisToken(ctx *gin.Context, userId primitive.ObjectID) (token string, errCode int, err error)
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

func (s *userService) CheckUser(ctx *gin.Context, userReq *requests.UserSignInRequest) (userId primitive.ObjectID, errCode int, err error) {
	userRepo := repositories.NewUser(ctx)
	queryOptions := repositories.NewOptions()
	queryOptions.SetOnlyFields("_id", "password")
	user, errCode, err := userRepo.FindOneByUsername(userReq.Username, queryOptions)
	if err != nil {
		if errCode == constants.ErrCodeUserNotFound {
			return primitive.NilObjectID, errCode, errors.New("invalid username")
		}
		return primitive.NilObjectID, errCode, err
	}
	if !hash.New().CheckPasswordHash(userReq.Password, user.Password) {
		return primitive.NilObjectID, constants.ErrCodeUserInvalidPassword, errors.New("invalid password")
	}
	return user.Id, constants.ErrCodeNoErr, nil
}

func (s *userService) RegisToken(ctx *gin.Context, userId primitive.ObjectID) (token string, errCode int, err error) {
	keyRepo := repositories.NewKey(ctx)
	queryOptions := repositories.NewOptions()
	queryOptions.SetOnlyFields("_id")
	key, errCode, err := keyRepo.FindOneByUserId(userId, queryOptions)
	if err != nil && errCode != constants.ErrCodeUserKeyNotFound {
		return "", errCode, err
	}

	if key != nil {
		if errCode, err = keyRepo.DeleteOneById(key.Id); err != nil {
			return "", errCode, err
		}
	}

	jwtService := jwt.NewJwtService(cfg.PrivateKeyPath, cfg.PublicKeyPath)
	tokenId := primitive.NewObjectID().Hex()
	accessToken, err := jwtService.GenerateToken(tokenId, false, cfg.AccessTokenExpired)
	if err != nil {
		return "", constants.ErrCodeUserGenerateTokenFailed, err
	}
	refreshToken, err := jwtService.GenerateToken(tokenId, true, cfg.RefreshTokenExpired)
	if err != nil {
		return "", constants.ErrCodeUserGenerateTokenFailed, err
	}

	if _, errCode, err = keyRepo.InsertOne(models.Key{
		UserId:  userId,
		TokenId: tokenId,
	}); err != nil {
		return "", errCode, err
	}

	ctx.SetCookie("refresh_token", refreshToken, int(cfg.RefreshTokenExpired.Seconds()), "/", "/", true, true)
	return accessToken, constants.ErrCodeNoErr, nil
}
