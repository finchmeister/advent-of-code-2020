package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type Notes struct {
	timestamp int
	busIds    []string
}

func FindEarliestBusMultipliedPt1(notes Notes) int {
	min := math.MaxInt32
	minBusId := 1
	for _, busIdRaw := range notes.busIds {
		if busIdRaw == "x" {
			continue
		}
		busId, _ := strconv.Atoi(busIdRaw)

		v := getNextBusIn(notes.timestamp, busId)

		if v < min {
			min = v
			minBusId = busId
		}
	}

	return min * minBusId
}

func getNextBusIn(timestamp int, busId int) int {
	a := int(math.Ceil(float64(timestamp)/float64(busId))) * busId
	return a % timestamp
}
func parse(input string) Notes {
	timestamp, _ := strconv.Atoi(splitByNewLine(input)[0])

	return Notes{
		timestamp,
		splitByComma(splitByNewLine(input)[1]),
	}
}

func splitByNewLine(input string) []string {
	return strings.Split(input, "\n")
}

func splitByComma(input string) []string {
	return strings.Split(input, ",")
}

func loadFile() string {
	data, err := ioutil.ReadFile("13_input.txt")
	if err != nil {
		panic(err)
	}

	return string(data)
}

func main() {
	fmt.Println("Pt1")
	fmt.Println(FindEarliestBusMultipliedPt1(parse(loadFile())))
}
