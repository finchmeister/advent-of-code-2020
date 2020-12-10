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
	accumulator, _ := getAccumulator(instructions)
	return accumulator
}

func FindAccValuePt2(instructions []Instruction) int {
	modifiedInstructions := make([]Instruction, len(instructions))
	for i, instruction := range instructions {
		if instruction.kind == "jmp" || instruction.kind == "nop" {
			copy(modifiedInstructions, instructions)
			modifiedInstructions[i] = getReplacementInstruction(instruction)
			accumulator, success := getAccumulator(modifiedInstructions)
			if success {
				return accumulator
			}
		}
	}
	panic("Solution not found")
}

func getReplacementInstruction(instruction Instruction) Instruction {
	var newInstructionKind string
	if instruction.kind == "jmp" {
		newInstructionKind = "nop"
	} else {
		newInstructionKind = "jmp"
	}

	return Instruction{
		newInstructionKind,
		instruction.value,
	}
}

// Returns accumulator and whether finished successfully
func getAccumulator(instructions []Instruction) (int, bool) {
	accumulator := 0
	visited := make(map[int]bool)
	i := 0
	for visited[i] != true && i != len(instructions) {
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
	if i == len(instructions) {
		return accumulator, true
	}

	return accumulator, false
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
	fmt.Println("Pt2")
	fmt.Println(FindAccValuePt2(parse(loadFile())))
}
