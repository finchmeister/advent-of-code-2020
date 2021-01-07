package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Coord struct {
	x int
	y int
}

var directionCoord = map[string]Coord{
	"ne": {1, 2},
	"e":  {2, 0},
	"se": {1, -2},
	"sw": {-1, -2},
	"w":  {-2, 0},
	"nw": {-1, 2},
}

func FindHowManyTilesLeftWithBlackSideUpPt1(allSteps [][]string) int {
	result := make(map[int]map[int]bool)

	for _, steps := range allSteps {
		position := getPositionAfterSteps(Coord{0, 0}, steps)

		if result[position.x] == nil {
			result[position.x] = make(map[int]bool)
		}

		result[position.x][position.y] = !result[position.x][position.y]
	}

	return getNoOfBlackTiles(result)
}

func getNoOfBlackTiles(tiles map[int]map[int]bool) int {
	countBlack := 0

	for x := range tiles {
		for y := range tiles[x] {
			if tiles[x][y] {
				countBlack++
			}
		}
	}

	return countBlack
}

func getPositionAfterSteps(startPosition Coord, steps []string) Coord {
	curPos := startPosition
	for _, step := range steps {
		curPos = Coord{curPos.x + directionCoord[step].x, curPos.y + directionCoord[step].y}
	}

	return curPos
}

func parseSteps(rawSteps string) []string {
	var steps []string
	for i := 0; i < len(rawSteps); i++ {
		step := string(rawSteps[i])
		if step != "e" && step != "w" {
			step = step + string(rawSteps[i+1])
			i++
		}
		steps = append(steps, step)
	}

	return steps
}

func parse(input string) [][]string {
	var allSteps [][]string
	for _, rawSteps := range strings.Split(input, "\n") {
		allSteps = append(allSteps, parseSteps(rawSteps))
	}

	return allSteps
}

func loadFile() string {
	data, err := ioutil.ReadFile("24_input.txt")
	if err != nil {
		panic(err)
	}

	return string(data)
}

func main() {
	fmt.Println("Pt1")
	fmt.Println(FindHowManyTilesLeftWithBlackSideUpPt1(parse(loadFile())))
}
