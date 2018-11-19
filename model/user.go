package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type UserView struct {
	gorm.Model
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

func (UserView) TableName() string {
	return "user_view"
}

func (User) TableName() string {
	return "users"
}
