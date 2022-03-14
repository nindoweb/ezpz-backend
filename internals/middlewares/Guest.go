package middlewares

import (
	"ezpz/internals/jwt"
	"ezpz/internals/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Guest() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token != "" {
			if _, err := jwt.VerifyToken(token); err == nil {
				guestResponseError(c)
				return
			}
			guestResponseError(c)
			return
		}
		c.Next()
	}
}

func guestResponseError(c *gin.Context) {
	c.JSON(http.StatusForbidden, response.JsonResponse{
		Message: "Forbidden",
	})
	c.Abort()
}
