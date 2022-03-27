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
		needGuest := auth.Use(handlers.Guest())
		needGuest.POST("register", api.Register)
		needGuest.POST("login", api.Login)
		needGuest.POST("verify/otp", api.VerifyOtp)

		needAuth := auth.Use(handlers.Auth())
		needAuth.GET("logout", api.Logout)
	}
}
