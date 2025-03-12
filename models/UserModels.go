package models

import (
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	gorm.Model
	ID           string `gorm:"primaryKey"`
	Email        string `gorm:"uniqueIndex"`
	Password     string
	Token        string
	RefreshToken string
}
