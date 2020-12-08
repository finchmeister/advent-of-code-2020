package main

import "testing"

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

func TestFindNumberOfTreesA(t *testing.T) {
	testSlope := parseSlope(testInput)
	got := FindNoOfTreesPt1(testSlope)
	expected := 7

	if got != expected {
		t.Error("Fail")
	}
}
