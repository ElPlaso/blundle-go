package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

type puzzleType struct {
	PuzzleId        string   `json:"puzzleid"`
	Fen             string   `json:"fen"`
	Rating          float64  `json:"rating"`
	RatingDeviation float64  `json:"ratingdeviation"`
	Moves           []string `json:"moves"`
	Themes          []string `json:"themes"`
}

type puzzleData struct {
	Key    string     `json:"key"`
	Puzzle puzzleType `json:"puzzle"`
}

func main() {
	router := gin.Default()
	router.GET("/daily-puzzle", getDailyPuzzle)

	router.Run("localhost:8080")
}

func getDailyPuzzle(c *gin.Context) {
	db, err := connect()

	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, "Failed to connect to database")
		return
	}

	// Get latest puzzle
	rows, err := db.Query(context.Background(), "SELECT * FROM puzzle ORDER BY id DESC LIMIT 1")
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError, "Error querying database")
		return
	}
	defer rows.Close()

	count := 0

	for rows.Next() {
		count++

		var id int
		var puzzle string

		rowErr := rows.Scan(&id, &puzzle)

		if rowErr != nil {
			log.Print(rowErr)
			c.JSON(http.StatusInternalServerError, "Error reading puzzle data")
			return
		}

		parsedPuzzle := puzzleType{}

		puzzleErr := json.Unmarshal([]byte(puzzle), &parsedPuzzle)

		if puzzleErr != nil {
			log.Println(puzzleErr)
			c.JSON(http.StatusInternalServerError, "Error parsing puzzle")
			return
		}

		puzzleData := puzzleData{
			Key:    strconv.Itoa(id),
			Puzzle: parsedPuzzle,
		}

		c.JSON(http.StatusOK, puzzleData)

		if count == 1 {
			return
		}
	}

	c.JSON(http.StatusNotFound, "Puzzle not found")
}

func connect() (*pgx.Conn, error) {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Fatalf("Error loading .env file: %s", envErr)
	}

	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	return conn, nil
}
