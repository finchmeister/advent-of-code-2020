package main

import (
	"testing"
)

var testInput = `0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: "a"
5: "b"

ababbb
bababa
abbbab
aaabbb
aaaabbb`

func TestFindMessageCountMatchRule0Pt1(t *testing.T) {
	got := FindMessageCountMatchRule0Pt1(parse(testInput))
	expected := 2

	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}
