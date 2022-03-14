package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type JsonResponse struct {
	Message string
	Data    interface{}
}

func InternalServerError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, JsonResponse{
		Message: "Internal server error",
	})
	c.Abort()
	return
}

func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, JsonResponse{
		Message: "not found",
	})
	c.Abort()
	return
}

func Forbidden(c *gin.Context) {
	c.JSON(http.StatusForbidden, JsonResponse{
		Message: "Forbidden",
	})
	c.Abort()
	return
}

func Error(c *gin.Context, data interface{}) {
	c.JSON(http.StatusInternalServerError, JsonResponse{
		Data: data,
	})
	c.Abort()
	return
}
