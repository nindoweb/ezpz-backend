package app

import (
	"log"
	"net/http"

	"ezpz/internals/common"
	"ezpz/internals/response"
	"ezpz/internals/validations"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	u := NewUser(false, false)
	validations.Struct(c, u)
	if c.Param("password") != c.Param("password_confirm") {
		c.JSON(http.StatusUnprocessableEntity, response.JsonResponse{
			Data: map[string]string{"password": "Password and confirm not the same"},
		})
	}

	if err := c.Bind(u); err != nil {
		log.Println(err)
	}
	password, err := common.Hash(c.Param("password"))
	if err != nil {
		log.Println(err)
	}
	u.Password = password

	c.JSON(http.StatusOK, response.JsonResponse{
		Message: "",
		Data:    u.ID,
	})
}

func Login(c *gin.Context) {
	c.JSON(http.StatusOK, "login page")
}

func Logout(c *gin.Context) {
	cookie := http.Cookie{
		Name:   "token",
		MaxAge: -1}
	c.SetCookie(cookie.Name, cookie.Value, cookie.MaxAge, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)

	c.JSON(http.StatusOK, response.JsonResponse{Message: "You are successfully logout"})
}

func VerifyOtp(c *gin.Context) {
	var u User
	token, err := common.CreateToken(u.Username, "secret")
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.JsonResponse{
			Message: "",
			Data:    map[string]string{"token": token},
		})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.JsonResponse{Data: map[string]string{"token": token}})
}
