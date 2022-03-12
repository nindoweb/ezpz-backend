package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	c.JSON(http.StatusOK, "register page")
}

func Login(c *gin.Context) {
	c.JSON(http.StatusOK, "login page")
}

func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, "logout page")
}
