package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type Food struct {
	ingredients []string
	allergens   []string
}

func FindCountOfIngredientsWithNoAllergensPt1(foods []Food) int {

	allergenIngredientsMap := make(map[string][][]string)

	for _, food := range foods {
		for _, allergen := range food.allergens {
			if allergenIngredientsMap[allergen] == nil {
				allergenIngredientsMap[allergen] = [][]string{food.ingredients}
			} else {
				allergenIngredientsMap[allergen] = append(allergenIngredientsMap[allergen], food.ingredients)
			}
		}
	}

	var ingredientsMustContainAllergen []string
	for allergen := range allergenIngredientsMap {
		for _, testIngredient := range allergenIngredientsMap[allergen][0] {
			if doAllFoodsIngredientsContainIngredient(allergenIngredientsMap[allergen], testIngredient) {
				if contains(testIngredient, ingredientsMustContainAllergen) == false {
					ingredientsMustContainAllergen = append(ingredientsMustContainAllergen, testIngredient)
				}
			}
		}
	}

	allIngredients := getAllIngredients(foods)
	ingredientsNotContainingAllergens := findIngredientsNotContainingAllergens(allIngredients, ingredientsMustContainAllergen)

	totalIngredientsNotContainingAllergensCount := 0
	for _, ingredientNotContainingAllergens := range ingredientsNotContainingAllergens {
		totalIngredientsNotContainingAllergensCount = totalIngredientsNotContainingAllergensCount + getIngredientOccurrenceCount(foods, ingredientNotContainingAllergens)
	}

	return totalIngredientsNotContainingAllergensCount
}

func FindCanonicalDangerousIngredientsListPt2(foods []Food) string {

	allergenIngredientsMap := make(map[string][][]string)

	for _, food := range foods {
		for _, allergen := range food.allergens {
			if allergenIngredientsMap[allergen] == nil {
				allergenIngredientsMap[allergen] = [][]string{food.ingredients}
			} else {
				allergenIngredientsMap[allergen] = append(allergenIngredientsMap[allergen], food.ingredients)
			}
		}
	}

	allergenMustBeIngredientMap := make(map[string][]string)
	for allergen := range allergenIngredientsMap {
		for _, testIngredient := range allergenIngredientsMap[allergen][0] {
			if doAllFoodsIngredientsContainIngredient(allergenIngredientsMap[allergen], testIngredient) {
				if contains(testIngredient, allergenMustBeIngredientMap[allergen]) == false {
					allergenMustBeIngredientMap[allergen] = append(allergenMustBeIngredientMap[allergen], testIngredient)
				}
			}
		}
	}

	allergenIngredientMap := make(map[string]string)

	for len(allergenMustBeIngredientMap) > 0 {
		for allergen := range allergenMustBeIngredientMap {
			if len(allergenMustBeIngredientMap[allergen]) == 1 {
				ingredient := allergenMustBeIngredientMap[allergen][0]
				allergenIngredientMap[allergen] = ingredient

				allergenMustBeIngredientMap = removeIngredientFromAllergenMustBeIngredientMap(ingredient, allergenMustBeIngredientMap)
			}
		}
	}

	var allergens []string
	for allergen := range allergenIngredientMap {
		allergens = append(allergens, allergen)
	}

	sort.Strings(allergens)

	var dangerousIngredients []string
	for _, allergen := range allergens {
		dangerousIngredients = append(dangerousIngredients, allergenIngredientMap[allergen])
	}

	return strings.Join(dangerousIngredients, ",")
}

func removeIngredientFromAllergenMustBeIngredientMap(ingredient string, allergenMustBeIngredientMap map[string][]string) map[string][]string {
	for allergen := range allergenMustBeIngredientMap {
		allergenMustBeIngredientMap[allergen] = removeFromSlice(allergenMustBeIngredientMap[allergen], ingredient)
		if len(allergenMustBeIngredientMap[allergen]) == 0 {
			delete(allergenMustBeIngredientMap, allergen)
		}
	}

	return allergenMustBeIngredientMap
}

func removeFromSlice(s []string, remove string) []string {
	for i := range s {
		if s[i] == remove {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func getIngredientOccurrenceCount(foods []Food, ingredient string) int {
	ingredientOccurrenceCount := 0
	for _, food := range foods {
		if contains(ingredient, food.ingredients) {
			ingredientOccurrenceCount++
		}
	}

	return ingredientOccurrenceCount
}

func getAllIngredients(foods []Food) []string {
	var allIngredients []string
	for _, food := range foods {
		for _, ingredient := range food.ingredients {
			if contains(ingredient, allIngredients) == false {
				allIngredients = append(allIngredients, ingredient)
			}
		}
	}

	return allIngredients
}

func findIngredientsNotContainingAllergens(allIngredients []string, containAllergens []string) []string {
	var ingredientsNotContainingAllergens []string
	for _, ingredient := range allIngredients {
		if contains(ingredient, containAllergens) == false {
			ingredientsNotContainingAllergens = append(ingredientsNotContainingAllergens, ingredient)
		}
	}

	return ingredientsNotContainingAllergens
}

func doAllFoodsIngredientsContainIngredient(foodsIngredients [][]string, ingredient string) bool {
	for _, foodIngredients := range foodsIngredients {
		if contains(ingredient, foodIngredients) == false {
			return false
		}
	}

	return true
}

func contains(value string, slice []string) bool {
	for _, v := range slice {
		if value == v {
			return true
		}
	}

	return false
}

func parse(input string) []Food {
	rawRows := strings.Split(input, "\n")

	var foods []Food

	for _, rawRow := range rawRows {
		foods = append(foods, parseFood(rawRow))
	}

	return foods
}

func parseFood(rawRow string) Food {
	splitRawRow := strings.Split(rawRow, " (contains ")

	return Food{
		strings.Split(splitRawRow[0], " "),
		strings.Split(splitRawRow[1][:len(splitRawRow[1])-1], ", "),
	}
}

func loadFile() string {
	data, err := ioutil.ReadFile("21_input.txt")
	if err != nil {
		panic(err)
	}

	return string(data)
}

func main() {
	fmt.Println("Pt1")
	fmt.Println(FindCountOfIngredientsWithNoAllergensPt1(parse(loadFile())))
	fmt.Println("Pt2")
	fmt.Println(FindCanonicalDangerousIngredientsListPt2(parse(loadFile())))
}
