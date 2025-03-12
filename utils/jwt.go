package utils

import (
	"errors"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

// JWTSecretKey retrieves the secret key from .env
func JWTSecretKey() string {
	return os.Getenv("JWT_SECRET")
}

// Claims struct
type Claims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

// GenerateToken generates a new JWT token
// func GenerateToken(userID uint, email string) (string, error) {
// 	// Define expiration time
// 	expirationTime := time.Now().Add(24 * time.Hour)

// 	// Create JWT claims
// 	claims := &Claims{
// 		UserID: userID,
// 		Email:  email,
// 		RegisteredClaims: jwt.RegisteredClaims{
// 			ExpiresAt: jwt.NewNumericDate(expirationTime),
// 		},
// 	}

// 	// Create JWT token
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	return token.SignedString([]byte(JWTSecretKey()))
// }

// ValidateToken validates the provided JWT token
func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTSecretKey()), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
