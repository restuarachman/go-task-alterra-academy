package controller

import (
	"belajar-go-echo/model"
	"belajar-go-echo/service"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestUserControllerAdd(t *testing.T) {
	e := echo.New()

	newUserJson, _ := json.Marshal(map[string]interface{}{
		"email": "restu@gmail.com",
	})
	req := httptest.NewRequest(http.MethodGet, "/", bytes.NewBuffer(newUserJson))
	req.Header.Set("Content-Type", "application/jsonn")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/")

	us := service.NewMockUserService()
	uc := NewUserController(us)
	uc.Add(c)

	users, err := us.Get()
	if err != nil {
		t.Error(err)
	}
	if len(users) != 1 {
		t.Errorf("Expecting len(users) to be 1, get %d", len(users))
	}
	if users[0].Email != "restu@gmail.com" {
		t.Errorf("Expectiong emails to be restu@gmail.com, get %s", users[0].Email)
	}
}

func TestUserControllerGet(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set("Content-Type", "application/jsonn")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/")

	us := service.NewMockUserService()
	us.Add(model.User{Email: "restu@gmail.com"})
	us.Add(model.User{Email: "aditya@gmail.com"})
	us.Add(model.User{Email: "rachman@gmail.com"})

	uc := NewUserController(us)
	uc.Get(c)

	var users []model.User
	if err := json.Unmarshal(rec.Body.Bytes(), &users); err != nil {
		t.Error("unmarhalling returned person failed")
	}
	if len(users) != 3 {
		t.Errorf("expecting len(users) is 3, get %d", len(users))
	}
	if users[0].Email != "restu@gmail.com" {
		t.Errorf("expection users[0].Email is restu@gmail.com, get %s", users[0].Email)
	}
	if users[1].Email != "aditya@gmail.com" {
		t.Errorf("expection users[0].Email is aditya@gmail.com, get %s", users[1].Email)
	}
	if users[2].Email != "rachman@gmail.com" {
		t.Errorf("expection users[0].Email is rachman@gmail.com, get %s", users[2].Email)
	}
}
