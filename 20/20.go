package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

type Tile struct {
	tile   [][]string
	number int
}

type Side struct {
	x int
	y int
}

var sides = []Side{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func FindCornerIdsMultipliedPt1(tiles []Tile) int {
	return getCornerValuesMultiplied(solvePuzzle(tiles))
}

func FindSeaMonstersWaterRoughnessPt2(tiles []Tile) int {
	solvedPuzzle := solvePuzzle(tiles)

	minY, maxY, minX, maxX := getSolvedPuzzleMinMaxXY(solvedPuzzle)

	var puzzleWithoutBorders [][][][]string

	for y := maxY; y >= minY; y-- {
		var puzzleWithoutBordersX [][][]string
		for x := minX; x <= maxX; x++ {
			puzzleWithoutBordersX = append(puzzleWithoutBordersX, removeBorder(solvedPuzzle[y][x].tile))
		}
		puzzleWithoutBorders = append(puzzleWithoutBorders, puzzleWithoutBordersX)
	}

	var image [][]string
	for y := range puzzleWithoutBorders {

		for tileY := range puzzleWithoutBorders[y][0] {
			var imageX []string
			for x := range puzzleWithoutBorders[y] {
				imageX = append(imageX, puzzleWithoutBorders[y][x][tileY]...)
			}
			image = append(image, imageX)
		}
	}

	for i := 0; i < 2; i++ {
		if i == 1 {
			image = flipVertically(image)
		}
		for j := 0; j < 4; j++ {
			image = rotate90CW(image)

			found := false
			for _, coord := range getSeaMonsterCoords(image) {
				image = updateImageWithSeaMonster(coord[0], coord[1], image)
				//fmt.Println(imageToString(image))
				found = true
			}

			if found {
				return getTotalWaterRoughness(image)
			}
		}
	}

	panic("Solution not found")
}

func imageToString(image [][]string) string {
	imageAsString := ""
	for y := range image {
		imageAsString = imageAsString + strings.Join(image[y], "") + "\n"
	}

	return imageAsString
}

func getTotalWaterRoughness(image [][]string) int {
	totalWaterRoughness := 0
	for y := range image {
		for x := range image[y] {
			if image[y][x] == "#" {
				totalWaterRoughness++
			}
		}
	}

	return totalWaterRoughness
}

func updateImageWithSeaMonster(x int, y int, image [][]string) [][]string {
	image[y+1][x+0] = "O"
	image[y+2][x+1] = "O"
	image[y+2][x+4] = "O"
	image[y+1][x+5] = "O"
	image[y+1][x+6] = "O"
	image[y+2][x+7] = "O"
	image[y+2][x+10] = "O"
	image[y+1][x+11] = "O"
	image[y+1][x+12] = "O"
	image[y+2][x+13] = "O"
	image[y+2][x+16] = "O"
	image[y+1][x+17] = "O"
	image[y+1][x+18] = "O"
	image[y+0][x+18] = "O"
	image[y+1][x+19] = "O"

	return image
}

func getSeaMonsterCoords(image [][]string) [][]int {
	var seaMonsterCoords [][]int
	for y := 0; y < len(image)-2; y++ {
		for x := 0; x < len(image[0])-19; x++ {
			if containsSeaMonster(x, y, image) {
				seaMonsterCoords = append(seaMonsterCoords, []int{x, y})
			}
		}

	}
	return seaMonsterCoords
}

func containsSeaMonster(x int, y int, image [][]string) bool {
	return image[y+1][x+0] == "#" &&
		image[y+2][x+1] == "#" &&
		image[y+2][x+4] == "#" &&
		image[y+1][x+5] == "#" &&
		image[y+1][x+6] == "#" &&
		image[y+2][x+7] == "#" &&
		image[y+2][x+10] == "#" &&
		image[y+1][x+11] == "#" &&
		image[y+1][x+12] == "#" &&
		image[y+2][x+13] == "#" &&
		image[y+2][x+16] == "#" &&
		image[y+1][x+17] == "#" &&
		image[y+1][x+18] == "#" &&
		image[y+0][x+18] == "#" &&
		image[y+1][x+19] == "#"
}

func removeBorder(tile [][]string) [][]string {
	var borderRemoved [][]string

	for y := 1; y < len(tile)-1; y++ {
		borderRemoved = append(borderRemoved, tile[y][1:len(tile[y])-1])
	}

	return borderRemoved
}

func solvePuzzle(tiles []Tile) map[int]map[int]Tile {
	solvedPuzzle := make(map[int]map[int]Tile)
	var tilesToCheck []Tile

	remainingTiles := make([]Tile, len(tiles))

	copy(remainingTiles, tiles)

	tile, remainingTiles := tilePop(remainingTiles)

	solvedPuzzle[0] = make(map[int]Tile)
	solvedPuzzle[0][0] = tile
	tilesToCheck = append(tilesToCheck, tile)

	for len(remainingTiles) > 0 {

		for _, tileToCheck := range tilesToCheck {

			for _, side := range sides {
				for _, remainingTile := range remainingTiles {

					doesTileFitSide, adjustedTile := doesTileFitSide(remainingTile, tileToCheck, side)

					if doesTileFitSide {
						adjustedTileX, adjustedTileY := getTileXY(solvedPuzzle, tileToCheck)
						adjustedTileX = adjustedTileX + side.x
						adjustedTileY = adjustedTileY + side.y

						if solvedPuzzle[adjustedTileY] == nil {
							solvedPuzzle[adjustedTileY] = make(map[int]Tile)
						}
						solvedPuzzle[adjustedTileY][adjustedTileX] = adjustedTile
						tilesToCheck = append(tilesToCheck, adjustedTile)
						remainingTiles = removeTile(remainingTiles, remainingTile)
					}
				}
			}
		}
	}

	return solvedPuzzle
}

func getCornerValuesMultiplied(solvedPuzzle map[int]map[int]Tile) int {
	minY, maxY, minX, maxX := getSolvedPuzzleMinMaxXY(solvedPuzzle)

	for y := range solvedPuzzle {
		if y < minY {
			minY = y
		}
		if y > maxY {
			maxY = y
		}
	}
	for x := range solvedPuzzle[0] {
		if x < minX {
			minX = x
		}
		if x > maxX {
			maxX = x
		}
	}

	return solvedPuzzle[minY][minX].number *
		solvedPuzzle[minY][maxX].number *
		solvedPuzzle[maxY][minX].number *
		solvedPuzzle[maxY][maxX].number
}

func getSolvedPuzzleMinMaxXY(solvedPuzzle map[int]map[int]Tile) (int, int, int, int) {
	minY := 0
	maxY := 0
	minX := 0
	maxX := 0
	for y := range solvedPuzzle {
		if y < minY {
			minY = y
		}
		if y > maxY {
			maxY = y
		}
	}
	for x := range solvedPuzzle[0] {
		if x < minX {
			minX = x
		}
		if x > maxX {
			maxX = x
		}
	}

	return minY, maxY, minX, maxX
}

func getTileXY(solvedPuzzle map[int]map[int]Tile, tile Tile) (int, int) {
	for y := range solvedPuzzle {
		for x := range solvedPuzzle[y] {
			if tile.number == solvedPuzzle[y][x].number {
				return x, y
			}
		}
	}
	panic("Cannot be called if tile is not in puzzle")
}

func doesTileFitSide(testTile Tile, sourceTile Tile, side Side) (bool, Tile) {
	sourceTileEdge := getEdge(side, sourceTile)
	for i := 0; i < 2; i++ {
		if i == 1 {
			testTile = flipTileVertically(testTile)
		}
		for j := 0; j < 4; j++ {
			testTile = rotateTile90CW(testTile)

			testTileEdge := getEdge(getComplimentingSide(side), testTile)
			if reflect.DeepEqual(sourceTileEdge, testTileEdge) {
				return true, testTile
			}
		}
	}

	return false, testTile
}

func getComplimentingSide(side Side) Side {
	return Side{side.x * -1, side.y * -1}
}

func tilePop(tiles []Tile) (Tile, []Tile) {
	return tiles[len(tiles)-1], tiles[:len(tiles)-1]
}

func removeTile(tiles []Tile, removeTile Tile) []Tile {
	for i := range tiles {
		if tiles[i].number == removeTile.number {
			return append(tiles[:i], tiles[i+1:]...)
		}
	}
	return tiles
}

func getEdge(side Side, tile Tile) []string {
	x := side.x
	y := side.y

	if x == 0 && y == 1 {
		return tile.tile[0]
	}

	if x == 0 && y == -1 {
		return tile.tile[len(tile.tile)-1]
	}

	if x == 1 && y == 0 {
		var edge []string
		for yI := range tile.tile {
			edge = append(edge, tile.tile[yI][len(tile.tile[0])-1])
		}
		return edge
	}
	if x == -1 && y == 0 {
		var edge []string
		for yI := range tile.tile {
			edge = append(edge, tile.tile[yI][0])
		}
		return edge
	}

	panic(fmt.Sprintf("Invalid args: x: %v, y: %v", x, y))
}

func flipTileVertically(tile Tile) Tile {
	return Tile{
		flipVertically(tile.tile),
		tile.number,
	}
}

func flipVertically(s [][]string) [][]string {
	var flippedS [][]string

	for y := range s {
		var flippedRow []string
		for x := len(s[0]) - 1; x >= 0; x-- {
			flippedRow = append(flippedRow, s[y][x])
		}
		flippedS = append(flippedS, flippedRow)
	}

	return flippedS
}

func rotateTile90CW(tile Tile) Tile {
	return Tile{
		rotate90CW(tile.tile),
		tile.number,
	}
}

func rotate90CW(s [][]string) [][]string {
	var rotatedS [][]string

	for y := range s {
		var rotatedRow []string
		for x := len(s[0]) - 1; x >= 0; x-- {
			rotatedRow = append(rotatedRow, s[x][y])
		}
		rotatedS = append(rotatedS, rotatedRow)
	}

	return rotatedS
}

func parse(input string) []Tile {
	puzzlePieces := strings.Split(input, "\n\n")

	var tiles []Tile

	for _, rawTile := range puzzlePieces {
		tiles = append(tiles, parseTile(rawTile))
	}

	return tiles
}

func parseTile(rawTile string) Tile {
	rows := strings.Split(rawTile, "\n")
	tileNumber, _ := strconv.Atoi(rows[0][5:9])
	var tile [][]string

	for i := 1; i < len(rows); i++ {
		value := strings.Split(rows[i], "")
		tile = append(tile, value)
	}

	return Tile{
		tile,
		tileNumber,
	}
}

func loadFile() string {
	data, err := ioutil.ReadFile("20_input.txt")
	if err != nil {
		panic(err)
	}

	return string(data)
}

func main() {
	fmt.Println("Pt1")
	fmt.Println(FindCornerIdsMultipliedPt1(parse(loadFile())))
	fmt.Println("Pt2")
	fmt.Println(FindSeaMonstersWaterRoughnessPt2(parse(loadFile())))
}
