// package routes

// import (
// 	"jwt/controllers"
// 	"jwt/middleware"

// 	"github.com/gin-gonic/gin"
// )

// func UserRoutes(r *gin.Engine) {
// 	r.POST("/register", controllers.RegisterUser())
// 	r.POST("/login", controllers.LoginUser())

// 	auth := r.Group("/auth")
// 	auth.Use(middleware.JwtOauthUser())
// 	{
// 		r.GET("/dashboard", controllers.DashBoard())
// 	}

// 	r.POST("/logout", controllers.LogoutUser())
// }

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

	r.POST("/logout", controllers.Logout)
}
