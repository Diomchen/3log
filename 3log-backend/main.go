package main

import (
	"3log-backend/config"
	"3log-backend/db"
	"3log-backend/router"
)

func main() {
	// Reading local config
	// viper.AddConfigPath("config/config.yaml")
	config.Init()

	db.InitDB()
	router.Include(router.PingRouter)
	r := router.Init()

	r.Run("0.0.0.0:8017")

}
