package main

import (
	"testing"
)

type Fixtures struct {
	input    string
	at       int
	expected int
}

func TestFindNumberSpokenExamplePt1(t *testing.T) {
	fixtures := []Fixtures{
		{"0,3,6", 3, 6},
		{"0,3,6", 4, 0},
		{"0,3,6", 5, 3},
		{"0,3,6", 6, 3},
		{"0,3,6", 7, 1},
		{"0,3,6", 8, 0},
		{"0,3,6", 9, 4},
		{"0,3,6", 10, 0},
	}

	for _, fixture := range fixtures {
		got := FindNumberSpoken(parse(fixture.input), fixture.at)
		expected := fixture.expected

		if got != expected {
			t.Errorf(
				"Expected %d; got %d",
				fixture.expected,
				got,
			)
		}
	}
}

func TestFindNumberSpokenPt1(t *testing.T) {
	fixtures := []Fixtures{
		{"0,3,6", 2020, 436},
		{"1,3,2", 2020, 1},
		{"2,1,3", 2020, 10},
		{"1,2,3", 2020, 27},
		{"2,3,1", 2020, 78},
		{"3,2,1", 2020, 438},
		{"3,1,2", 2020, 1836},
	}

	for _, fixture := range fixtures {
		got := FindNumberSpoken(parse(fixture.input), fixture.at)
		expected := fixture.expected

		if got != expected {
			t.Errorf(
				"Expected %d; got %d",
				fixture.expected,
				got,
			)
		}
	}
}

func TestFindNumberSpokenPt2(t *testing.T) {
	fixtures := []Fixtures{
		{"0,3,6", 30000000, 175594},
		{"1,3,2", 30000000, 2578},
		{"2,1,3", 30000000, 3544142},
		{"1,2,3", 30000000, 261214},
		{"2,3,1", 30000000, 6895259},
		{"3,2,1", 30000000, 18},
		{"3,1,2", 30000000, 362},
	}

	for _, fixture := range fixtures {
		got := FindNumberSpoken(parse(fixture.input), fixture.at)
		expected := fixture.expected

		if got != expected {
			t.Errorf(
				"Expected %d; got %d",
				fixture.expected,
				got,
			)
		}
	}
}
