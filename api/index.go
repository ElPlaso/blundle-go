package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Handler(c *gin.Context) {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Fatalf("Error loading .env file: %s", envErr)
	}

	router := gin.Default()
	router.GET("/daily-puzzle", getDailyPuzzle)
	router.GET("/add-puzzle", addPuzzle)

	router.Run("localhost:8080")
}
