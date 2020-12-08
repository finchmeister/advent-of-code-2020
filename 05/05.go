package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

type BoardingPass struct {
	row    int
	column int
	seatId int
}

func FindHighestSeatIdPt1(boardingPasses []string) int {
	return getMax(getAllSeatIds(boardingPasses))
}

func FindMissingSeatIdPt2(boardingPasses []string) int {
	seatIds := getAllSeatIds(boardingPasses)
	for i := getMin(seatIds) + 1; i < getMax(seatIds)-1; i++ {
		if contains(i-1, seatIds) == true && contains(i+1, seatIds) == true && contains(i, seatIds) == false {
			return i
		}
	}

	return 0
}

func getAllSeatIds(boardingPasses []string) []int {
	var seatIds []int
	for _, boardingPass := range boardingPasses {
		r := calculateBoardingPass(boardingPass)
		seatIds = append(seatIds, r.seatId)
	}

	return seatIds
}

func getMin(values []int) int {
	min := values[0]
	for _, v := range values {
		if v < min {
			min = v
		}
	}

	return min
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

func contains(value int, slice []int) bool {
	for _, v := range slice {
		if value == v {
			return true
		}
	}

	return false
}

func calculateBoardingPass(rawBoardingPass string) BoardingPass {
	row := binarySpacePartition(rawBoardingPass[:7], "F")
	column := binarySpacePartition(rawBoardingPass[7:], "L")
	seatId := row*8 + column

	return BoardingPass{row, column, seatId}
}

func binarySpacePartition(input string, lowerHalf string) int {
	max := math.Pow(float64(2), float64(len(input))) - 1
	min := float64(0)
	var value int
	for _, c := range input {
		half := math.Ceil((max - min) / 2)
		if string(c) == lowerHalf {
			max = max - half
			value = int(max)
		} else {
			min = min + half
			value = int(min)
		}
	}

	return value
}

func loadBoardingPasses() []string {
	data, err := ioutil.ReadFile("05_input.txt")
	if err != nil {
		panic(err)
	}

	return strings.Split(string(data), "\n")
}

func main() {
	fmt.Println("Pt1 Highest Seating Id")
	fmt.Println(FindHighestSeatIdPt1(loadBoardingPasses()))
	fmt.Println("Pt2 Find Missing Seating Id")
	fmt.Println(FindMissingSeatIdPt2(loadBoardingPasses()))
}
