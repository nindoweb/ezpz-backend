package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"ezpz/routes"

	"github.com/spf13/viper"
)

func main() {
	logging()

	viper.SetConfigName("config") 
	viper.SetConfigType("yaml")
	dir, _ := os.Getwd()
	viper.AddConfigPath(dir)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("please cp config.example.yaml to config.yaml and setup your configureation")
		} else {
			panic(err)
		}
	}

	addr := fmt.Sprintf("%v:%v", viper.GetString("HOST"), viper.GetString("PORT"))
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
