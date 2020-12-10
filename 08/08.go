package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Instruction struct {
	kind  string
	value int
}

func FindAccValuePt1(instructions []Instruction) int {
	accumulator := 0
	visited := make(map[int]bool)
	i := 0
	for visited[i] != true {
		visited[i] = true
		if instructions[i].kind == "nop" {
			i++
			continue
		}
		if instructions[i].kind == "jmp" {
			i = i + instructions[i].value
			continue
		}
		if instructions[i].kind == "acc" {
			accumulator = accumulator + instructions[i].value
			i++
			continue
		}
	}

	return accumulator
}

func parse(input string) []Instruction {
	var instructions []Instruction
	for _, row := range splitByNewLine(input) {
		instructions = append(instructions, createFromRow(row))
	}

	return instructions
}

func splitByNewLine(input string) []string {
	return strings.Split(input, "\n")
}

func createFromRow(row string) Instruction {
	split := strings.Split(row, " ")
	value, _ := strconv.Atoi(split[1])
	return Instruction{
		split[0],
		value,
	}
}

func loadFile() string {
	data, err := ioutil.ReadFile("08_input.txt")
	if err != nil {
		panic(err)
	}

	return string(data)
}

func main() {
	fmt.Println("Pt1")
	fmt.Println(FindAccValuePt1(parse(loadFile())))
}
