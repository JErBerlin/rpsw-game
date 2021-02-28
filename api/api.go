// api package provides the rest api to play the game rock, paper, scissors.
// It provides a server start function and a handler for a play post request
package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"rpsw-game/game"
)

const port = ":8080"

// PlayRequest represents the move (or throw) request from the player
type PlayRequest struct {
	Throw string
}

// PlayOutcome has information about the outcome of the game and which was the computer move
type PlayOutcome struct {
	Outcome      string
	ComputerMove string
}

// StartGameServer calls and runs a default gin server and assigns the only post handler
// (we use the gin library, more practical and expressive than the standard go http libraries)
func StartGameServer() {
	// some initialization to use game library
	InitAPI()

	// default gin router (logger and recovery functions)
	r := gin.Default()

	// just need one post handler to submit movement and get outcome
	r.POST("/play", PlayGame)

	r.Run(port)
}

// InitAPI() makes initialization jobs
func InitAPI() {
	// we need to do that to use the library game later
	rand.Seed(time.Now().Unix())
}

// PlayGame gets the post request, decodes and validates, calls the game and send back the outcome
func PlayGame(c *gin.Context) {
	body := c.Request.Body
	reqJSON, err := ioutil.ReadAll(body)
	if err != nil {
		log.Println("body from post request could not be read:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "request body could not be processed"})
		return
	}
	playerMove, err := MakeMoveFromJSON(reqJSON)
	if err != nil {
		log.Println("a valid move could not be generated from json:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad formatted json, the request could not be completed"})
		return
	}

	computerMove := game.ComputerChooseMove()
	outcome := game.Game(playerMove, computerMove).String()

	playOutcome := PlayOutcome{
		outcome,
		computerMove.String(),
	}

	c.JSON(http.StatusOK, playOutcome)
}

// MakeMoveFromJSON decodes and validates the move request
// and returns a valid move or an error
func MakeMoveFromJSON(reqJSON []byte) (game.Move, error) {
	var throw PlayRequest
	err := json.Unmarshal(reqJSON, &throw)
	if err != nil {
		return 0, err
	}
	for i := 0; i < int(game.MovesLength); i++ {
		m := game.Move(i)
		if strings.ToLower(throw.Throw) == m.String() {
			return m, nil
		}
	}
	return 0, fmt.Errorf("throw requested is not a valid move")
}
