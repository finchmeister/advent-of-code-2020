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

	return solvedPuzzle[minY][minX].number *
		solvedPuzzle[minY][maxX].number *
		solvedPuzzle[maxY][minX].number *
		solvedPuzzle[maxY][maxX].number
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
	var flippedTile [][]string

	for y := range tile.tile {
		var flippedTileRow []string
		for x := len(tile.tile[0]) - 1; x >= 0; x-- {
			flippedTileRow = append(flippedTileRow, tile.tile[y][x])
		}
		flippedTile = append(flippedTile, flippedTileRow)
	}

	return Tile{
		flippedTile,
		tile.number,
	}
}

func rotateTile90CW(tile Tile) Tile {
	var flippedTile [][]string

	for y := range tile.tile {
		var flippedTileRow []string
		for x := len(tile.tile[0]) - 1; x >= 0; x-- {
			flippedTileRow = append(flippedTileRow, tile.tile[x][y])
		}
		flippedTile = append(flippedTile, flippedTileRow)
	}

	return Tile{
		flippedTile,
		tile.number,
	}
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
}
