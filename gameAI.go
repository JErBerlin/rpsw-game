package main

import "math/rand"

// ComputerChooseMove returns a randomly generated move according
// a given strategy
func ComputerChooseMove() Move {
	// mixed strategy: equiprobable
	return Move(rand.Intn(int(MovesLength)))
}
