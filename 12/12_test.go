package main

import (
	"testing"
)

var testInput = `F10
N3
F7
R90
F11`

func TestFindManhattanDistancePt1(t *testing.T) {
	got := FindManhattanDistancePt1(parse(testInput))
	expected := 25

	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}
