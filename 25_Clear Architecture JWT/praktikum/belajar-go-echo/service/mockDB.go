package service

import (
	"belajar-go-echo/model"

	"gorm.io/gorm"
)

type DBUserService struct {
	db gorm.DB
}

func NewDBUserService(Db gorm.DB) DBUserService {
	return DBUserService{
		db: Db,
	}
}

func (us DBUserService) Add(user model.User) (model.User, error) {
	tx := us.db.Save(&user)
	if tx.Error != nil {
		return model.User{}, tx.Error
	}
	return user, nil
}

func (us DBUserService) Get() ([]model.User, error) {
	users := []model.User{}
	tx := us.db.Find(&users)
	if tx.Error != nil {
		return []model.User{}, tx.Error
	}
	return users, nil
}
