package handlers

import (
	"net/http"

	"ezpz/pkg/response"

	"github.com/gin-gonic/gin"
)

func errorMessage(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, response.JsonResponse{
		Message: message,
	})
	c.Abort()
}