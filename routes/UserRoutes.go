package routes

import (
	"jwt/controllers"
	"jwt/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {

	{
		r.POST("/register", controllers.Register)
		r.POST("/login", controllers.Login)

	}
	auth := r.Group("/auth")
	auth.Use(middleware.AuthMiddleware()) // Applying JWT authentication middleware
	{
		auth.GET("/dashboard", controllers.Dashboard)
	}
}

// package routes

// import (
// 	"jwt/controllers"

// 	"github.com/gin-gonic/gin"
// )

// // (/dashboard, /login, /register, /logout)

// func UserRoutes(incomingRoutes *gin.Engine) {
// 	// auth:= incomingRoutes.Use(middleware.JwtOauthUser()) // Not required
// 	incomingRoutes.POST("/register", controllers.RegisterUser())
// 	incomingRoutes.POST("/login", controllers.LoginUser())
// 	// incomingRoutes.POST("/login", controllers.LoginUser())
// 	// incomingRoutes.POST("/dashboard", middleware.JwtOauthUser(), controllers.DashBoard())
// 	incomingRoutes.POST("/logout", controllers.LogoutUser())
// }

// //
