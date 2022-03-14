package main

import (
	"ezpz/config"
	"ezpz/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"time"
)

func main() {
	logging()
	routing()
}

func routing() {
	gin.SetMode(config.AppConfig()["env"])
	r := gin.Default()
	r.Use(gin.Logger())
	api := r.Group("api")
	routes.RouteApi(api)

	addr := fmt.Sprintf("%v:%v", config.AppConfig()["host"], config.AppConfig()["port"])
	if err := r.Run(addr); err != nil {
		log.Println(err)
	}
}

func logging() {
	day := time.Now().String()
	fileName := fmt.Sprintf("logs/logs-%s.txt", day)
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
}
