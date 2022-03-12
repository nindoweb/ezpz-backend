package routes

import (
	"ezpz/app"
	"github.com/gin-gonic/gin"
)

func RouteApi(r *gin.RouterGroup) {
	v1(r.Group("v1"))
}

func v1(r *gin.RouterGroup) {
	auth := r.Group("auth")
	{
		auth.POST("register", app.Register)
		auth.POST("login", app.Login)
		auth.GET("logout", app.Logout)
		auth.POST("verify/otp", app.VerifyOtp)
	}
}
