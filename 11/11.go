package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"reflect"
	"strings"
)

func FindStabilisingValuePt1(seatLayout [][]string) int {
	for true {
		nextLayout := getNextLayout(seatLayout)
		if isSeatLayoutEqual(seatLayout, nextLayout) {
			break
		}
		seatLayout = nextLayout
	}

	return countNoOfOccupiedSeats(seatLayout)
}

func getNextLayout(seatLayout [][]string) [][]string {
	nextLayout := makeNewLayout(seatLayout)
	for y := range seatLayout {
		for x := range seatLayout[y] {
			nextLayout[y][x] = getNextSeat(y, x, seatLayout)
		}
	}

	return nextLayout
}

func countNoOfOccupiedSeats(seatLayout [][]string) int {
	count := 0
	for y := range seatLayout {
		for x := range seatLayout[y] {
			if seatLayout[y][x] == "#" {
				count++
			}
		}
	}
	return count
}

func makeNewLayout(seatLayout [][]string) [][]string {
	nextLayout := make([][]string, len(seatLayout))
	for i := range nextLayout {
		nextLayout[i] = make([]string, len(seatLayout[0]))
	}

	return nextLayout
}

func isSeatLayoutEqual(seatLayoutA [][]string, seatLayoutB [][]string) bool {
	return reflect.DeepEqual(seatLayoutA, seatLayoutB)
}

func occupiedAdjacentCount(y int, x int, seatLayout [][]string) int {
	y0 := int(math.Max(0, float64(y-1)))
	yN := int(math.Min(float64(len(seatLayout)-1), float64(y+1)))
	x0 := int(math.Max(0, float64(x-1)))
	xN := int(math.Min(float64(len(seatLayout[0])-1), float64(x+1)))
	occupiedCount := 0
	for yI := y0; yI <= yN; yI++ {
		for xI := x0; xI <= xN; xI++ {
			if yI == y && xI == x {
				continue
			}

			if seatLayout[yI][xI] == "#" {
				occupiedCount++
			}
		}
	}

	return occupiedCount
}

func getNextSeat(y int, x int, seatLayout [][]string) string {
	currentSeat := seatLayout[y][x]
	if currentSeat == "." {
		return "."
	}
	occupiedCount := occupiedAdjacentCount(y, x, seatLayout)
	if currentSeat == "L" && occupiedCount == 0 {
		return "#"
	}
	if currentSeat == "#" && occupiedCount >= 4 {
		return "L"
	}

	return currentSeat
}

func parse(input string) [][]string {
	var data [][]string
	for _, row := range splitByNewLine(input) {
		value := strings.Split(row, "")
		data = append(data, value)
	}

	return data
}

func splitByNewLine(input string) []string {
	return strings.Split(input, "\n")
}

func loadFile() string {
	data, err := ioutil.ReadFile("11_input.txt")
	if err != nil {
		panic(err)
	}

	return string(data)
}

func main() {
	fmt.Println("Pt1")
	fmt.Println(FindStabilisingValuePt1(parse(loadFile())))
}
