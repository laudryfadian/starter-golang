package helpers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/laudryfadian/starter-golang/internal/entity"
)

func SuccessResponse(c *gin.Context, message string, statusCode int, data interface{}) {
	c.JSON(statusCode, entity.SuccessResponse{
		StatusCode: statusCode,
		Message:    message,
		DateTime:   time.Now().Format(time.RFC3339),
		Data:       data,
	})
}

func ErrorResponse(c *gin.Context, message string, statusCode int) {
	c.JSON(statusCode, entity.ErrorResponse{
		StatusCode: statusCode,
		Message:    message,
		DateTime:   time.Now().Format(time.RFC3339),
		Data:       nil,
	})
}
