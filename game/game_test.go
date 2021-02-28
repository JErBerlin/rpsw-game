package game

import (
	"fmt"
	"testing"
)

func TestGame(t *testing.T) {
	var tests = []struct {
		desc         string
		humanMove    Move
		computerMove Move
		want         Outcome
	}{
		{"rock on rock", Rock, Rock, Draw},
		{"paper on paper", Paper, Paper, Draw},
		{"scissors on scissors", Scissors, Scissors, Draw},
		{"rock on paper", Rock, Paper, Lose},
		{"rock on scissors", Rock, Scissors, Win},
		{"paper on rock", Paper, Rock, Win},
		{"paper on scissors", Paper, Scissors, Lose},
		{"scissors on rock", Scissors, Rock, Lose},
		{"scissors on paper", Scissors, Paper, Win},
	}
	for _, test := range tests {
		fmt.Printf("Test run: %v\n", test)
		got := Game(test.humanMove, test.computerMove)
		if got != test.want {
			t.Errorf("%s: %v on %v = %v, want %v", test.desc, test.humanMove, test.computerMove, got, test.want)
		}
	}
}
