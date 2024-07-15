package v1

import (
	v1 "back-end/internal/businesses/v1"
	"back-end/internal/constants"
	"back-end/internal/http/handlers/serializers"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userHandler struct {
	service v1.UserService
}

type UserHandler interface {
	GetUserByID(c *gin.Context)
	GetUserByUsername(c *gin.Context)
}

func NewUserHandler() UserHandler {
	return &userHandler{
		service: v1.NewUserService(),
	}
}

func (h *userHandler) GetUserByID(ctx *gin.Context) {
	var req serializers.UserGetUserByIDRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, constants.ErrCodeInvalidRequest, constants.ErrInvalidRequest.Error())
		return
	}
	userInfo, err := h.service.GetInfoUserById(ctx, req.Id)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, constants.ErrCodeUnknown, constants.ErrUnknown.Error())
		return
	}
	NewSuccessResponse(ctx, http.StatusOK, userInfo)
}

func (h *userHandler) GetUserByUsername(c *gin.Context) {
	NewErrorResponse(c, http.StatusBadRequest, 1000, "username is required")
	return
}
