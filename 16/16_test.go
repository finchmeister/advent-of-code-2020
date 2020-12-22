package main

import (
	"reflect"
	"testing"
)

var testInput = `class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12`

var testInputPt2 = `class: 0-1 or 4-19
row: 0-5 or 8-19
seat: 0-13 or 16-19

your ticket:
11,12,13

nearby tickets:
3,9,18
15,1,5
5,14,9`

func TestFindInvalidTicketScanningErrorRatePt1(t *testing.T) {
	got := FindInvalidTicketScanningErrorRatePt1(parse(testInput))
	expected := 71

	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}

func TestGetRuleColumnMapping(t *testing.T) {
	got := getRuleColumnMapping(parse(testInputPt2))
	expected := map[int]int{
		0: 1,
		1: 0,
		2: 2,
	}

	if reflect.DeepEqual(got, expected) == false {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}

func TestFindMultipleOfDepartedPt2(t *testing.T) {
	got := FindMultipleOfDepartedPt2(parse(loadFile()))
	expected := 1439429522627

	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}
