package middlewares

import (
	"net/http"
	"strings"

	"back-end/internal/config"
	"back-end/internal/constants"
	"back-end/internal/http/datatransfers/responses"
	handler "back-end/internal/http/handlers/v1"
	"back-end/pkg/jwt"
	"back-end/pkg/utils/local"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AuthMiddleware() gin.HandlerFunc {
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
		localService := local.New(c)
		userId, err := primitive.ObjectIDFromHex(claims.Issuer)
		if err != nil {
			return
		}
		localService.SetUserId(userId)
		c.Next()
	}
}
