package main

import (
	"math/rand"
	"time"
)

func main() {
	InitComputerPlayer()
	StartGameServer()
}

// InitComputerPlayer initialize the random seed to be used later
func InitComputerPlayer() {
	rand.Seed(time.Now().Unix())
}
