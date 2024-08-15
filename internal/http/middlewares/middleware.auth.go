package middlewares

import (
	"net/http"
	"strings"

	"back-end/internal/config"
	"back-end/internal/constants"
	"back-end/internal/datasource/repositories"
	"back-end/internal/http/datatransfers/responses"
	handler "back-end/internal/http/handlers/v1"
	"back-end/pkg/jwt"
	"back-end/pkg/utils/local"
	"github.com/gin-gonic/gin"
)

func ValidateAccess() gin.HandlerFunc {
	jwtService := jwt.NewJwtService(config.GetConfig().PrivateKeyPath, config.GetConfig().PublicKeyPath)

	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			handler.NewErrorResponse(c, http.StatusUnauthorized, &responses.ErrorResponse{
				ErrorCode: constants.ErrCodeMissingToken,
				Message:   constants.ErrMissingToken,
			})
			return
		}
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 {
			handler.NewErrorResponse(c, http.StatusUnauthorized, &responses.ErrorResponse{
				ErrorCode: constants.ErrCodeMissingToken,
				Message:   constants.ErrMissingToken,
			})
			return
		}
		accessToken := bearerToken[1]
		claims, err := jwtService.ValidateToken(accessToken)
		if err != nil {
			handler.NewErrorResponse(c, http.StatusUnauthorized, &responses.ErrorResponse{
				ErrorCode: constants.ErrCodeWrongToken,
				Message:   constants.ErrWrongToken,
			})
			return
		}
		if !claims.Data.IsAccess {
			handler.NewErrorResponse(c, http.StatusUnauthorized, &responses.ErrorResponse{
				ErrorCode: constants.ErrCodeWrongToken,
				Message:   constants.ErrWrongToken,
			})
			return
		}

		optionQuery := repositories.NewOptions()
		optionQuery.SetOnlyFields("user_id")
		key, _, err := repositories.NewKey(c).FindOneByTokenId(claims.ID, optionQuery)
		if err != nil {
			handler.NewErrorResponse(c, http.StatusUnauthorized, &responses.ErrorResponse{
				ErrorCode: constants.ErrCodeWrongToken,
				Message:   constants.ErrWrongToken,
			})
			return
		}

		local.New(c).SetUserId(key.UserId)
		c.Next()
	}
}

func ValidateRefresh() gin.HandlerFunc {
	jwtService := jwt.NewJwtService(config.GetConfig().PrivateKeyPath, config.GetConfig().PublicKeyPath)

	return func(c *gin.Context) {
		refreshToken, err := c.Cookie("refresh_token")
		if err != nil {
			handler.NewErrorResponse(c, http.StatusUnauthorized, &responses.ErrorResponse{
				ErrorCode: constants.ErrCodeMissingToken,
				Message:   constants.ErrMissingToken,
			})
			return
		}
		claims, err := jwtService.ValidateToken(refreshToken)
		if err != nil {
			handler.NewErrorResponse(c, http.StatusUnauthorized, &responses.ErrorResponse{
				ErrorCode: constants.ErrCodeWrongToken,
				Message:   constants.ErrWrongToken,
			})
			return
		}
		if claims.Data.IsAccess {
			handler.NewErrorResponse(c, http.StatusUnauthorized, &responses.ErrorResponse{
				ErrorCode: constants.ErrCodeWrongToken,
				Message:   constants.ErrWrongToken,
			})
			return
		}

		optionQuery := repositories.NewOptions()
		optionQuery.SetOnlyFields("user_id")
		key, _, err := repositories.NewKey(c).FindOneByTokenId(claims.ID, optionQuery)
		if err != nil {
			handler.NewErrorResponse(c, http.StatusUnauthorized, &responses.ErrorResponse{
				ErrorCode: constants.ErrCodeWrongToken,
				Message:   constants.ErrWrongToken,
			})
			return
		}

		local.New(c).SetUserId(key.UserId)
		c.Next()
	}
}
