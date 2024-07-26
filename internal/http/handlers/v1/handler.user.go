package v1

import (
	"net/http"
	"time"

	bussiness "back-end/internal/businesses/v1"
	"back-end/internal/constants"
	"back-end/internal/datasource/models"
	"back-end/internal/http/datatransfers/requests"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service bussiness.UserService
}

type UserHandler interface {
	GetUserById(c *gin.Context)
	GetUserByUsername(c *gin.Context)
	SignUp(ctx *gin.Context)
}

func NewUserHandler() UserHandler {
	return &userHandler{
		service: bussiness.NewUserService(),
	}
}

func (h *userHandler) SignUp(ctx *gin.Context) {
	var req requests.UserSignUpRequest
	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, constants.ErrCodeParseRequestFailed, constants.ErrInvalidRequest)
		return
	}
	if err := req.Validate(); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, constants.ErrCodeInvalidRequest, err.Error())
		return
	}
	currentTime := time.Now()
	newUser := models.User{
		Username:  req.Username,
		Email:     req.Email,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Password:  req.Password,
		Phone:     req.Phone,
		Status:    constants.UserStatusRegistered,
		RoleId:    constants.UserRoleReader,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}
	if errCode, err := h.service.Create(ctx, &newUser); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, errCode, err.Error())
		return
	}
}

func (h *userHandler) GetUserById(ctx *gin.Context) {
	var req requests.UserGetUserByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, constants.ErrCodeParseRequestFailed, constants.ErrInvalidRequest)
		return
	}
	userInfo, errCode, err := h.service.GetInfoById(ctx, req.Id)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, errCode, constants.ErrUnknown)
		return
	}
	NewSuccessResponse(ctx, http.StatusOK, userInfo)
}

func (h *userHandler) GetUserByUsername(c *gin.Context) {
	NewErrorResponse(c, http.StatusBadRequest, 1000, "username is required")
}
