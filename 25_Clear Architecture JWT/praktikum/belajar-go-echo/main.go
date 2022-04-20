package main

import (
	"belajar-go-echo/config"
	"belajar-go-echo/controller"
	"belajar-go-echo/service"

	"github.com/labstack/echo/v4"
)

func main() {
	db, err := config.ConnectDB()
	if err != nil {
		panic(err)
	}

	err = config.MigrateDB(db)
	if err != nil {
		panic(err)
	}

	us := service.NewDBUserService(*db)
	uc := controller.NewUserController(us)
	app := echo.New()
	app.GET("/users", uc.Get)
	app.POST("/users", uc.Add)
	app.Start(":8080")
}
