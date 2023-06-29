package game

import (
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
		{"paper on scissors", Paper, Scissors, Lose},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			got := Game(test.humanMove, test.computerMove)
			if got != test.want {
				t.Errorf("%v on %v = %v, want %v", test.humanMove, test.computerMove, got, test.want)
			}
		})
	}
}
