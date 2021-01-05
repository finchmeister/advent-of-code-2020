package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func FindLabelsOnCupsAfter1Pt1(cupLabelling []int, noOfMoves int) string {
	currentCup := cupLabelling[0]
	for i := 0; i < noOfMoves; i++ {
		var pickedUpCups []int
		nextCup := currentCup
		for j := 0; j < 3; j++ {
			nextCup = getNextCup(nextCup, cupLabelling)
			pickedUpCups = append(pickedUpCups, nextCup)
		}

		destinationCup := getDestinationCup(currentCup, cupLabelling, pickedUpCups)

		cupLabelling = removeCups(pickedUpCups, cupLabelling)
		cupLabelling = addCupsAfterDestination(pickedUpCups, cupLabelling, destinationCup)

		currentCup = getNextCup(currentCup, cupLabelling)
	}

	return getLabelsOrderedAfterOne(cupLabelling)
}

func getLabelsOrderedAfterOne(cupLabelling []int) string {
	var orderedAfterOne []string
	cup := 1
	for getNextCup(cup, cupLabelling) != 1 {
		cup = getNextCup(cup, cupLabelling)
		orderedAfterOne = append(orderedAfterOne, strconv.Itoa(cup))
	}

	return strings.Join(orderedAfterOne, "")
}

func removeCups(cupsToRemove, cupLabelling []int) []int {
	var newCupLabelling []int
	for _, cup := range cupLabelling {
		if contains(cup, cupsToRemove) == false {
			newCupLabelling = append(newCupLabelling, cup)
		}
	}

	return newCupLabelling
}

func addCupsAfterDestination(cupsToAdd, cupLabelling []int, destination int) []int {
	var newCupLabelling []int
	indexToInsert := getCupIndex(destination, cupLabelling)

	newCupLabelling = append(cupLabelling[:indexToInsert+1], append(cupsToAdd, cupLabelling[indexToInsert+1:]...)...)

	return newCupLabelling
}

func getNextCup(currentCup int, cupLabelling []int) int {
	for i := range cupLabelling {
		if currentCup == cupLabelling[i] {
			key := mod(i+1, len(cupLabelling))
			return cupLabelling[key]
		}
	}
	panic("Cup not found")
}

func getCupIndex(cup int, cupLabelling []int) int {
	for i := range cupLabelling {
		if cupLabelling[i] == cup {
			return i
		}
	}

	panic("Cup not found")
}

func getDestinationCup(currentCup int, cupLabelling []int, pickedUpCups []int) int {
	modulo := getMax(cupLabelling)
	destinationCup := subtractWithWrap(currentCup, 1, modulo)

	for contains(destinationCup, pickedUpCups) {
		destinationCup = subtractWithWrap(destinationCup, 1, modulo)
	}

	return destinationCup
}

func subtractWithWrap(a, b int, modulo int) int {
	x := mod(a-b, modulo)
	if x == 0 {
		return modulo
	}

	return x
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func contains(value int, slice []int) bool {
	for _, v := range slice {
		if value == v {
			return true
		}
	}

	return false
}

func getMax(values []int) int {
	max := values[0]
	for _, v := range values {
		if v > max {
			max = v
		}
	}

	return max
}

func parse(input string) []int {
	var cupLabelling []int
	for _, cup := range strings.Split(input, "") {
		value, _ := strconv.Atoi(cup)
		cupLabelling = append(cupLabelling, value)
	}

	return cupLabelling
}

func loadFile() string {
	data, err := ioutil.ReadFile("23_input.txt")
	if err != nil {
		panic(err)
	}

	return string(data)
}

func main() {
	fmt.Println("Pt1")
	fmt.Println(FindLabelsOnCupsAfter1Pt1(parse(loadFile()), 100))
}
