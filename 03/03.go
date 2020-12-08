package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func FindNoOfTreesPt1(slope [][]string) int {
	width := len(slope[0])
	noOfTrees := 0
	for y := range slope {
		x := int(math.Mod(float64(y*3), float64(width)))
		if slope[y][x] == "#" {
			noOfTrees++
		}
	}

	return noOfTrees
}

func parseSlope(rawString string) [][]string {
	lines := strings.Split(rawString, "\n")

	var slope [][]string

	for _, line := range lines {
		slope = append(slope, strings.Split(line, ""))
	}

	return slope
}

func loadSlope() string {
	data, err := ioutil.ReadFile("03_input.txt")
	if err != nil {
		panic(err)
	}

	return string(data)
}

func main() {
	slope := parseSlope(loadSlope())
	fmt.Println("Pt1 No of Trees")
	fmt.Println(FindNoOfTreesPt1(slope))
}
