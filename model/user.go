package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	ID        int64   `gorm:"primary_key" json:"id"`
	Name string   `json:"name,omitempty"`
	Email  string   `json:"email,omitempty"`
	Password  string   `json:"password,omitempty"`
}
