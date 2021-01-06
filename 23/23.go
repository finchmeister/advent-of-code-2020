package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const TenMillion = 10000000
const OneMillion = 1000000

func FindLabelsOnCupsAfter1Pt1(cupLabelling []int, noOfMoves int) string {
	return getLabelsOrderedAfterOne(playGame(cupLabelling, noOfMoves))
}

func FindMultipleOfLabelsAfterCup1Pt2(cupLabelling []int, noOfMoves int) int {
	cupLabelling = updateCupLabellingPt2(cupLabelling, OneMillion)

	cupLabellingLinkedList := playGame(cupLabelling, noOfMoves)

	return cupLabellingLinkedList[1] * cupLabellingLinkedList[cupLabellingLinkedList[1]]
}

func updateCupLabellingPt2(cupLabelling []int, size int) []int {
	newCupLabelling := make([]int, len(cupLabelling))
	copy(newCupLabelling, cupLabelling)

	for i := len(cupLabelling) + 1; i <= size; i++ {
		newCupLabelling = append(newCupLabelling, i)
	}

	return newCupLabelling
}

func playGame(cupLabelling []int, noOfMoves int) map[int]int {
	cupLabellingLinkedList := makeCupLabellingLinkedList(cupLabelling)

	currentCup := cupLabelling[0]
	for i := 0; i < noOfMoves; i++ {
		var pickedUpCups []int
		nextCup := currentCup
		for j := 0; j < 3; j++ {
			nextCup = cupLabellingLinkedList[nextCup]
			pickedUpCups = append(pickedUpCups, nextCup)
		}
		nextCup = cupLabellingLinkedList[nextCup]

		destinationCup := getDestinationCup(currentCup, cupLabellingLinkedList, pickedUpCups)
		cupLabellingLinkedList[currentCup] = nextCup

		pickedUpCupsNextCup := cupLabellingLinkedList[destinationCup]
		cupLabellingLinkedList[destinationCup] = pickedUpCups[0]
		cupLabellingLinkedList[pickedUpCups[2]] = pickedUpCupsNextCup

		currentCup = nextCup
	}

	return cupLabellingLinkedList
}

func getDestinationCup(currentCup int, cupLabellingLinkedList map[int]int, pickedUpCups []int) int {
	modulo := len(cupLabellingLinkedList)
	destinationCup := subtractWithWrap(currentCup, 1, modulo)

	for contains(destinationCup, pickedUpCups) {
		destinationCup = subtractWithWrap(destinationCup, 1, modulo)
	}

	return destinationCup
}

func makeCupLabellingLinkedList(cupLabelling []int) map[int]int {
	cupLabellingLinkedList := make(map[int]int, len(cupLabelling))

	for i := range cupLabelling {
		nextI := subtractWithNoWrap(i, -1, len(cupLabelling))
		cupLabellingLinkedList[cupLabelling[i]] = cupLabelling[nextI]
	}

	return cupLabellingLinkedList
}

func getLabelsOrderedAfterOne(cupLabellingLinkedList map[int]int) string {
	var orderedAfterOne []string
	cup := 1
	for cupLabellingLinkedList[cup] != 1 {
		cup = cupLabellingLinkedList[cup]
		orderedAfterOne = append(orderedAfterOne, strconv.Itoa(cup))
	}

	return strings.Join(orderedAfterOne, "")
}

func subtractWithWrap(a, b int, modulo int) int {
	x := mod(a-b, modulo)
	if x == 0 {
		return modulo
	}

	return x
}

func subtractWithNoWrap(a, b int, modulo int) int {
	x := mod(a-b, modulo)

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
	fmt.Println("Pt2")
	fmt.Println(FindMultipleOfLabelsAfterCup1Pt2(parse(loadFile()), TenMillion))
}
