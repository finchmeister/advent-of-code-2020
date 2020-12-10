package main

import (
	"testing"
)

var testInput = `35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576`

func TestFindFirstNoToNotFollowRulePt1(t *testing.T) {
	got := FindFirstNoToNotFollowRulePt1(parse(testInput), 5)
	expected := 127

	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}

func TestCanSumBeMade(t *testing.T) {
	got := canSumBeMade([]int{35, 20, 15, 25, 47}, 40)
	expected := true

	if got != expected {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}
