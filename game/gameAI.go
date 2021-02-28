// gameAI provides a way for the computer to choose a move according to a strategy
// You need to do an initialization of the random seed before using gameAI
package game

import (
	"math/rand"
)

// ComputerChooseMove returns a randomly generated move according
// a given strategy
func ComputerChooseMove() Move {
	// mixed strategy: equiprobable of strong elements, never choose rock
	// observe that we choose moves randomly from second element to the last
	return Move(1+rand.Intn(int(MovesLength)-1))
}
