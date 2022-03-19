package app

import (
	"fmt"
	"log"
	"net/http"

	"ezpz/internals/common"
	"ezpz/internals/jwt"
	"ezpz/internals/redis"
	"ezpz/internals/response"
	"ezpz/internals/validations"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

func Register(c *gin.Context) {
	u := NewUser(false, false)

	if verr := validations.ValidateStruct(c, u); verr != nil {
		response.ValidationError(c, verr)
		return
	}

	if c.Param("password") != c.Param("password_confirm") {
		response.Error(c, map[string]string{"password": "Password and confirm not the same"})
	}

	password, err := common.Hash(c.Param("password"))
	if err != nil {
		log.Println(err)
	}

	u.Password = password
	Create(UserCollection, u)

	key := fmt.Sprintf("auth:username:%s", u.Username)
	if err := redis.Set(key, key, 60); err != nil {
		log.Println(err)
		response.InternalServerError(c)
		return
	}

	c.JSON(http.StatusOK, response.JsonResponse{
		Data: map[string]interface{}{
			"username": u.Username,
		},
	})
}

func Login(c *gin.Context) {
	var u User
	data := Find(UserCollection, "username", c.Param("username"))
	if data == nil {
		response.Error(c, map[string]string{"username": "username or password is invalid!"})
	}

	if err := mapstructure.Decode(data, &u); err != nil {
		log.Println(err)
		response.InternalServerError(c)
	}

	if err := common.CheckHash(u.Password, c.Param("password")); err != nil {
		response.Error(c, map[string]string{"username": "username or password is invalid!"})
	}

	c.JSON(http.StatusOK, response.JsonResponse{
		Data: map[string]interface{}{
			"user_id": u.ID.String(),
		},
	})
}

func Logout(c *gin.Context) {
	cookie := http.Cookie{
		Name:   "token",
		MaxAge: -1}
	c.SetCookie(cookie.Name, cookie.Value, cookie.MaxAge, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)

	c.JSON(http.StatusOK, response.JsonResponse{Message: "User successfully logout"})
}

func VerifyOtp(c *gin.Context) {
	var u User

	key := fmt.Sprintf("auth:user:id:%s", u.Username)
	_, err := redis.Get(key)
	if err != nil {
		response.Error(c, map[string]string{"otp": "otp has been expired"})
	}

	data := Find(UserCollection, "id", c.Param("user_id"))
	if data == nil {
		response.Error(c, map[string]string{"user": "User not found"})
	}

	if err := mapstructure.Decode(data, &u); err != nil {
		log.Println(err)
		response.InternalServerError(c)
	}

	token, err := jwt.CreateToken(jwt.NewPayload(u.Username))
	if err != nil {
		response.Error(c, map[string]string{"token": token})
	}

	c.JSON(http.StatusOK, response.JsonResponse{Data: map[string]string{"token": token}})
}
