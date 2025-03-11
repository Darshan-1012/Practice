package controllers

import (
	"fmt"
	"jwt/models"
	"jwt/utils"
	_"jwt/database"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	_ "gorm.io/gorm"
	_ "gorm.io/driver/mysql"
)

var validate = validator.New()

func HashPassword(Password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(Password), 14)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func VerifyPassword(UserPassword string, ProvidedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(ProvidedPassword), []byte(UserPassword))
	check := true
	msg := ""
	if err != nil {
		msg = fmt.Sprintf("Email of password is wrong")
		check = false
	}
	return check, msg

}
func RegisterUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var existingUser models.User
		models.DB.Where("Email = ?", user.Email).First(&existingUser)
		if existingUser.ID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user does not exist"})
			return
		}

		token, refreshToken, _ := utils.GenToken(user.ID, user.Email)
		user.token = token
		user.refreshToken = refreshToken
		

		models.DB.Model(&user).Updates(models.User{token: token, refreshToken: refreshToken})

		c.JSON(http.StatusOK, user)
	}
}

func LoginUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginData models.User
		var foundUser models.User

		if err := c.ShouldBindJSON(&loginData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		if err := models.DB.Where("Email=?", loginData.Email).First(&foundUser).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "The email or password is incorrect"})
		}

		if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(loginData.Password)).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Email or Password is wrong"})
		}

		token, refreshToken, _ := utils.GenToken(foundUser.ID, foundUser.Email)
		foundUser.token = token
		foundUser.refreshToken = refreshToken

		models.DB.Model(&foundUser).Updates(models.User{token: token, refreshToken: refreshToken})
		c.JSON(http.StatusOK, foundUser)
	}
}

func DashBoard() gin.HandlerFunc{
	return func(c *gin.Context){
		c.String(http.StatusOK, "Welcome to the dashboard")
	}
}
func LogoutUser() gin.HandlerFunc{
	return func(c *gin.Context){
		c.SetCookie("token", "", -1, "/", "", false, true)
		c.JSON(http.StatusOK, gin.H{"message": "Logout Successful"})
	}
}
