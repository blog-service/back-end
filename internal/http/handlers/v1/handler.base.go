package v1

import (
	"back-end/internal/http/datatransfers/responses"
	"github.com/gin-gonic/gin"
)

func NewSuccessResponse(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, responses.BaseResponse{
		Success: true,
		Data:    data,
	})
}

func NewErrorResponse(c *gin.Context, statusCode, errorCode int, mess string) {
	c.JSON(statusCode, responses.BaseResponse{
		Success:   false,
		Message:   mess,
		ErrorCode: errorCode,
	})

}
