package config

import (
	"belajar-go-echo/model"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	connectionString := os.Getenv("APP_DB_CONNECTION_STRING")
	if connectionString == "" {
		connectionString = "root:@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	}
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	return gorm.Open(mysql.Open(connectionString), &gorm.Config{})
}

func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(
		model.User{},
	)
}
