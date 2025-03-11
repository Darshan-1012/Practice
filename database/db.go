package database

import (
	"log"
	"os"
	"jwt/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(){

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	MysqlDb := os.Getenv("MYSQL_DSN")

	db1, err := gorm.Open(mysql.Open(MysqlDb), &gorm.Config{})

	if err != nil {
		log.Panic(err)
	}

	db1.AutoMigrate(&models.User{})

	DB = db1
	models.DB = db1
}
