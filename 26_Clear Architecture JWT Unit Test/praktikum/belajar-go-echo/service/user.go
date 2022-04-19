package service

import (
	"belajar-go-echo/model"
)

type UserService interface {
	Add(user model.User) (model.User, error)
	Get() ([]model.User, error)
}
