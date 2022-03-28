package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"ezpz/app/models"
	"ezpz/pkg/common"
	"ezpz/pkg/jwt"
	"ezpz/pkg/redis"
	"ezpz/pkg/response"
	"ezpz/pkg/validations"

	"github.com/gin-gonic/gin"
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

	go func() {
		u.HashPassword(c.Param("password"))
		u.Create()
		u.SendOtp()
	}()

	key := fmt.Sprintf("auth:username:%s", u.Username)
	redis.Set(key, key, time.Minute * 2)

	c.JSON(http.StatusOK, response.JsonResponse{
		Data: map[string]interface{}{
			"username": u.Username,
		},
	})
}

func Login(c *gin.Context) {
	var u models.User
	// todo: check username is required
	u.FindByUser(c.Param("username"))

	if err := common.CheckHash(u.Password, c.Param("password")); err != nil {
		response.Error(c, map[string]string{"username": "username or password is invalid!"})
	}

	key := fmt.Sprintf("auth:username:%s:otp:%s", u.Username, c.Param("otp"))
	redis.Set(key, key, time.Minute * 2)

	go u.SendOtp()

	c.JSON(http.StatusOK, response.JsonResponse{
		Data: map[string]interface{}{
			"username": u.Username,
		},
	})
}

func Logout(c *gin.Context) {
	cookie := http.Cookie{
		Name:   "token",
		MaxAge: -1}
	c.SetCookie(cookie.Name, cookie.Value, cookie.MaxAge, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)

	c.JSON(http.StatusNoContent, response.JsonResponse{Message: "User successfully logout"})
}

func VerifyOtp(c *gin.Context) {
	byteData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		response.InternalServerError(c)
	}
	var data map[string]string
	json.Unmarshal(byteData, &data)

	if data["otp"] == "" {
		response.ValidationError(c, map[string]interface{}{"otp":"otp is required"})
		return
	}

	if data["username"] == "" {
		response.ValidationError(c, map[string]interface{}{"username":"username is required"})
		return
	}

	var u models.User

	key := fmt.Sprintf("auth:username:%s:otp:%s", data["username"], data["otp"])
	value, _ := redis.Get(key)
	if value == ""  {
		response.Error(c, map[string]string{"otp": "otp has been expired"})
		return
	}

	redis.Forget(key)
	u.FindByUser(data["username"])

	token, err := jwt.CreateToken(jwt.NewPayload(u.Username))
	if err != nil {
		response.Error(c, map[string]interface{}{"token": err})
		return
	}

	c.JSON(http.StatusOK, response.JsonResponse{Data: map[string]string{"token": token}})
}
