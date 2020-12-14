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

// Credit to https://www.dave4math.com/mathematics/chinese-remainder-theorem
func FindEarliestBusTimestampPt2(notes Notes) int {
	n := make(map[int]int)
	a := make(map[int]int)
	i := 0
	for j, busIdRaw := range notes.busIds {
		if busIdRaw != "x" {
			busId, _ := strconv.Atoi(busIdRaw)
			n[i] = busId
			a[i] = -j
			i++
		}
	}

	N := getMultiple(n)

	nHat := make(map[int]int)
	for i := range n {
		nHat[i] = N / n[i]
	}

	u := make(map[int]int)
	for i := range n {
		u[i] = getU(nHat[i], n[i])
	}

	toSum := make(map[int]int)

	for i := range n {
		toSum[i] = a[i] * nHat[i] * u[i]
	}

	return mod(getSum(toSum), N)
}

func getU(nHat int, nI int) int {
	i := 0
	for (nHat*i)%nI != 1 {
		i++
	}

	return i
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func getMultiple(values map[int]int) int {
	x := 1
	for _, i := range values {
		x = x * i
	}

	return x
}

func getSum(values map[int]int) int {
	x := 0
	for _, i := range values {
		x = x + i
	}

	return x
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
	fmt.Println("Pt2")
	fmt.Println(FindEarliestBusTimestampPt2(parse(loadFile())))
}
