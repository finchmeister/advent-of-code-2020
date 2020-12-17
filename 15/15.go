package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func FindNumberSpoken(startingNumbers []int, at int) int {
	i := 1
	numbersSpoken := make(map[int][]int)
	var lastSpokenNumber int
	for _, startingNumber := range startingNumbers {
		numbersSpoken[startingNumber] = append(numbersSpoken[startingNumber], i)
		i++
		lastSpokenNumber = startingNumber
	}

	var spokenNumber int
	for i <= at {
		if len(numbersSpoken[lastSpokenNumber]) == 1 {
			spokenNumber = 0
		}
		if len(numbersSpoken[lastSpokenNumber]) >= 2 {
			spokenNumber = getDifferenceOfLastElements(numbersSpoken[lastSpokenNumber])
		}
		numbersSpoken[spokenNumber] = append(numbersSpoken[spokenNumber], i)
		i++
		lastSpokenNumber = spokenNumber
	}

	return lastSpokenNumber
}

func getDifferenceOfLastElements(s []int) int {
	return s[len(s)-1:][0] - s[len(s)-2:][0]
}

func parse(input string) []int {
	var startingNumbers []int
	for _, value := range splitByComma(input) {
		number, _ := strconv.Atoi(value)
		startingNumbers = append(startingNumbers, number)
	}

	return startingNumbers
}

func splitByComma(input string) []string {
	return strings.Split(input, ",")
}

func loadFile() string {
	data, err := ioutil.ReadFile("15_input.txt")
	if err != nil {
		panic(err)
	}

	return string(data)
}

func main() {
	fmt.Println("Pt1")
	fmt.Println(FindNumberSpoken(parse(loadFile()), 2020))
	fmt.Println("Pt2")
	fmt.Println(FindNumberSpoken(parse(loadFile()), 30000000))
}
