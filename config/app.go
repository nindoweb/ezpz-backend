package config

import "github.com/gin-gonic/gin"

func AppConfig() map[string]string {
	return map[string]string{
		"env":            gin.DebugMode,
		"debug":          "true",
		"host":           "127.0.0.1",
		"port":           "8000",
		"jwt_secret":     "secret",
		"jwt_expire":     "5",
		"redis_host":     "127.0.0.1",
		"redis_port":     "6379",
		"redis_password": "",
		"redis_db":       "0",
	}
}
