package middlewares

import "github.com/gin-gonic/gin"

func Guest() gin.HandlerFunc {
	return func(c *gin.Context) {
		// before request
		c.Next()
		// after request
	}
}
