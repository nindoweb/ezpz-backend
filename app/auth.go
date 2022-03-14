package app

import (
	"ezpz/internals/redis"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"ezpz/internals/common"
	"ezpz/internals/jwt"
	"ezpz/internals/response"
	"ezpz/internals/validations"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	u := NewUser(false, false)
	validations.Struct(c, u)
	if c.Param("password") != c.Param("password_confirm") {
		response.Error(c, map[string]string{"password": "Password and confirm not the same"})
	}

	if err := c.Bind(u); err != nil {
		log.Println(err)
	}
	password, err := common.Hash(c.Param("password"))
	if err != nil {
		log.Println(err)
	}
	u.Password = password

	key := fmt.Sprintf("auth:user:id:%s", strconv.FormatUint(uint64(u.ID), 10))
	value := strconv.FormatUint(uint64(u.ID), 10)
	if err := redis.Set(key, value, 60); err != nil {
		log.Println(err)
		response.InternalServerError(c)
	}

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

	key := fmt.Sprintf("auth:user:id:%s", strconv.FormatUint(uint64(u.ID), 10))
	_, err := redis.Get(key)
	if err != nil {
		response.Error(c, map[string]string{"otp": "otp has been expired"})
	}

	token, err := jwt.CreateToken(jwt.NewPayload(u.Username))
	if err != nil {
		response.Error(c, map[string]string{"token": token})
	}

	c.JSON(http.StatusOK, response.JsonResponse{Data: map[string]string{"token": token}})
}
