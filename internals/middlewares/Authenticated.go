package middlewares

import (
	"ezpz/internals/jwt"
	"ezpz/internals/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			unAuthorizeError(c)
			return
		}

		if _, err := jwt.VerifyToken(token); err != nil {
			unAuthorizeError(c)
			return
		}
		c.Next()
	}
}

func unAuthorizeError(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, response.JsonResponse{
		Message: "Unauthorized",
	})
	c.Abort()
}
