package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Reading local config
	// viper.AddConfigPath("config/config.yaml")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})

	r.Run("0.0.0.0:8017")

}
