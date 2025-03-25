package main

import (
	"log"
	"os"

	. "github.com/ElPlaso/blundle-go/shared"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Printf("Error loading .env file: %s", envErr)
	}

	router := gin.New()
	router.Use(cors.Default())
	router.GET("/daily-puzzle", GetDailyPuzzle)
	router.GET("/add-puzzle", AddPuzzle)

	envPort, envPortFound := os.LookupEnv("PORT")

	address := ""
	if envPortFound {
		address = "0.0.0.0:" + envPort // For render
	} else {
		address = "localhost:8080"
	}

	router.Run(address)
}
