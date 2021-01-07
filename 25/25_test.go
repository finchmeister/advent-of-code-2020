package main

import (
	"testing"
)

var testInput = `5764801
17807724`

func TestFindEncryptionKeyPt1(t *testing.T) {
	got := FindEncryptionKeyPt1(parse(testInput))
	expected := 14897079

	if got != expected {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}
