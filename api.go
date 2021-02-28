package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
)

const port = ":8080"

// PlayRequest represents the move (or throw) request from the player
type PlayRequest struct {
	Throw string
}

// PlayRequest has information about the outcome of the game and which was the computer move
type PlayOutcome struct {
	Outcome      string
	ComputerMove string
}

// StartGameServer calls and runs a default gin server and assigns the only post handler
// (we use the gin library, more practical and expressive than the standard go http libraries)
func StartGameServer() {
	// default gin router (logger and recovery functions)
	r := gin.Default()
	// just need one post handler to submit movement and get outcome
	r.POST("/play", PlayGame)
	r.Run(port)
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

	computerMove := ComputerChooseMove()
	outcome := Game(playerMove, computerMove).String()

	playOutcome := PlayOutcome{
		outcome,
		computerMove.String(),
	}

	c.JSON(http.StatusOK, playOutcome)
}

// MakeMoveFromJSON decodes and validates the move request
// and returns a valid move or an error
func MakeMoveFromJSON(reqJSON []byte) (Move, error) {
	var throw PlayRequest
	err := json.Unmarshal(reqJSON, &throw)
	if err != nil {
		return 0, err
	}
	switch strings.ToLower(throw.Throw) {
	case "rock":
		return Rock, nil
	case "paper":
		return Paper, nil
	case "scissors":
		return Scissors, nil
	}
	return 0, fmt.Errorf("throw requested is not a valid move")
}

// ComputerChooseMove returns a randomly generated move according
// a given strategy
func ComputerChooseMove() Move {
	// mixed strategy: equiprobable
	return Move(rand.Intn(int(MovesLength)))
}
