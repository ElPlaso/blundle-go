package handler

import (
	"log"
	"net/http"

	. "github.com/ElPlaso/blundle-go/shared"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	app *gin.Engine
)

func init() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Fatalf("Error loading .env file: %s", envErr)
	}

	app = gin.Default()
	app.GET("/daily-puzzle", GetDailyPuzzle)
	app.GET("/add-puzzle", AddPuzzle)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
