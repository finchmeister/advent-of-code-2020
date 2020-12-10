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

func FindMinMaxSumOfContiguousSetPt2(data []int, targetSum int) int {
	for sliceLength := 2; sliceLength < len(data); sliceLength++ {
		for startIndex := 0; startIndex <= (len(data) - sliceLength); startIndex++ {
			targetSlice := data[startIndex : startIndex+sliceLength]
			if sumsToValue(targetSlice, targetSum) == true {
				return getMin(targetSlice) + getMax(targetSlice)
			}
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

func sumsToValue(data []int, value int) bool {
	sum := 0
	for _, v := range data {
		sum = sum + v
	}

	return sum == value
}

func getMin(values []int) int {
	min := values[0]
	for _, v := range values {
		if v < min {
			min = v
		}
	}

	return min
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
	pt1Solution := FindFirstNoToNotFollowRulePt1(parse(loadFile()), 25)
	fmt.Println(pt1Solution)
	fmt.Println("Pt2")
	fmt.Println(FindMinMaxSumOfContiguousSetPt2(parse(loadFile()), pt1Solution))
}
