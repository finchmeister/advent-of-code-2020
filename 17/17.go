package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func FindCubesInActiveStatePt1(cubes map[int]map[int]map[int]bool, cycles int) int {
	for i := 0; i < cycles; i++ {
		cubes = getNextStatePt1(cubes)
	}

	activeCubeCount := getActiveCubeCountPt1(cubes)

	return activeCubeCount
}

func getNextStatePt1(cubes map[int]map[int]map[int]bool) map[int]map[int]map[int]bool {
	nextCubes := make(map[int]map[int]map[int]bool)
	expandedCubes := expandRangePt1(cubes)
	for z := range expandedCubes {
		nextCubes[z] = make(map[int]map[int]bool)
		for y := range expandedCubes[z] {
			nextCubes[z][y] = make(map[int]bool)
			for x := range expandedCubes[z][y] {
				nextCubes[z][y][x] = getNextCubeStatePt1(x, y, z, cubes)
			}
		}
	}

	return nextCubes
}

func expandRangePt1(cubes map[int]map[int]map[int]bool) map[int]map[int]map[int]bool {

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

func getActiveCubeCountPt1(cubes map[int]map[int]map[int]bool) int {
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

func getNextCubeStatePt1(x int, y int, z int, cubes map[int]map[int]map[int]bool) bool {
	activeNeighbours := countActiveNeighboursPt1(x, y, z, cubes)
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

func countActiveNeighboursPt1(x int, y int, z int, cubes map[int]map[int]map[int]bool) int {
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

func parsePt1(input string) map[int]map[int]map[int]bool {
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

func FindCubesInActiveStatePt2(cubes map[int]map[int]map[int]map[int]bool, cycles int) int {
	for i := 0; i < cycles; i++ {
		cubes = getNextStatePt2(cubes)
	}

	activeCubeCount := getActiveCubeCountPt2(cubes)

	return activeCubeCount
}

func getNextStatePt2(cubes map[int]map[int]map[int]map[int]bool) map[int]map[int]map[int]map[int]bool {
	nextCubes := make(map[int]map[int]map[int]map[int]bool)
	expandedCubes := expandRangePt2(cubes)

	for w := range expandedCubes {
		nextCubes[w] = make(map[int]map[int]map[int]bool)
		for z := range expandedCubes[w] {
			nextCubes[w][z] = make(map[int]map[int]bool)
			for y := range expandedCubes[w][z] {
				nextCubes[w][z][y] = make(map[int]bool)
				for x := range expandedCubes[w][z][y] {
					nextCubes[w][z][y][x] = getNextCubeStatePt2(x, y, z, w, cubes)
				}
			}
		}
	}

	return nextCubes
}

func expandRangePt2(cubes map[int]map[int]map[int]map[int]bool) map[int]map[int]map[int]map[int]bool {

	minX := 0
	maxX := 0
	minY := 0
	maxY := 0
	minZ := 0
	maxZ := 0
	minW := 0
	maxW := 0

	for w := range cubes {
		if w > maxW {
			maxW = w
		}
		if w < minW {
			minW = w
		}
	}

	for z := range cubes[0] {
		if z > maxZ {
			maxZ = z
		}
		if z < minZ {
			minZ = z
		}
	}

	for y := range cubes[0][0] {
		if y > maxY {
			maxY = y
		}
		if y < minY {
			minY = y
		}
	}

	for x := range cubes[0][0][0] {
		if x > maxX {
			maxX = x
		}
		if x < minX {
			minX = x
		}
	}

	expandedCubes := make(map[int]map[int]map[int]map[int]bool)

	for w := minW - 1; w <= maxW+1; w++ {
		expandedCubes[w] = make(map[int]map[int]map[int]bool)
		for z := minZ - 1; z <= maxZ+1; z++ {
			expandedCubes[w][z] = make(map[int]map[int]bool)
			for y := minY - 1; y <= maxY+1; y++ {
				expandedCubes[w][z][y] = make(map[int]bool)
				for x := minX - 1; x <= maxX+1; x++ {
					if cubes[w][z][y][x] == true {
						expandedCubes[w][z][y][x] = true
						continue
					}
					expandedCubes[w][z][y][x] = false
				}
			}
		}
	}

	return expandedCubes
}

func getActiveCubeCountPt2(cubes map[int]map[int]map[int]map[int]bool) int {
	activeCubeCount := 0
	for w := range cubes {
		for z := range cubes[w] {
			for y := range cubes[w][z] {
				for x := range cubes[w][z][y] {
					if cubes[w][z][y][x] {
						activeCubeCount++
					}
				}
			}
		}
	}

	return activeCubeCount
}

func getNextCubeStatePt2(x int, y int, z int, w int, cubes map[int]map[int]map[int]map[int]bool) bool {
	activeNeighbours := countActiveNeighboursPt2(x, y, z, w, cubes)
	isCubeActive := cubes[w][z][y][x] == true

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

func countActiveNeighboursPt2(x int, y int, z int, w int, cubes map[int]map[int]map[int]map[int]bool) int {
	activeNeighbours := 0

	for wI := w - 1; wI <= w+1; wI++ {
		for zI := z - 1; zI <= z+1; zI++ {
			for yI := y - 1; yI <= y+1; yI++ {
				for xI := x - 1; xI <= x+1; xI++ {
					if xI == x && yI == y && zI == z && wI == w {
						continue
					}
					if cubes[wI][zI][yI][xI] == true {
						activeNeighbours++
					}
				}
			}
		}
	}

	return activeNeighbours
}

func parsePt2(input string) map[int]map[int]map[int]map[int]bool {
	rows := strings.Split(input, "\n")
	cubes := make(map[int]map[int]map[int]map[int]bool)
	cubes[0] = make(map[int]map[int]map[int]bool)
	cubes[0][0] = make(map[int]map[int]bool)
	for y, row := range rows {
		yRow := make(map[int]bool)
		for x, cell := range strings.Split(row, "") {
			yRow[x] = cell == "#"
		}
		cubes[0][0][y] = yRow
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
	fmt.Println(FindCubesInActiveStatePt1(parsePt1(loadFile()), 6))
	fmt.Println("Pt2")
	fmt.Println(FindCubesInActiveStatePt2(parsePt2(loadFile()), 6))
}
