package main

import (
	"testing"
)

var testInput = `nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`

func TestFindAccValuePt1(t *testing.T) {
	got := FindAccValuePt1(parse(testInput))
	expected := 5

	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}
