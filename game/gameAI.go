// gameAI.go
// gameAI provides a way for the computer to choose a move according to a strategy
// You need to do an initialization of the random seed before using gameAI
package game

import (
	"math/rand"
)

// ComputerChooseMove returns a randomly generated move.
// Each move is equally likely to be chosen.
func ComputerChooseMove() Move {
	return Move(rand.Intn(int(MovesLength)))
}
