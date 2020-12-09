package main

import (
	"testing"
)

var testInput string = `abc

a
b
c

ab
ac

a
a
a
a

b`

func TestFindSumOfGroupCountsPt1(t *testing.T) {
	got := FindSumOfGroupCountsPt1(parse(testInput))
	expected := 11

	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}
