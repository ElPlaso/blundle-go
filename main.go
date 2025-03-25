package main

import (
	"log"

	. "github.com/ElPlaso/blundle-go/shared"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Printf("Error loading .env file: %s", envErr)
	}

	router := gin.Default()
	router.GET("/daily-puzzle", GetDailyPuzzle)
	router.GET("/add-puzzle", AddPuzzle)

	router.Run("localhost:8080")
}
