package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type JsonResponse struct {
	Message string `json:"message"`
	Data    interface{} `json:"data"`
}

func InternalServerError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, JsonResponse{
		Message: "Internal server error",
	})
	c.Abort()
}

func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, JsonResponse{
		Message: "not found",
	})
	c.Abort()
}

func Forbidden(c *gin.Context) {
	c.JSON(http.StatusForbidden, JsonResponse{
		Message: "Forbidden",
	})
	c.Abort()
}

func Error(c *gin.Context, data interface{}) {
	c.JSON(http.StatusBadRequest, JsonResponse{
		Data: data,
	})
	c.Abort()
}

func ValidationError(c *gin.Context, data interface{}) {
	c.JSON(http.StatusUnprocessableEntity, JsonResponse{
		Data: data,
	})
	c.Abort()
}