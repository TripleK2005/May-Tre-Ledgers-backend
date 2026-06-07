package response

import "github.com/gin-gonic/gin"

func Success[T any](c *gin.Context, status int, message string, data T) {
	c.JSON(status, Response[T]{
		Success: true,
		Message: message,
		Data:    data,
	})
}
