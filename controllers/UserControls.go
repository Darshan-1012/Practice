// package controllers

// import (
// 	_ "jwt/database"
// 	"jwt/models"
// 	"jwt/utils"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/go-playground/validator/v10"
// 	"golang.org/x/crypto/bcrypt"
// 	_ "gorm.io/driver/mysql"
// 	_ "gorm.io/gorm"
// )

// var validate = validator.New()

// func HashPassword(Password string) (string, error) {
// 	bytes, err := bcrypt.GenerateFromPassword([]byte(Password), 14)
// 	// if err != nil {
// 	// 	log.Panic(err)
// 	// }
// 	return string(bytes), err
// }

// func VerifyPassword(UserPassword string, ProvidedPassword string) bool {
// 	err := bcrypt.CompareHashAndPassword([]byte(ProvidedPassword), []byte(UserPassword))
// 	// check := true
// 	// msg := ""
// 	// if err != nil {
// 	// 	msg = fmt.Sprintf("Email of password is wrong")
// 	// 	check = false
// 	// }
// 	return err == nil

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
// 		// if existingUser.ID != "" {
// 		// 	c.JSON(http.StatusBadRequest, gin.H{"error": "user does not exist"})
// 		// 	return
// 		// }
// 		// user.Password, _ = HashPassword(user.Password)/
// 		// hashedPassword, err := utils.HashPassword(user.Password)

// 		hashedPassword, err := HashPassword(user.Password)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
// 			return
// 		}
// 		user.Password = hashedPassword

// 		if err := models.DB.Create(&user).Error; err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
// 			return
// 		}

// 		c.JSON(http.StatusOK, gin.H{"message": "User Registration Successful"})
// 	}
// }

// func LoginUser() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// var loginData models.User
// 		// var foundUser models.User

// 		var req struct {
// 			Email    string `json:"Email"`
// 			Password string `json:"Password"`
// 		}

// 		if err := c.ShouldBindJSON(&req); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request"})
// 			return
// 		}

// 		var user models.User
// 		if err := models.DB.Where("Email=?", req.Email).First(&user).Error; err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "The email or password is incorrect"})
// 			return
// 		}

// 		// 	BCRYPT AGAIN REPEAT
// 		// if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(loginData.Password)); err != nil {
// 		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Email or Password is wrong"})
// 		// }

// 		if !VerifyPassword(user.Password, req.Password) {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "The password or email invalid"})
// 		}

// 		token, err := utils.GenToken(user)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
// 			return
// 		}
// 		c.JSON(http.StatusOK, gin.H{"message": "User login successful", "token": token})
// 	}
// }

// func DashBoard() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.String(http.StatusOK, "Welcome to the dashboard")

//		}
//	}
//
//	func LogoutUser() gin.HandlerFunc {
//		return func(c *gin.Context) {
//			c.SetCookie("token", "", -1, "/", "", false, true)
//			c.JSON(http.StatusOK, gin.H{"message": "Logout Successful"})
//		}
//	}
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"jwt/database"
	"jwt/models"
	"jwt/utils"
)

// Secret Key for JWT

// Register Handler
func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Hash Password
	// hashedPassword, err := hashPassword(user.Password)
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not hash password"})
		return
	}
	user.Password = hashedPassword

	// Save User
	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not register user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

// Login Handler
func Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var user models.User
	if err := database.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Check Password
	if !utils.CheckPassword(user.Password, req.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Generate Token
	token, err := utils.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"messsage": "Login successfully done", "token": token})
}

func Dashboard(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the dashboard!",
	})
}

func Logout(c *gin.Context) {
	// Set the cookie with an expired time to remove it
	c.SetCookie("token", "", -1, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Logout successful",
	})
}
