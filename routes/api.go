package routes

import (
	"ezpz/app/controllers/api"
	"ezpz/app/handlers"

	"github.com/gin-gonic/gin"
)

func RouteApi(r *gin.RouterGroup) {
	v1(r.Group("v1"))
}

func v1(r *gin.RouterGroup) {
	auth := r.Group("auth")
	{
		needAuth := auth.Use(handlers.Auth())
		needAuth.POST("register", api.Register)
		needAuth.POST("login", api.Login)
		needAuth.POST("verify/otp", api.VerifyOtp)

		needGuest := auth.Use(handlers.Guest())
		needGuest.GET("logout", api.Logout)
	}
}
