package utils

import (
	"github.com/gin-gonic/gin"
)

// ErrorResponse sends a standard error response with a given message.
func ErrorResponse(c *gin.Context, message string) {
	c.JSON(400, gin.H{
		"error": message,
	})
}

// SuccessResponse sends a standard success response with the provided data.
func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"data": data,
	})
}
