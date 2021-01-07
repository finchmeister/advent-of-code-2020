package main

import (
	"testing"
)

var testInput = `sesenwnenenewseeswwswswwnenewsewsw
neeenesenwnwwswnenewnwwsewnenwseswesw
seswneswswsenwwnwse
nwnwneseeswswnenewneswwnewseswneseene
swweswneswnenwsewnwneneseenw
eesenwseswswnenwswnwnwsewwnwsene
sewnenenenesenwsewnenwwwse
wenwwweseeeweswwwnwwe
wsweesenenewnwwnwsenewsenwwsesesenwne
neeswseenwwswnwswswnw
nenwswwsewswnenenewsenwsenwnesesenew
enewnwewneswsewnwswenweswnenwsenwsw
sweneswneswneneenwnewenewwneswswnese
swwesenesewenwneswnwwneseswwne
enesenwswwswneneswsenwnewswseenwsese
wnwnesenesenenwwnenwsewesewsesesew
nenewswnwewswnenesenwnesewesw
eneswnwswnwsenenwnwnwwseeswneewsenese
neswnwewnwnwseenwseesewsenwsweewe
wseweeenwnesenwwwswnew`

type Fixtures struct {
	input    int
	expected int
}

func TestFindHowManyTilesLeftWithBlackSideUpPt1(t *testing.T) {
	got := FindHowManyTilesLeftWithBlackSideUpPt1(parse(testInput))
	expected := 10

	if got != expected {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}

func TestFindHowManyTilesLeftWithBlackSideUpPt2(t *testing.T) {
	fixtures := []Fixtures{
		{1, 15},
		{2, 12},
		{3, 25},
		{4, 14},
		{5, 23},
		{6, 28},
		{7, 41},
		{8, 37},
		{9, 49},
		{10, 37},
		{20, 132},
		{30, 259},
		{40, 406},
		{50, 566},
		{60, 788},
		{70, 1106},
		{80, 1373},
		{90, 1844},
		{100, 2208},
	}

	for _, fixture := range fixtures {
		got := FindHowManyTilesLeftWithBlackSideUpPt2(parse(testInput), fixture.input)
		expected := fixture.expected

		if got != expected {
			t.Errorf(
				"Expected %d; got %d",
				fixture.expected,
				got,
			)
		}
	}
}
