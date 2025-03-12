package main

import (
	"fmt"
	"jwt/database"
	"jwt/models"
	"jwt/routes"

	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Database Connection
	dsn := "root:@tcp(127.0.0.1:3306)/userdata?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	database.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto Migrate Models
	database.DB.AutoMigrate(&models.User{})

	// Initialize Gin Router
	r := gin.Default()

	// Authentication Routes
	routes.AuthRoutes(r)

	// Run Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "7000"
	}
	fmt.Println("Server running on port:", port)
	log.Fatal(r.Run(":" + port))
}

// package main

// import (
// 	"fmt"
// 	"jwt/database"
// 	"jwt/routes"
// 	"log"

// 	"github.com/gin-gonic/gin"
// 	"github.com/joho/godotenv"
// )

// func main() {

// 	// r := gin.Default()
// 	database.Connect()

// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// 	r := gin.Default()

// 	routes.UserRoutes(r)
// 	fmt.Println(" Server running on port 7000")

// 	r.Run(":7000")
// }
