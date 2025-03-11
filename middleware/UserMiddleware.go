package middleware

import (
	"fmt"
	"jwt/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JwtOauthUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No authentication header provided")})
			c.Abort()
			return
		}

		claims, err := utils.AuthToken(clientToken)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error:": err})
			c.Abort()
			return
		}

		c.Set("Email", claims.Email)
		c.Set("ID", claims.ID)
		c.Next()
	}
}
