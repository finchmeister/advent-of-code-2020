package main

import (
	"testing"
)

var testInput = `.#.
..#
###`

func TestFindCubesInActiveStatePt1(t *testing.T) {
	got := FindCubesInActiveStatePt1(parse(testInput), 6)
	expected := 112

	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}
