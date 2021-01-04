package main

import (
	"testing"
)

var testInput = `Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10`

func TestFindWinningPlayersScorePt1(t *testing.T) {
	got := FindWinningPlayersScorePt1(parse(testInput))
	expected := 306

	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}
