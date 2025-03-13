// package models

// import (
// 	_ "gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// var DB *gorm.DB

// //	type User struct {
// //		gorm.Model
// //		ID           string `json:"ID"`
// //		Email        string `json:"Email"`
// //		Password     string `json:"Password"`
// //		Token        string `json:"Token"`
// //		RefreshToken string `json:"RefreshToken"`
// //	}
// type User struct {
// 	gorm.Model
// 	Username string `gorm:"unique;not null" json:"username"`
// 	Email    string `gorm:"unique;not null" json:"email"`
// 	Password string `json:"password"`
// }

package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username"`
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `json:"password"`
}
