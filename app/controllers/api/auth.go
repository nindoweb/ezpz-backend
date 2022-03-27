package api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"ezpz/app/models"
	"ezpz/pkg/common"
	"ezpz/pkg/jwt"
	"ezpz/pkg/redis"
	"ezpz/pkg/response"
	"ezpz/pkg/validations"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

func Register(c *gin.Context) {
	u := models.NewUser(false, false)

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
	go models.Create(models.USER_COLLECTION, u)

	key := fmt.Sprintf("auth:username:%s", u.Username)
	redis.Set(key, key, time.Minute * 2)

	// if err := notification.SendMail("register", []string{u.Email}, "6565"); err != nil {
	// 	log.Println(err)
	// 	panic(err)
	// }

	c.JSON(http.StatusOK, response.JsonResponse{
		Data: map[string]interface{}{
			"username": u.Username,
		},
	})
}

func Login(c *gin.Context) {
	var u models.User
	data := models.Find(models.USER_COLLECTION, "username", u.Username)
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
	var u models.User

	key := fmt.Sprintf("auth:user:id:%s", u.Username)
	value, _ := redis.Get(key)
	if value == ""  {
		response.Error(c, map[string]string{"otp": "otp has been expired"})
	}

	data := models.Find(models.USER_COLLECTION, "username", u.Username)
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
