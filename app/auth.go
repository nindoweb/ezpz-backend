package app

import (
	"ezpz/internals/common"
	"ezpz/internals/jwt"
	"ezpz/internals/redis"
	"ezpz/internals/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"log"
	"net/http"
)

func Register(c *gin.Context) {
	u := NewUser(false, false)
	//if errs := validations.Struct(u); errs != nil {
	//	response.Error(c, errs)
	//}
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
	id := Create(UserCollection, u)

	key := fmt.Sprintf("auth:user:id:%s", id)
	if err := redis.Set(key, key, 60); err != nil {
		log.Println(err)
		response.InternalServerError(c)
	}

	c.JSON(http.StatusOK, response.JsonResponse{
		Data: map[string]interface{}{
			"user_id": id,
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
	data := Find(UserCollection, "id", c.Param("user_id"))
	if data == nil {
		response.Error(c, map[string]string{"user": "User not found"})
	}

	if err := mapstructure.Decode(data, &u); err != nil {
		log.Println(err)
		response.InternalServerError(c)
	}

	key := fmt.Sprintf("auth:user:id:%s", u.ID.String())
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
