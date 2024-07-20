package middlewares

import (
	"back-end/internal/config"
	"back-end/internal/constants"
	handler "back-end/internal/http/handlers/v1"
	"back-end/pkg/jwt"
	"back-end/pkg/utils/local"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	jwtService := jwt.NewJwtService(config.GetConfig().PrivateKeyPath, config.GetConfig().PublicKeyPath)

	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			handler.NewErrorResponse(c, http.StatusUnauthorized, constants.ErrCodeMissingToken, constants.ErrMissingToken)
			return
		}
		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 {
			handler.NewErrorResponse(c, http.StatusUnauthorized, constants.ErrCodeMissingToken, constants.ErrMissingToken)
			return
		}
		accessToken := bearerToken[1]
		claims, err := jwtService.ValidateToken(accessToken)
		if err != nil {
			handler.NewErrorResponse(c, http.StatusUnauthorized, constants.ErrCodeWrongToken, constants.ErrWrongToken)
			return
		}
		if !claims.Data.IsAccess {
			handler.NewErrorResponse(c, http.StatusUnauthorized, constants.ErrCodeWrongToken, constants.ErrWrongToken)
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
