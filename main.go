package main

import (
	"ezpz/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"time"
)

func main() {
	logging()
	routeing()
}

func routeing() {
	r := gin.Default()
	api := r.Group("api")
	routes.RouteApi(api)

	if err := r.Run(); err != nil {
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
