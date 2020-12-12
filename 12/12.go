package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var compass = []string{"N", "E", "S", "W"}

type Instruction struct {
	action string
	value  int
}

type Position struct {
	n int
	e int
}

func FindManhattanDistancePt1(instructions []Instruction) int {
	currentDir := "E"
	position := Position{0, 0}
	for _, instruction := range instructions {
		position = getNextPositionCompassMove(position, instruction)
		if instruction.action == "F" {
			position = getNextPositionForwardMove(position, currentDir, instruction)
		}

		currentDir = getDir(currentDir, instruction)
	}

	return Abs(position.n) + Abs(position.e)
}

func getNextPositionCompassMove(position Position, instruction Instruction) Position {
	if instruction.action == "N" {
		return Position{position.n + instruction.value, position.e}
	}
	if instruction.action == "S" {
		return Position{position.n - instruction.value, position.e}
	}
	if instruction.action == "E" {
		return Position{position.n, position.e + instruction.value}
	}
	if instruction.action == "W" {
		return Position{position.n, position.e - instruction.value}
	}
	return position
}

func getNextPositionForwardMove(position Position, currentDir string, instruction Instruction) Position {
	return getNextPositionCompassMove(position, Instruction{currentDir, instruction.value})
}

func getDir(currentDir string, instruction Instruction) string {
	if instruction.action == "L" || instruction.action == "R" {
		return getDirRotate(currentDir, instruction.action, instruction.value)
	}

	return currentDir
}

func getDirRotate(currentDir string, turnDirection string, degrees int) string {
	modifier := 1
	if turnDirection == "L" {
		modifier = -1
	}
	compassPosition := getIndex(compass, currentDir)
	rotate := degrees / 90
	newDirIndex := mod(compassPosition+modifier*rotate, len(compass))

	return compass[newDirIndex]
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getIndex(s []string, value string) int {
	for i := range s {
		if s[i] == value {
			return i
		}
	}

	return -1
}

func parse(input string) []Instruction {
	var instructions []Instruction
	for _, row := range splitByNewLine(input) {
		action := string(row[0])
		value, _ := strconv.Atoi(row[1:])
		instructions = append(instructions, Instruction{
			action,
			value,
		})
	}

	return instructions
}

func splitByNewLine(input string) []string {
	return strings.Split(input, "\n")
}

func loadFile() string {
	data, err := ioutil.ReadFile("12_input.txt")
	if err != nil {
		panic(err)
	}

	return string(data)
}

func main() {
	fmt.Println("Pt1")
	fmt.Println(FindManhattanDistancePt1(parse(loadFile())))
}
