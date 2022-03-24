package handlers

import (
	"ezpz/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			errorMessage(c, "Unauthorized")
			return
		}

		if _, err := jwt.VerifyToken(token); err != nil {
			errorMessage(c, "Unauthorized")
			return
		}
		c.Next()
	}
}



