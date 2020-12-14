package main

import (
	"testing"
)

var testInput = `939
7,13,x,x,59,x,31,19`

func TestFindEarliestBusMultipliedPt1(t *testing.T) {
	got := FindEarliestBusMultipliedPt1(parse(testInput))
	expected := 295

	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}
