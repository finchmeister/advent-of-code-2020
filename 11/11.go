package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"reflect"
	"strings"
)

func FindStabilisingValuePt1(seatLayout [][]string) int {
	return findStabilisingValue(seatLayout, getNextSeatPt1)
}

func FindStabilisingValuePt2(seatLayout [][]string) int {
	return findStabilisingValue(seatLayout, getNextSeatPt2)
}

func findStabilisingValue(seatLayout [][]string, getNextSeat func(y int, x int, seatLayout [][]string) string) int {
	for true {
		nextLayout := getNextLayout(seatLayout, getNextSeat)
		if isSeatLayoutEqual(seatLayout, nextLayout) {
			break
		}
		seatLayout = nextLayout
	}

	return countNoOfOccupiedSeats(seatLayout)
}

func getNextLayout(seatLayout [][]string, getNextSeat func(y int, x int, seatLayout [][]string) string) [][]string {
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

func occupiedAdjacentCountPt1(y int, x int, seatLayout [][]string) int {
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

func getNextSeatPt1(y int, x int, seatLayout [][]string) string {
	currentSeat := seatLayout[y][x]
	if currentSeat == "." {
		return "."
	}
	occupiedCount := occupiedAdjacentCountPt1(y, x, seatLayout)
	if currentSeat == "L" && occupiedCount == 0 {
		return "#"
	}
	if currentSeat == "#" && occupiedCount >= 4 {
		return "L"
	}

	return currentSeat
}

type Dir struct {
	y int
	x int
}

func occupiedAdjacentCountPt2(y int, x int, seatLayout [][]string) int {
	occupiedCount := 0
	routes := getRoutes(y, x, seatLayout)
	for _, route := range routes {
		for _, dir := range route {
			if seatLayout[dir.y][dir.x] == "L" {
				break
			}
			if seatLayout[dir.y][dir.x] == "#" {
				occupiedCount++
				break
			}
		}
	}

	return occupiedCount
}

func getRoutes(y int, x int, seatLayout [][]string) [][]Dir {
	dirs := []Dir{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	routes := make([][]Dir, len(dirs))
	for i, dir := range dirs {
		var route []Dir
		for i := 1; y+i*dir.y < len(seatLayout) && y+i*dir.y >= 0 && x+i*dir.x < len(seatLayout[0]) && x+i*dir.x >= 0; i++ {
			route = append(route, Dir{y + i*dir.y, x + i*dir.x})
		}
		routes[i] = route
	}

	return routes
}

func getNextSeatPt2(y int, x int, seatLayout [][]string) string {
	currentSeat := seatLayout[y][x]
	if currentSeat == "." {
		return "."
	}
	occupiedCount := occupiedAdjacentCountPt2(y, x, seatLayout)
	if currentSeat == "L" && occupiedCount == 0 {
		return "#"
	}
	if currentSeat == "#" && occupiedCount >= 5 {
		return "L"
	}

	return currentSeat
}

func seatLayoutToString(seatLayout [][]string) string {
	stringSeatLayout := ""
	for y := range seatLayout {
		stringSeatLayout = stringSeatLayout + strings.Join(seatLayout[y], "") + "\n"
	}

	return stringSeatLayout
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
	fmt.Println("Pt2")
	fmt.Println(FindStabilisingValuePt2(parse(loadFile())))
}
