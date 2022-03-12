package routes

import (
	"ezpz/app/controllers/api"
	"github.com/gin-gonic/gin"
)

func RouteApi(r *gin.RouterGroup) {
	v1(r.Group("v1"))
}

func v1(r *gin.RouterGroup) {
	auth := r.Group("auth")
	{
		auth.POST("register", api.Register)
		auth.POST("login", api.Login)
		auth.GET("logout", api.Logout)
	}
}
