package v1

import (
	"back-end/internal/constants"
	"back-end/internal/datasource/repositories"
	"back-end/internal/http/datatransfers/responses"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type KeyService interface {
}

type keyService struct {
}

func NewKeyService() KeyService {
	return &keyService{}
}

func (s *keyService) GetInfoById(ctx *gin.Context, userId primitive.ObjectID) (*responses.UserGetInfoResponse, int, error) {
	queryOptions := repositories.NewOptions()
	queryOptions.SetOnlyFields("username", "email", "first_name", "last_name", "phone", "role", "created_at", "created_by", "updated_at", "updated_by")
	userInfo, errCode, err := repositories.NewUser(ctx).FindOneByID(userId, queryOptions)
	if err != nil {
		return nil, errCode, err
	}
	user := &responses.UserGetInfoResponse{
		Username:  userInfo.Username,
		Email:     userInfo.Email,
		FirstName: userInfo.FirstName,
		LastName:  userInfo.LastName,
		Phone:     userInfo.Phone,
		RoleId:    userInfo.RoleId,
		RoleName:  constants.MapUserRoles[userInfo.RoleId],
		CreatedAt: userInfo.CreatedAt,
		CreatedBy: userInfo.CreatedBy,
		UpdatedAt: userInfo.UpdatedAt,
		UpdatedBy: userInfo.UpdatedBy,
	}
	return user, constants.ErrCodeNoErr, nil
}
