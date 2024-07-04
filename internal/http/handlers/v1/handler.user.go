package v1

import (
	userService "back-end/internal/businesses/v1/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userHandler struct {
	service userService.Service
}

type UserHandler interface {
	GetUserByID(c *gin.Context)
	GetUserByUsername(c *gin.Context)
}

func New() UserHandler {
	return &userHandler{
		service: userService.New(),
	}
}

func (h *userHandler) GetUserByID(c *gin.Context) {
	NewSuccessResponse(c, http.StatusOK, nil)
	return
}

func (h *userHandler) GetUserByUsername(c *gin.Context) {
	NewErrorResponse(c, http.StatusBadRequest, 1000, "username is required")
	return
}
