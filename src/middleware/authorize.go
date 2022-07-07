package middleware

import (
	"api-2/src/helper"

	"github.com/gin-gonic/gin"
)

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")

		_, errVal := helper.ValidateToken(token)

		if errVal != nil {
			c.AbortWithStatusJSON(401, gin.H{
				"status":  "failed",
				"message": "unauthorized",
			})
		} else {
			c.Next()
		}
	}
}
