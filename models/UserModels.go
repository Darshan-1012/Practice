package models

import (
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	gorm.Model
	ID           string `json:"ID"`
	Email        string `json:"Email"`
	Password     string `json:"Password"`
	Token        string `json:"Token"`
	RefreshToken string `json:"RefreshToken"`
}
