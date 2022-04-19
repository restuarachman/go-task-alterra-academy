package controller

import (
	"belajar-go-echo/model"
	"belajar-go-echo/service"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	us service.UserService
}

func NewUserController(us service.UserService) UserController {
	return UserController{
		us: us,
	}
}

func (uc UserController) Get(c echo.Context) error {
	users, err := uc.us.Get()
	if err != nil {
		return c.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(200, echo.Map{
		"data": users,
	})
}

func (uc UserController) Add(c echo.Context) error {
	newUser := model.User{}
	c.Bind(&newUser)
	user, err := uc.us.Add(newUser)
	if err != nil {
		return c.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(200, echo.Map{
		"data": user,
	})
}
