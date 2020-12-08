package main

import (
	"testing"
)

type Fixture struct {
	s                    string
	expectedBoardingPass BoardingPass
}

func TestCalculateBoardingPass(t *testing.T) {
	fixtures := []Fixture{
		{"FBFBBFFRLR", BoardingPass{44, 5, 357}},
		{"BFFFBBFRRR", BoardingPass{70, 7, 567}},
		{"FFFBBBFRRR", BoardingPass{14, 7, 119}},
		{"BBFFBBFRLL", BoardingPass{102, 4, 820}},
	}

	for _, fixture := range fixtures {
		got := calculateBoardingPass(fixture.s)
		expected := fixture.expectedBoardingPass

		if got != expected {
			t.Errorf(
				"Expected BoardingPass(row: %d, column: %d, id: %d); got BoardingPass(row: %d, column: %d, id: %d);",
				fixture.expectedBoardingPass.row,
				fixture.expectedBoardingPass.column,
				fixture.expectedBoardingPass.seatId,
				got.row,
				got.column,
				got.seatId,
			)
		}
	}
}

func TestFindHighestSeatIdPt1(t *testing.T) {
	got := FindHighestSeatIdPt1([]string{
		"FBFBBFFRLR",
		"BFFFBBFRRR",
		"FFFBBBFRRR",
		"BBFFBBFRLL",
	})
	expected := 820

	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}
