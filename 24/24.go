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

var directionCoordMap = map[string]Coord{
	"ne": {1, 1},
	"e":  {2, 0},
	"se": {1, -1},
	"sw": {-1, -1},
	"w":  {-2, 0},
	"nw": {-1, 1},
}

func FindHowManyTilesLeftWithBlackSideUpPt1(allSteps [][]string) int {
	return getNoOfBlackTiles(computeInitialTilingState(allSteps))
}

func FindHowManyTilesLeftWithBlackSideUpPt2(allSteps [][]string, noOfDays int) int {
	tilingState := computeInitialTilingState(allSteps)
	for i := 0; i < noOfDays; i++ {
		tilingState = getNextTilingState(tilingState)
	}

	return getNoOfBlackTiles(tilingState)
}

func computeInitialTilingState(allSteps [][]string) map[int]map[int]bool {
	result := make(map[int]map[int]bool)

	for _, steps := range allSteps {
		position := getPositionAfterSteps(Coord{0, 0}, steps)

		if result[position.x] == nil {
			result[position.x] = make(map[int]bool)
		}

		result[position.x][position.y] = !result[position.x][position.y]
	}

	return result
}

func expandTilingState(currentTiles map[int]map[int]bool) map[int]map[int]bool {
	expandedTiles := make(map[int]map[int]bool)

	minX := 0
	maxX := 0
	minY := 0
	maxY := 0

	for x := range currentTiles {
		if x > maxX {
			maxX = x
		}
		if x < minX {
			minX = x
		}

		for y := range currentTiles[x] {
			if y > maxY {
				maxY = y
			}
			if y < minY {
				minY = y
			}
		}
	}

	for x := minX - 2; x <= maxX+2; x++ {
		if expandedTiles[x] == nil {
			expandedTiles[x] = make(map[int]bool)
		}
		for y := minY - 1; y <= maxY+1; y++ {
			if mod(x, 2) == 0 {
				if mod(y, 2) == 0 {
					expandedTiles[x][y] = isTileBlack(currentTiles, Coord{x, y})
				}
			} else {
				if mod(y, 2) == 1 {
					expandedTiles[x][y] = isTileBlack(currentTiles, Coord{x, y})
				}
			}
		}
	}

	return expandedTiles
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func getNextTilingState(currentTiles map[int]map[int]bool) map[int]map[int]bool {
	nextTiles := make(map[int]map[int]bool)
	currentTiles = expandTilingState(currentTiles)
	for x := range currentTiles {
		if nextTiles[x] == nil {
			nextTiles[x] = make(map[int]bool)
		}
		for y := range currentTiles[x] {
			nextTiles[x][y] = getNextTileState(currentTiles, Coord{x, y})
		}
	}

	return nextTiles
}

func getNextTileState(currentTiles map[int]map[int]bool, position Coord) bool {
	adjacentBlack := 0
	for _, coords := range directionCoordMap {
		if isTileBlack(currentTiles, Coord{position.x + coords.x, position.y + coords.y}) {
			adjacentBlack++
		}
	}

	isTileBlack := isTileBlack(currentTiles, position)
	if isTileBlack {
		if adjacentBlack == 0 || adjacentBlack > 2 {
			return false
		}
	} else {
		if adjacentBlack == 2 {
			return true
		}
	}

	return isTileBlack
}

func isTileBlack(currentTiles map[int]map[int]bool, position Coord) bool {
	if currentTiles[position.x] == nil || currentTiles[position.x][position.y] == false {
		return false
	}

	return currentTiles[position.x][position.y]
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
		curPos = Coord{curPos.x + directionCoordMap[step].x, curPos.y + directionCoordMap[step].y}
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
	fmt.Println("Pt2")
	fmt.Println(FindHowManyTilesLeftWithBlackSideUpPt2(parse(loadFile()), 100))
}
