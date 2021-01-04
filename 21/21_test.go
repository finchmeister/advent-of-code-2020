package main

import (
	"testing"
)

var testInput = `mxmxvkd kfcds sqjhc nhms (contains dairy, fish)
trh fvjkl sbzzf mxmxvkd (contains dairy)
sqjhc fvjkl (contains soy)
sqjhc mxmxvkd sbzzf (contains fish)`

func TestFindCountOfIngredientsWithNoAllergensPt1(t *testing.T) {
	got := FindCountOfIngredientsWithNoAllergensPt1(parse(testInput))
	expected := 5

	if got != expected {
		t.Errorf("Expected %d, got %d", expected, got)
	}
}
