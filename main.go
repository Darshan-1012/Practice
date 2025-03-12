package main

import (
	"jwt/database"
	"jwt/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	r := gin.Default()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.Connect()

	routes.UserRoutes(r)
	r.Run(":8080")

}
