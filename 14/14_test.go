package main

import (
	"testing"
)

var testInput = `mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`

var testInputPt2 = `mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1`

func TestFindSumOfValuesLeftInMemoryPt1(t *testing.T) {
	got := FindSumOfValuesLeftInMemoryPt1(parse(testInput))
	expected := 165

	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}

func TestFindSumOfValuesLeftInMemoryPt2(t *testing.T) {
	got := FindSumOfValuesLeftInMemoryPt2(parse(testInputPt2))
	expected := 208

	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}
