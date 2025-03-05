package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	// Reading local config
	// viper.AddConfigPath("config/config.yaml")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "hello"
		})
	})

	r.Run("0.0.0.0:8077")

}
