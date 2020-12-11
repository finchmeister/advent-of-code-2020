package main

import (
	"testing"
)

var testInput1 = `16
10
15
5
1
11
7
19
6
12
4`

var testInput2 = `28
33
18
42
31
14
46
20
48
47
24
23
49
45
19
38
39
11
1
32
25
35
8
17
7
9
4
2
34
10
3`

type Fixtures struct {
	input    string
	expected int
}

func TestFindDifferencesMultipliedPt1(t *testing.T) {
	fixtures := []Fixtures{
		{testInput1, 35},
		{testInput2, 220},
	}

	for _, fixture := range fixtures {
		got := FindDifferencesMultipliedPt1(parse(fixture.input))
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

func TestFindCountValidPermutationsPt2(t *testing.T) {
	fixtures := []Fixtures{
		{testInput1, 8},
		{testInput2, 19208},
	}

	for _, fixture := range fixtures {
		got := FindCountValidPermutationsPt2(parse(fixture.input))
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
