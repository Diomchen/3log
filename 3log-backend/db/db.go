package db

import (
	"3log-backend/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	CDP := config.C.DB.Postgres
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", CDP.Host, CDP.User, CDP.Password, CDP.Dbname, CDP.Port, CDP.Sslmode)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	} else {
		fmt.Println("Connected to database")
	}
}
