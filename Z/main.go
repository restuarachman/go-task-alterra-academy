package main

import (
	"CA/model"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()
	dsn := "root:@tcp(127.0.0.1:3306)/go_latihan?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		e.Logger.Fatal((err))
	}
	db.AutoMigrate(&model.UserRequest{})

	// route

	e.Logger.Fatal(e.Start(":3000"))
}
