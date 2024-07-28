package v1

import (
	"net/http"

	bussiness "back-end/internal/businesses/v1"
	"back-end/internal/constants"
	"back-end/internal/http/datatransfers/requests"
	"back-end/internal/http/datatransfers/responses"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service bussiness.UserService
}

type UserHandler interface {
	GetUserById(c *gin.Context)
	SignUp(ctx *gin.Context)
	SignIn(ctx *gin.Context)
}

func NewUserHandler() UserHandler {
	return &userHandler{
		service: bussiness.NewUserService(),
	}
}

func (h *userHandler) SignUp(ctx *gin.Context) {
	var req requests.UserSignUpRequest
	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, &responses.ErrorResponse{
			ErrorCode: constants.ErrCodeParseRequestFailed,
			Message:   constants.ErrInvalidRequest,
		})
		return
	}
	if err := req.Validate(); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, &responses.ErrorResponse{
			ErrorCode: constants.ErrCodeInvalidRequest,
			Message:   err.Error(),
		})
		return
	}
	if errCode, err := h.service.Create(ctx, &req); err != nil {
		if errCode == constants.ErrCodeDuplicateData {
			NewErrorResponse(ctx, http.StatusConflict, &responses.ErrorResponse{
				ErrorCode: errCode,
				Message:   err.Error(),
			})
			return
		}
		NewErrorResponse(ctx, http.StatusInternalServerError, &responses.ErrorResponse{
			ErrorCode: errCode,
			Message:   err.Error(),
		})
		return
	}
	NewSuccessResponse(ctx, http.StatusCreated, nil)
}

func (h *userHandler) SignIn(ctx *gin.Context) {
	var req requests.UserSignInRequest
	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, &responses.ErrorResponse{
			ErrorCode: constants.ErrCodeParseRequestFailed,
			Message:   constants.ErrInvalidRequest,
		})
		return
	}
	if err := req.Validate(); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, &responses.ErrorResponse{
			ErrorCode: constants.ErrCodeInvalidRequest,
			Message:   err.Error(),
		})
		return
	}

	userId, errCode, err := h.service.CheckUser(ctx, &req)
	if err != nil {
		if errCode == constants.ErrCodeUserNotFound || errCode == constants.ErrCodeUserInvalidPassword {
			NewErrorResponse(ctx, http.StatusUnauthorized, &responses.ErrorResponse{
				ErrorCode: errCode,
				Message:   err.Error(),
			})
			return
		}
		NewErrorResponse(ctx, http.StatusInternalServerError, &responses.ErrorResponse{
			ErrorCode: errCode,
			Message:   err.Error(),
		})
		return
	}

	accessToken, errCode, err := h.service.RegisToken(ctx, userId)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, &responses.ErrorResponse{
			ErrorCode: errCode,
			Message:   err.Error(),
		})
		return
	}

	NewSuccessResponse(ctx, http.StatusOK, &responses.SuccessResponse{
		Data: responses.UserSignInResponse{
			Token: accessToken,
		},
	})
}

func (h *userHandler) GetUserById(ctx *gin.Context) {
	var req requests.UserGetUserByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, &responses.ErrorResponse{
			ErrorCode: constants.ErrCodeParseRequestFailed,
			Message:   constants.ErrInvalidRequest,
		})
		return
	}
	userInfo, errCode, err := h.service.GetInfoById(ctx, req.Id)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, &responses.ErrorResponse{
			ErrorCode: errCode,
			Message:   err.Error(),
		})
		return
	}
	NewSuccessResponse(ctx, http.StatusOK, &responses.SuccessResponse{
		Data: userInfo,
	})
}
