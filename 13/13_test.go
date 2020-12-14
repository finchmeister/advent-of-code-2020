package main

import (
	"testing"
)

var testInput = `939
7,13,x,x,59,x,31,19`

type Fixtures struct {
	input    Notes
	expected int
}

func TestFindEarliestBusMultipliedPt1(t *testing.T) {
	got := FindEarliestBusMultipliedPt1(parse(testInput))
	expected := 295

	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}

func TestFindEarliestBusTimestampPt2(t *testing.T) {
	fixtures := []Fixtures{
		{parse(testInput), 1068781},
		{Notes{0, []string{"17", "x", "13", "19"}}, 3417},
		{Notes{0, []string{"17", "x", "13"}}, 102},
		{Notes{0, []string{"67", "7", "59", "61"}}, 754018},
		{Notes{0, []string{"67", "x", "7", "59", "61"}}, 779210},
		{Notes{0, []string{"67", "7", "x", "59", "61"}}, 1261476},
		{Notes{0, []string{"1789", "37", "47", "1889"}}, 1202161486},
	}

	for _, fixture := range fixtures {
		got := FindEarliestBusTimestampPt2(fixture.input)
		expected := fixture.expected

		if got != expected {
			t.Errorf(
				"Expected %d; got %d",
				fixture.expected,
				got,
			)
		}
	}
}
