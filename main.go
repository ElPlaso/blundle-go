package main

import (
	"github.com/gin-gonic/gin"
)

type puzzle struct {
	PuzzleId        string   `json:"puzzleid"`
	Fen             string   `json:"fen"`
	Rating          float64  `json:"rating"`
	RatingDeviation float64  `json:"ratingdeviation"`
	Moves           []string `json:"moves"`
	Themes          []string `json:"themes"`
}

type puzzleData struct {
	Key    string `json:"key"`
	Puzzle puzzle `json:"puzzle"`
}

func main() {
	router := gin.Default()
	router.GET("/daily-puzzle", getDailyPuzzle)

	router.Run("localhost:8080")
}

func getDailyPuzzle(c *gin.Context) {
	// TODO
}
