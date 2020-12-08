package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

type Directions struct {
	right int
	down  int
}

func FindNoOfTreesPt1(slope [][]string) int {
	return FindNoOfTrees(slope, Directions{3, 1})
}

func FindTotalNoOfTreesPt2(slope [][]string, directions []Directions) int {
	totalMultiple := 1
	for _, direction := range directions {
		totalMultiple = totalMultiple * FindNoOfTrees(slope, direction)
	}

	return totalMultiple
}

func FindNoOfTrees(slope [][]string, direction Directions) int {
	width := len(slope[0])
	noOfTrees := 0
	i := 0
	for y := range slope {
		if int(math.Mod(float64(y), float64(direction.down))) != 0 {
			continue
		}

		x := int(math.Mod(float64(i*direction.right), float64(width)))
		if slope[y][x] == "#" {
			noOfTrees++
		}
		i++
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
	fmt.Println("Pt2 Total No of Trees")
	fmt.Println(FindTotalNoOfTreesPt2(slope, []Directions{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}))
}
