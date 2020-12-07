package main

import (
	"testing"
)

func TestFindTwoSumTo2020(t *testing.T) {
	got := FindTwoSumTo2020([]int{1721, 979, 366, 299, 675, 1456})
	expected := TwoSumTo2020{1721, 299}

	if got != expected {
		t.Error("Fail")
	}
}

func TestFindThreeSumTo2020(t *testing.T) {
	got := FindThreeSumTo2020([]int{1721, 979, 366, 299, 675, 1456})
	expected := ThreeSumTo2020{979, 366, 675}

	if got != expected {
		t.Error("Fail")
	}
}
