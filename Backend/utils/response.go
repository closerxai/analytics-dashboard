package utils

import (
	"github.com/gin-gonic/gin"
)

func CustomResponse(c *gin.Context, status int, success bool, message string, data interface{}) {
	c.JSON(status, gin.H{
		"success": success,
		"message": message,
		"data": data,
	})
}
