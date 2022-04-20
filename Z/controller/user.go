package controller

import (
	"CA/model"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserController struct {
	Db gorm.DB
}

func (uc UserController) NewUserController(db gorm.DB) UserController {
	return UserController{
		Db: db,
	}
}

func (uc UserController) GetAllUsers(c echo.Context) error {
	users := []model.UserRequest{}
	tx := uc.Db.Find(&users)
	if tx.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot get users",
		})
	}

	return c.JSON(http.StatusInternalServerError, map[string]interface{}{
		"message": "success get users",
		"users":   users,
	})
}
