package middlewares

import (
	"net/http"
	"strings"

	"back-end/internal/constants"
	stringUtils "back-end/pkg/utils/string-utils"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", constants.AllowOrigin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", constants.AllowCredential)
		c.Writer.Header().Set("Access-Control-Allow-Headers", constants.AllowHeader)
		c.Writer.Header().Set("Access-Control-Allow-Methods", constants.AllowMethods)
		c.Writer.Header().Set("Access-Control-Max-Age", constants.MaxAge)

		if !stringUtils.New().IsArrayContains(strings.Split(constants.AllowMethods, ", "), c.Request.Method) {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden with CORS policy"})
			return
		}

		for key := range c.Request.Header {
			if !stringUtils.New().IsArrayContains(strings.Split(constants.AllowHeader, ", "), key) {
				c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden with CORS policy"})
				return
			}
		}

		c.Next()
	}
}
