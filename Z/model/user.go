package model

type UserRequest struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `json:"username"`
	Password string `json:"password"`
}
