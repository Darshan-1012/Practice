package routes

import (
	"jwt/controllers"
	"jwt/middleware"

	"github.com/gin-gonic/gin"
)

// (/dashboard, /login, /register, /logout)

func UserRoutes(incomingRoutes *gin.Engine) {
	// auth:= incomingRoutes.Use(middleware.JwtOauthUser()) // Not required
	incomingRoutes.POST("/register", controllers.RegisterUser())
	// incomingRoutes.GET("/login", controllers.LoginUser())
	incomingRoutes.POST("/login", controllers.LoginUser())
	incomingRoutes.POST("/dashboard", middleware.JwtOauthUser(), controllers.DashBoard())
	incomingRoutes.POST("/logout", controllers.LogoutUser())
}
