package main

import (
	"testing"
)

var testInput string = `..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`

type Fixtures struct {
	directions    Directions
	expectedTrees int
}

func TestFindNumberOfTreesPt1(t *testing.T) {
	testSlope := parseSlope(testInput)
	got := FindNoOfTreesPt1(testSlope)
	expected := 7

	if got != expected {
		t.Error("Fail")
	}
}

func TestFindNumberOfTrees(t *testing.T) {
	testSlope := parseSlope(testInput)
	fixtures := []Fixtures{
		{Directions{1, 1}, 2},
		{Directions{3, 1}, 7},
		{Directions{5, 1}, 3},
		{Directions{7, 1}, 4},
		{Directions{1, 2}, 2},
	}

	for _, fixture := range fixtures {
		got := FindNoOfTrees(testSlope, fixture.directions)
		expected := fixture.expectedTrees

		if got != expected {
			t.Errorf(
				"Expected FindNoOfTreesPt2(right: %d, down: %d) = %d; want %d",
				fixture.directions.right,
				fixture.directions.down,
				got,
				fixture.expectedTrees)
		}
	}
}

func TestFindTotalNoOfTreesPt2(t *testing.T) {
	testSlope := parseSlope(testInput)
	got := FindTotalNoOfTreesPt2(testSlope, []Directions{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	})
	expected := 336

	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}
