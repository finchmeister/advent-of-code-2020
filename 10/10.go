package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

func FindDifferencesMultipliedPt1(data []int) int {
	data = append(data, 0, getMax(data)+3)
	sort.Ints(data)
	differences := make(map[int]int)
	for i := 0; i < len(data)-1; i++ {
		difference := data[i+1] - data[i]
		differences[difference]++
	}
	return differences[1] * differences[3]
}

func FindCountValidPermutationsPt2(data []int) int {
	data = append(data, 0, getMax(data)+3)
	sort.Ints(data)
	islands := findIslands(data)
	multiple := 1
	for island, count := range islands {
		multiple = multiple * getIslandPermutations(island, count)
	}

	return multiple
}

func getIslandPermutations(island int, pow int) int {
	var base int
	if island == 2 {
		base = 2
	}
	if island == 3 {
		base = 4
	}
	if island == 4 {
		base = 7
	}

	return int(math.Pow(float64(base), float64(pow)))
}

func findIslands(data []int) map[int]int {
	data = append(data, 0, getMax(data)+3)
	sort.Ints(data)

	islands := make(map[int]int)
	i := 0
	for i < len(data)-2 {
		j := 0
		for data[i+j]+1 == data[i+j+1] {
			j++
		}
		i = i + j + 1
		if j > 1 {
			islands[j]++
		}
	}

	return islands
}

func getMax(values []int) int {
	max := values[0]
	for _, v := range values {
		if v > max {
			max = v
		}
	}

	return max
}

func parse(input string) []int {
	var data []int
	for _, row := range splitByNewLine(input) {
		value, _ := strconv.Atoi(row)
		data = append(data, value)
	}

	return data
}

func splitByNewLine(input string) []string {
	return strings.Split(input, "\n")
}

func loadFile() string {
	data, err := ioutil.ReadFile("10_input.txt")
	if err != nil {
		panic(err)
	}

	return string(data)
}

func main() {
	fmt.Println("Pt1")
	fmt.Println(FindDifferencesMultipliedPt1(parse(loadFile())))
	fmt.Println("Pt2")
	fmt.Println(FindCountValidPermutationsPt2(parse(loadFile())))
}
