// game.go
// game package provide the rules and the interface to play rock, paper, scissors.
// It provides the mean to compute the outcome, given the moves of the two players.
// You can use the functions in gameAI.go to choose a move for the computer.
package game

// Move codes the possible throws or moves of the game
type Move int

const (
	Rock Move = iota
	Paper
	Scissors
	MovesLength
)

// String returns the string labels of enum ints
func (m Move) String() string {
	switch m {
	case Rock:
		return "rock"
	case Paper:
		return "paper"
	case Scissors:
		return "scissors"
	}
	return ""
}

// Player codes the id of the players
type Player int

const (
	Human Player = iota
	Computer
)

// Outcome codes lose, draw and win
type Outcome int

const (
	Lose Outcome = iota - 1
	Draw
	Win
)

func (o Outcome) String() string {
	switch o {
	case Lose:
		return "lose"
	case Draw:
		return "draw"
	case Win:
		return "win"
	}
	return ""
}

// Outcome table tells you if player one will lose, draw or win against player two
// given moves of player one and two. We write the whole matrix although not needed.
// (by design of the game, it is an anti-symmetric, hollow matrix)
var OutcomeTable = [][]Outcome{
	{Draw, Lose, Win},
	{Win, Draw, Lose},
	{Lose, Win, Draw},
}

// Game returns if Player one will lose, draw or win against player two given
// the moves of player one and two by looking it up in the OutcomeTable
func Game(playerOneMove, playerTwoMove Move) Outcome {
	return OutcomeTable[playerOneMove][playerTwoMove]
}
