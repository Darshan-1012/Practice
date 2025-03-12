package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Secret Key for JWT
var jwtSecret = []byte("your_secret_key")

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		// Parse Token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Set User in Context
		claims, _ := token.Claims.(jwt.MapClaims)
		c.Set("user", claims)
		c.Next()
	}
}

// import (
// 	"fmt"
// 	"jwt/utils"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func JwtOauthUser() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		clientToken := c.Request.Header.Get("token")
// 		if clientToken == "" {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No authentication header provided")})
// 			c.Abort()
// 			return
// 		}

// 		claims, err := utils.AuthToken(clientToken)
// 		if err != "" {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error:": err})
// 			c.Abort()
// 			return
// 		}

// 		c.Set("Email", claims.Email)
// 		c.Set("ID", claims.ID)
// 		c.Next()
// 	}
// }
