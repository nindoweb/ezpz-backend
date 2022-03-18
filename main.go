package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"ezpz/config"
	"ezpz/routes"
)

func main() {
	logging()

	addr := fmt.Sprintf("%v:%v", config.AppConfig()["host"], config.AppConfig()["port"])
	if err := routes.Routing().Run(addr); err != nil {
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
