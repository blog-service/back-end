package v1

import (
	"back-end/internal/http/datatransfers/responses"
	"github.com/gin-gonic/gin"
)

func NewSuccessResponse(c *gin.Context, statusCode int, response *responses.SuccessResponse) {
	if response == nil {
		response = &responses.SuccessResponse{}
	}
	response.Success = true
	c.JSON(statusCode, response)
}

func NewErrorResponse(c *gin.Context, statusCode int, response *responses.ErrorResponse) {
	if response == nil {
		response = &responses.ErrorResponse{}
	}
	response.Success = false
	c.JSON(statusCode, response)
}
