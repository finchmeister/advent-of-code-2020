package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func FindCubesInActiveStatePt1(cubes map[int]map[int]map[int]bool, cycles int) int {
	for i := 0; i < cycles; i++ {
		cubes = getNextState(cubes)
	}

	activeCubeCount := getActiveCubeCount(cubes)

	return activeCubeCount
}

func getNextState(cubes map[int]map[int]map[int]bool) map[int]map[int]map[int]bool {
	nextCubes := make(map[int]map[int]map[int]bool)
	expandedCubes := expandRange(cubes)
	for z := range expandedCubes {
		nextCubes[z] = make(map[int]map[int]bool)
		for y := range expandedCubes[z] {
			nextCubes[z][y] = make(map[int]bool)
			for x := range expandedCubes[z][y] {
				nextCubes[z][y][x] = getNextCubeState(x, y, z, cubes)
			}
		}
	}

	return nextCubes
}

func expandRange(cubes map[int]map[int]map[int]bool) map[int]map[int]map[int]bool {

	minX := 0
	maxX := 0
	minY := 0
	maxY := 0
	minZ := 0
	maxZ := 0

	for z := range cubes {
		if z > maxZ {
			maxZ = z
		}
		if z < minZ {
			minZ = z
		}
	}

	for y := range cubes[0] {
		if y > maxY {
			maxY = y
		}
		if y < minY {
			minY = y
		}
	}

	for x := range cubes[0][0] {
		if x > maxX {
			maxX = x
		}
		if x < minX {
			minX = x
		}
	}

	expandedCubes := make(map[int]map[int]map[int]bool)
	for z := minZ - 1; z <= maxZ+1; z++ {
		expandedCubes[z] = make(map[int]map[int]bool)
		for y := minY - 1; y <= maxY+1; y++ {
			expandedCubes[z][y] = make(map[int]bool)
			for x := minX - 1; x <= maxX+1; x++ {
				if cubes[z][y][x] == true {
					expandedCubes[z][y][x] = true
					continue
				}
				expandedCubes[z][y][x] = false
			}
		}
	}

	return expandedCubes
}

func getActiveCubeCount(cubes map[int]map[int]map[int]bool) int {
	activeCubeCount := 0
	for z := range cubes {
		for y := range cubes[z] {
			for x := range cubes[z][y] {
				if cubes[z][y][x] {
					activeCubeCount++
				}
			}
		}
	}

	return activeCubeCount
}

func getNextCubeState(x int, y int, z int, cubes map[int]map[int]map[int]bool) bool {
	activeNeighbours := countActiveNeighbours(x, y, z, cubes)
	isCubeActive := cubes[z][y][x] == true

	if isCubeActive {
		if activeNeighbours == 2 || activeNeighbours == 3 {
			return true
		}
		return false
	}

	if activeNeighbours == 3 {
		return true
	}

	return false
}

func countActiveNeighbours(x int, y int, z int, cubes map[int]map[int]map[int]bool) int {
	activeNeighbours := 0

	for zI := z - 1; zI <= z+1; zI++ {
		for yI := y - 1; yI <= y+1; yI++ {
			for xI := x - 1; xI <= x+1; xI++ {
				if xI == x && yI == y && zI == z {
					continue
				}
				if cubes[zI][yI][xI] == true {
					activeNeighbours++
				}
			}
		}
	}

	return activeNeighbours
}

func parse(input string) map[int]map[int]map[int]bool {
	rows := strings.Split(input, "\n")
	cubes := make(map[int]map[int]map[int]bool)
	cubes[0] = make(map[int]map[int]bool)
	for y, row := range rows {
		yRow := make(map[int]bool)
		for x, cell := range strings.Split(row, "") {
			yRow[x] = cell == "#"
		}
		cubes[0][y] = yRow
	}

	return cubes
}

func loadFile() string {
	data, err := ioutil.ReadFile("17_input.txt")
	if err != nil {
		panic(err)
	}

	return string(data)
}

func main() {
	fmt.Println("Pt1")
	fmt.Println(FindCubesInActiveStatePt1(parse(loadFile()), 6))
}
