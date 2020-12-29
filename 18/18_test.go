package main

import (
	"testing"
)

type Fixtures struct {
	input    string
	expected int
}

type TransformFixtures struct {
	input    string
	expected string
}

var testInputPt1 = `1 + 2 * 3 + 4 * 5 + 6
1 + (2 * 3) + (4 * (5 + 6))`

func TestEvaluateExpressionPt1(t *testing.T) {
	fixtures := []Fixtures{
		{"1 + 2 * 3 + 4 * 5 + 6", 71},
		{"1 + (2 * 3) + (4 * (5 + 6))", 51},
		{"2 * 3 + (4 * 5)", 26},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 437},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 12240},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 13632},
	}

	for _, fixture := range fixtures {
		got := evaluateExpressionPt1(fixture.input)
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

func TestFindSumOfExpressionsPt1(t *testing.T) {
	got := FindSumOfExpressionsPt1(parse(testInputPt1))
	expected := 122

	if got != expected {
		t.Errorf(
			"Expected %d; got %d",
			expected,
			got,
		)
	}
}

func TestTransformExpression(t *testing.T) {
	fixtures := []TransformFixtures{
		{"(1 + 2) + 4", "(((1 + 2)) + 4)"},
		{"1 + (2 + 4)", "(1 + ((2 + 4)))"},
		{"1 + 2 * 3 + 4 * 5 + 6", "(1 + 2) * (3 + 4) * (5 + 6)"},
		{"1 + (2 * 3) + (4 * (5 + 6))", "((1 + (2 * 3)) + (4 * ((5 + 6))))"},
		{"2 * 3 + (4 * 5)", "2 * (3 + (4 * 5))"},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", "(5 + (8 * ((3 + 9) + 3) * 4 * 3))"},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))",
			"5 * 9 * (7 * 3 * (3 + 9) * (3 + ((8 + 6) * 4)))"},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2",
			"(((((2 + 4) * 9) * (((6 + 9) * (8 + 6)) + 6)) + 2) + 4) * 2"},
	}

	for _, fixture := range fixtures {
		got := transformExpressionPt2(fixture.input)
		expected := fixture.expected

		if got != expected {
			t.Errorf(
				"Expected %v; got %v",
				fixture.expected,
				got,
			)
		}
	}
}

func TestEvaluateExpressionPt2(t *testing.T) {
	fixtures := []Fixtures{
		{"1 + 2 * 3 + 4 * 5 + 6", 231},
		{"1 + (2 * 3) + (4 * (5 + 6))", 51},
		{"2 * 3 + (4 * 5)", 46},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 1445},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 669060},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 23340},
	}

	for _, fixture := range fixtures {
		got := evaluateExpressionPt2(fixture.input)
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

func TestTransformedExpressionPt2(t *testing.T) {
	fixtures := []Fixtures{
		{"(1 + 2) * (3 + 4) * (5 + 6)", 231},
		{"((1 + (2 * 3)) + (4 * ((5 + 6))))", 51},
		{"2 * (3 + (4 * 5))", 46},
		{"5 + (8 * ((3 + 9) + 3) * 4 * 3)", 1445},
		{"5 * 9 * (7 * 3 * (3 + 9) * (3 + (8 + 6 * 4)))", 669060},
		{"(((((2 + 4) * 9) * (((6 + 9) * (8 + 6)) + 6)) + 2) + 4) * 2", 23340},
	}

	for _, fixture := range fixtures {
		got := evaluateExpressionPt1(fixture.input)
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
