package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand/v2"
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
	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Fatalf("Error loading .env file: %s", envErr)
	}

	router := gin.Default()
	router.GET("/daily-puzzle", getDailyPuzzle)
	// router.POST("/puzzles", addPuzzle)

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

		var parsedPuzzle puzzleType

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

func addPuzzle(c *gin.Context) {
	// TODO: Authenticate
	// TODO: Get random puzzle and add to database
}

func connect() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func getRandomChessPuzzle() {
	// get puzzle with 3 to 4 moves
	numMoves := rand.IntN(4-3) + 3
	url := fmt.Sprintf("https://chess-puzzles.p.rapidapi.com/?rating=2000&themesType=ALL&playerMoves=%d&count=1", numMoves)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-key", os.Getenv("CHESS_PUZZLE_API_KEY"))
	req.Header.Add("x-rapidapi-host", "chess-puzzles.p.rapidapi.com")

	res, resErr := http.DefaultClient.Do(req)

	if resErr != nil {
		log.Println(resErr)
		return
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	var parsedBody any

	bodyErr := json.Unmarshal(body, &parsedBody)

	if bodyErr != nil {
		log.Println(bodyErr)
		return
		// TODO: Handle return
	}

	// puzzles := parsedBody.puzzles // TODO: Handle

	// puzzle := puzzles[0]

	// return puzzle
}
