package main

import (
	"testing"
)

var testInput = `L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`

func TestFindStabilisingValuePt1(t *testing.T) {
	got := FindStabilisingValuePt1(parse(testInput))
	expected := 37

	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}

func TestFindStabilisingValuePt2(t *testing.T) {
	got := FindStabilisingValuePt2(parse(testInput))
	expected := 26

	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}

func TestOccupiedAdjacentCountPt2A(t *testing.T) {
	input := `.......#.
...#.....
.#.......
.........
..#L....#
....#....
.........
#........
...#.....`

	got := occupiedAdjacentCountPt2(4, 3, parse(input))
	expected := 8

	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}

func TestOccupiedAdjacentCountPt2B(t *testing.T) {
	input := `.............
.L.L.#.#.#.#.
.............`

	got := occupiedAdjacentCountPt2(1, 2, parse(input))
	expected := 0

	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}

func TestOccupiedAdjacentCountPt2C(t *testing.T) {
	input := `.##.##.
#.#.#.#
##...##
...L...
##...##
#.#.#.#
.##.##.`

	got := occupiedAdjacentCountPt2(3, 3, parse(input))
	expected := 0

	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}

func TestOccupiedAdjacentCountPt2D(t *testing.T) {
	input := `#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##`

	got := occupiedAdjacentCountPt2(0, 0, parse(input))
	expected := 3

	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}

func TestGetNextLayout2(t *testing.T) {
	got := getNextLayout(parse(testInput), getNextSeatPt2)
	expected := parse(`#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##`)

	if isSeatLayoutEqual(got, expected) == false {
		t.Errorf("Expected\n%v \ngot\n%v",
			seatLayoutToString(expected), seatLayoutToString(got),
		)
	}
}

func TestGetNextLayout2A(t *testing.T) {
	got := getNextLayout(parse(`#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##`), getNextSeatPt2)
	expected := parse(`#.LL.LL.L#
#LLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLL#
#.LLLLLL.L
#.LLLLL.L#`)

	if isSeatLayoutEqual(got, expected) == false {
		t.Errorf("Expected\n%v \ngot\n%v",
			seatLayoutToString(expected), seatLayoutToString(got),
		)
	}
}

func TestGetNextLayout2AA(t *testing.T) {
	got := getNextLayout(parse(`#.L#.L#.L#
#LLLLLL.LL
L.L.L..#..
##L#.#L.L#
L.L#.#L.L#
#.L####.LL
..#.#.....
LLL###LLL#
#.LLLLL#.L
#.L#LL#.L#`), getNextSeatPt2)
	expected := parse(`#.L#.L#.L#
#LLLLLL.LL
L.L.L..#..
##L#.#L.L#
L.L#.LL.L#
#.LLLL#.LL
..#.L.....
LLL###LLL#
#.LLLLL#.L
#.L#LL#.L#`)

	if isSeatLayoutEqual(got, expected) == false {
		t.Errorf("Expected\n%v \ngot\n%v",
			seatLayoutToString(expected), seatLayoutToString(got),
		)
	}
}
