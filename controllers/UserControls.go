package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// import (
// 	"fmt"
// 	_ "jwt/database"
// 	"jwt/models"
// 	"jwt/utils"
// 	"log"
// 	"net/http"
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// 	"github.com/go-playground/validator/v10"
// 	"golang.org/x/crypto/bcrypt"
// 	_ "gorm.io/driver/mysql"
// 	_ "gorm.io/gorm"
// )

// var validate = validator.New()

// func HashPassword(Password string) string {
// 	bytes, err := bcrypt.GenerateFromPassword([]byte(Password), 14)
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	return string(bytes)
// }

// func VerifyPassword(UserPassword string, ProvidedPassword string) (bool, string) {
// 	err := bcrypt.CompareHashAndPassword([]byte(ProvidedPassword), []byte(UserPassword))
// 	check := true
// 	msg := ""
// 	if err != nil {
// 		msg = fmt.Sprintf("Email of password is wrong")
// 		check = false
// 	}
// 	return check, msg

// }
// func RegisterUser() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var user models.User

// 		// Bind JSON request
// 		if err := c.ShouldBindJSON(&user); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		// Check if the user already exists
// 		var existingUser models.User
// 		err := models.DB.Where("email = ?", user.Email).First(&existingUser).Error
// 		if err == nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
// 			return
// 		}

// 		// Hash the password
// 		user.Password = HashPassword(user.Password)

// 		// Log user data before insertion
// 		log.Println("User Data Before Insert:", user)

// 		// Insert the user into the database
// 		if err := models.DB.Create(&user).Error; err != nil {
// 			log.Println("Database Insert Error:", err)
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
// 			return
// 		}

// 		// Ensure ID is generated (if ID is `uint`, check `user.ID == 0`)
// 		if user.ID == 0 {
// 			log.Println("User ID is not set after creation")
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user ID"})
// 			return
// 		}

// 		// Respond with success
// 		c.JSON(http.StatusOK, gin.H{
// 			"message": "User registered successfully",
// 			"user_id": user.ID,
// 			"email":   user.Email,
// 		})
// 	}
// }

// func RegisterUser() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var user models.User

// 		// Bind JSON request
// 		if err := c.ShouldBindJSON(&user); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		// Check if the user already exists
// 		var existingUser models.User
// 		if err := models.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
// 			return
// 		}

// 		// Hash the password before saving
// 		user.Password = HashPassword(user.Password)

// 		// Create the user in the database
// 		if err := models.DB.Create(&user).Error; err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
// 			return
// 		}

// 		// Debugging: Check if ID is assigned
// 		if user.ID == 0 {
// 			log.Println("User ID is not set after creation")
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user ID"})
// 			return
// 		}

// 		// Respond with user details
// 		c.JSON(http.StatusOK, gin.H{
// 			"message": "User registered successfully",
// 			"user_id": user.ID,
// 			"email":   user.Email,
// 		})
// 	}
// }

// func RegisterUser() gin.HandlerFunc {
// 	return func(c *gin.Context) {

// 		var user models.User
// 		if err := c.ShouldBindJSON(&user); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		// var existingUser models.User
// 		// models.DB.Where("Email = ?", user.Email).First(&existingUser)
// 		// if existingUser.ID == "" {
// 		// 	c.JSON(http.StatusBadRequest, gin.H{"error": "user does not exist"})
// 		// 	return
// 		// }
// 		// Check if the user already exists
// 		var existingUser models.User
// 		err := models.DB.Where("email = ?", user.Email).First(&existingUser).Error
// 		if err == nil {
// 			// If a record is found, return an error (user already exists)
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
// 			return
// 		}
// 		user.Password = HashPassword(user.Password)
// 		// token, refreshToken, _ := utils.GenToken(user.ID, user.Email)
// 		// // user.token = token
// 		// user.Token = token
// 		// // user.refreshToken = refreshToken
// 		// user.RefreshToken = refreshToken

// 		// models.DB.Model(&user).Updates(models.User{token: token, refreshToken: refreshToken})
// 		if err := models.DB.Create(&user).Error; err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
// 			return
// 		}
// 		// if user.ID == "0" {
// 		// 	log.Println("User ID is not set after creation")
// 		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user ID"})
// 		// 	return
// 		// }
// 		// c.JSON(http.StatusOK, user)
// 		// user.Password = ""
// 		c.JSON(http.StatusOK, gin.H{
// 			"message": "User registered successfully",
// 			"user_id": user.ID,
// 			"email":   user.Email,
// 		})
// 	}
// }

// func LoginUser() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var loginData models.User
// 		var foundUser models.User

// 		// Bind JSON request
// 		if err := c.ShouldBindJSON(&loginData); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}

// 		// Find user by email
// 		if err := models.DB.Where("email = ?", loginData.Email).First(&foundUser).Error; err != nil {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "The email or password is incorrect"})
// 			return
// 		}

// 		// Check password
// 		if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(loginData.Password)); err != nil {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Email or Password is wrong"})
// 			return
// 		}

// 		// Convert uint ID to string for GenToken
// 		userIDStr := strconv.FormatUint(uint64(foundUser.ID), 10)

// 		// Generate tokens
// 		token, refreshToken, err := utils.GenToken(userIDStr, foundUser.Email)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
// 			return
// 		}

// 		// Save tokens to foundUser
// 		foundUser.Token = token
// 		foundUser.RefreshToken = refreshToken

// 		// Update the user record with new tokens
// 		models.DB.Model(&foundUser).Updates(map[string]interface{}{
// 			"token":         token,
// 			"refresh_token": refreshToken,
// 		})

// 		// Respond with user details & token
// 		c.JSON(http.StatusOK, gin.H{
// 			"message":       "Login successful",
// 			"user_id":       foundUser.ID,
// 			"email":         foundUser.Email,
// 			"token":         token,
// 			"refresh_token": refreshToken,
// 		})
// 	}
// }

// func LoginUser() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		var loginData models.User
// 		var foundUser models.User

// 		if err := c.ShouldBindJSON(&loginData); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		}

// 		if err := models.DB.Where("Email=?", loginData.Email).First(&foundUser).Error; err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "The email or password is incorrect"})
// 		}

// 		if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(loginData.Password)).Error; err != nil {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Email or Password is wrong"})
// 		}

// 		token, refreshToken, _ := utils.GenToken(foundUser.ID, foundUser.Email)
// 		// foundUser.token = token
// 		// foundUser.refreshToken = refreshToken
// 		foundUser.Token = token
// 		foundUser.RefreshToken = refreshToken

// 		// models.DB.Model(&foundUser).Updates(models.User{token: token, refreshToken: refreshToken})
// 		models.DB.Save(&foundUser)
// 		c.JSON(http.StatusOK, foundUser)
// 	}
// }

//	func DashBoard() gin.HandlerFunc {
//		return func(c *gin.Context) {
//			c.String(http.StatusOK, "Welcome to the dashboard")
//		}
//	}
func LogoutUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.SetCookie("token", "", -1, "/", "", false, true)
		c.JSON(http.StatusOK, gin.H{"message": "Logout Successful"})
	}
}
