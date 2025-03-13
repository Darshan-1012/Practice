// package utils

// import (
// 	"jwt/models"
// 	"log"
// 	"os"
// 	"time"

// 	"github.com/golang-jwt/jwt/v4"
// )

// // Declared at Verify Token
// // type RegisterDetails struct {
// // 	Username    string `json:"username"`
// // 	Email string `json:"email"`

// // 	jwt.RegisteredClaims
// // }

// var SECRET_KEY string = os.Getenv("SECRET_KEY")

// func GenToken(User models.User) (signedToken string, err error) {

// 	// A SIMPLE WAY TO MAP THE MODELS DATA TO THE TOKEN(Change)
// 	// claims := &RegisterDetails{
// 	// 	Username:    Username,
// 	// 	Email: Email,
// 	// 	RegisteredClaims: jwt.RegisteredClaims{
// 	// 		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
// 	// 	},

// 	claims := jwt.MapClaims{
// 		"Username": User.Username,
// 		"Email":    User.Email,
// 		"exp":      time.Now().Add(time.Hour * 24).Unix(),
// 	}

// 	// EXTRA REFRESH TOKEN
// 	// refreshClaims := &RegisterDetails{
// 	// 	RegisteredClaims: jwt.RegisteredClaims{
// 	// 		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 168)),
// 	// 	},
// 	// }

// 	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))

// 	if err != nil {
// 		log.Panic(err)
// 		return
// 	}

// 	// EXTRA ERROR DETECTION FOR REFRESH TOKEN
// 	// refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))

// 	// if err != nil {
// 	// 	log.Panic(err)
// 	// 	return
// 	// }

// 	return token, err
// }

// // Different File for Verification of token
// // func AuthToken(signedToken string) (claims *RegisterDetails, msg string) {
// // 	token, err := jwt.ParseWithClaims(
// // 		signedToken,
// // 		&RegisterDetails{},
// // 		func(token *jwt.Token) (interface{}, error) {
// // 			return []byte(SECRET_KEY), nil
// // 		},
// // 	)

// // 	if err != nil {
// // 		msg = err.Error()
// // 		return
// // 	}

// // 	claims, ok := token.Claims.(*RegisterDetails)
// // 	if !ok {
// // 		msg = fmt.Sprintf("the token is invalid")
// // 		return
// // 	}

// // 	// ERROR TO CHECK IF TOKEN IS EXPIRED OR NOT
// // 	// if claims.ExpiresAt == nil || claims.ExpiresAt.Time.Unix() < time.Now().Unix() {
// // 	// 	msg = fmt.Sprint("token is expired")
// // 	// 	return
// // 	// }

// // 	return claims, msg
// // }

package utils

import (
	"jwt/models"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("your_secret_key")

// Hash Password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// Check Password
func CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

// Generate JWT Token
func GenerateToken(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"username": user.Username,
		"email":    user.Email,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
