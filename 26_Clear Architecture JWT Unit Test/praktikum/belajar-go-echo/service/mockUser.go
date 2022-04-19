package service

import "belajar-go-echo/model"

type MockUserService struct {
	data []model.User
}

func NewMockUserService() MockUserService {
	return MockUserService{
		data: []model.User{},
	}
}

func (us *MockUserService) Add(user model.User) ([]model.User, error) {
	us.data = append(us.data, user)
	return us.data, nil
}

func (us *MockUserService) Get() ([]model.User, error) {
	return us.data, nil
}
