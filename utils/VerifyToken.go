// package utils

// import (
// 	"fmt"

// 	"github.com/golang-jwt/jwt/v4"
// )

// type RegisterDetails struct {
// 	UserId uint   `json:"userid"`
// 	Email  string `json:"email"`
// 	jwt.RegisteredClaims
// }

// func AuthToken(signedToken string) (claims *RegisterDetails, msg string) {
// 	token, err := jwt.ParseWithClaims(
// 		signedToken,
// 		&RegisterDetails{},
// 		func(token *jwt.Token) (interface{}, error) {
// 			return []byte(SECRET_KEY), nil
// 		},
// 	)

// 	if err != nil {
// 		msg = err.Error()
// 		return
// 	}

// 	claims, ok := token.Claims.(*RegisterDetails)
// 	if !ok {
// 		msg = fmt.Sprintf("the token is invalid")
// 		return
// 	}

// 	// ERROR TO CHECK IF TOKEN IS EXPIRED OR NOT
// 	// if claims.ExpiresAt == nil || claims.ExpiresAt.Time.Unix() < time.Now().Unix() {
// 	// 	msg = fmt.Sprint("token is expired")
// 	// 	return
// 	// }

// 	return claims, msg
// }

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
