package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func FindFirstNoToNotFollowRulePt1(data []int, preambleLen int) int {
	for i := 0; i < (len(data) - preambleLen); i++ {
		nextValue := data[i+preambleLen]
		if canSumBeMade(data[i:i+preambleLen], nextValue) == false {
			return nextValue
		}
	}
	panic("Solution not found")
}

func canSumBeMade(data []int, expectedSum int) bool {
	for i := range data {
		for j := range data {
			if data[i]+data[j] == expectedSum {
				return true
			}
		}
	}
	return false
}

func parse(input string) []int {
	var data []int
	for _, row := range splitByNewLine(input) {
		value, _ := strconv.Atoi(row)
		data = append(data, value)
	}

	return data
}

func splitByNewLine(input string) []string {
	return strings.Split(input, "\n")
}

func loadFile() string {
	data, err := ioutil.ReadFile("09_input.txt")
	if err != nil {
		panic(err)
	}

	return string(data)
}

func main() {
	fmt.Println("Pt1")
	fmt.Println(FindFirstNoToNotFollowRulePt1(parse(loadFile()), 25))
}
