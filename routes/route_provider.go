package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Routing() *gin.Engine {
	gin.SetMode(viper.GetString("ENV"))
	r := gin.Default()
	r.Use(gin.Logger())
	api := r.Group("api")
	RouteApi(api)

	return r
}
