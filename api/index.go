package handler

import (
	"net/http"

	. "github.com/ElPlaso/blundle-go/shared"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	app *gin.Engine
)

func init() {
	app = gin.Default()
	app.Use(cors.Default())
	app.GET("/api/daily-puzzle", GetDailyPuzzle)
	app.GET("/api/add-puzzle", AddPuzzle)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
