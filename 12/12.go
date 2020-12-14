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

func FindManhattanDistancePt2(instructions []Instruction) int {
	wayPoint := Position{1, 10}
	position := Position{0, 0}
	for _, instruction := range instructions {
		if instruction.action == "F" {
			position = getNextPositionForwardMovePt2(position, wayPoint, instruction)
			continue
		}
		if instruction.action == "L" || instruction.action == "R" {
			wayPoint = getNextWaypointAfterRotation(wayPoint, instruction)
			continue
		}
		wayPoint = getNextPositionCompassMove(wayPoint, instruction)
	}

	return Abs(position.n) + Abs(position.e)
}

func getNextPositionForwardMovePt2(position Position, waypoint Position, instruction Instruction) Position {
	if instruction.action != "F" {
		return position
	}

	return sumPosition(position, multiplyPosition(waypoint, instruction.value))
}

func multiplyPosition(position Position, x int) Position {
	return Position{position.n * x, position.e * x}
}

func getNextWaypointAfterRotation(waypoint Position, instruction Instruction) Position {
	if instruction.action != "L" && instruction.action != "R" {
		return waypoint
	}
	if instruction.action == "L" {
		instruction = convertToRInstruction(instruction)
	}

	return getWaypointAfterRRotation(waypoint, instruction)
}

func getWaypointAfterRRotation(waypoint Position, instruction Instruction) Position {
	if instruction.action != "R" {
		panic("Fed a bad instruction")
	}

	switch instruction.value {
	case 90:
		return Position{-1 * waypoint.e, waypoint.n}
	case 180:
		return Position{-1 * waypoint.n, -1 * waypoint.e}
	case 270:
		return Position{waypoint.e, -1 * waypoint.n}
	default:
		return waypoint
	}
}

func convertToRInstruction(instruction Instruction) Instruction {
	if instruction.action != "L" && instruction.action != "R" {
		return instruction
	}
	cwRotations := []int{0, 90, 180, 270}
	modifier := 1
	if instruction.action == "L" {
		modifier = -1
	}
	rotate := instruction.value / 90

	return Instruction{"R", cwRotations[mod(modifier*rotate, len(compass))]}
}

func sumPosition(positionA Position, positionB Position) Position {
	return Position{positionA.n + positionB.n, positionA.e + positionB.e}
}

func getNextPositionCompassMove(position Position, instruction Instruction) Position {
	switch instruction.action {
	case "N":
		return Position{position.n + instruction.value, position.e}
	case "E":
		return Position{position.n, position.e + instruction.value}
	case "S":
		return Position{position.n - instruction.value, position.e}
	case "W":
		return Position{position.n, position.e - instruction.value}
	default:
		return position
	}
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
	fmt.Println("Pt2")
	fmt.Println(FindManhattanDistancePt2(parse(loadFile())))
}
