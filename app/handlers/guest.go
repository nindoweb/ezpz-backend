package handlers

import (
	"ezpz/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func Guest() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != "" {
			if _, err := jwt.VerifyToken(token); err == nil {
				errorMessage(c, "Forbidden")
				return
			}
			errorMessage(c, "Forbidden")
			return
		}
		c.Next()
	}
}
