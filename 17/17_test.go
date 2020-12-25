package main

import (
	"testing"
)

var testInput = `.#.
..#
###`

func TestFindCubesInActiveStatePt1(t *testing.T) {
	got := FindCubesInActiveStatePt1(parsePt1(testInput), 6)
	expected := 112

	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}

func TestFindCubesInActiveStatePt2(t *testing.T) {
	got := FindCubesInActiveStatePt2(parsePt2(testInput), 6)
	expected := 848

	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}
