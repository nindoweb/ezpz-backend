package routes

import (
	"ezpz/config"
	"github.com/gin-gonic/gin"
)

func Routing() *gin.Engine {
	gin.SetMode(config.AppConfig()["env"])
	r := gin.Default()
	r.Use(gin.Logger())
	api := r.Group("api")
	RouteApi(api)

	return r
}
