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
