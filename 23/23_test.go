package main

import (
	"testing"
)

var testInput = `389125467`

func TestFindLabelsOnCupsAfter1Pt1TenMoves(t *testing.T) {
	got := FindLabelsOnCupsAfter1Pt1(parse(testInput), 10)
	expected := "92658374"

	if got != expected {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}

func TestFindLabelsOnCupsAfter1Pt1OneHundredMoves(t *testing.T) {
	got := FindLabelsOnCupsAfter1Pt1(parse(testInput), 100)
	expected := "67384529"

	if got != expected {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}

func TestFindMultipleOfLabelsAfterCup1Pt2(t *testing.T) {
	got := FindMultipleOfLabelsAfterCup1Pt2(parse(testInput), TenMillion)
	expected := 149245887792

	if got != expected {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}
