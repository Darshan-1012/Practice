// package models

// import (
// 	_ "gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// var DB *gorm.DB

//	type User struct {
//		gorm.Model
//		ID           uint   `gorm:"primaryKey;autoIncrement"`
//		Email        string `gorm:"uniqueIndex"`
//		Password     string
//		Token        string
//		RefreshToken string
//	}
package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username"`
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `json:"password"`
}
