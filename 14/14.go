package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Instruction struct {
	mask   string
	memory Memory
}

type Memory struct {
	address int
	value   int
}

func FindSumOfValuesLeftInMemoryPt1(instructions []Instruction) int {
	memory := make(map[int]int)
	mask := ""
	for _, instruction := range instructions {
		if instruction.mask != "" {
			mask = instruction.mask
		} else {
			memory[instruction.memory.address] = applyMask(mask, instruction.memory.value)
		}
	}

	return getSum(memory)
}

func FindSumOfValuesLeftInMemoryPt2(instructions []Instruction) int {
	memory := make(map[int]int)
	var allPossibleValues []string
	mask := ""
	for _, instruction := range instructions {
		if instruction.mask != "" {
			mask = instruction.mask
			continue
		}
		resultMask := applyMaskPt2(mask, instruction.memory.address)
		allPossibleValues = getAllPossibleValues(resultMask, 0)
		for _, value := range allPossibleValues {
			memory[binary2Int(value)] = instruction.memory.value
		}
	}

	return getSum(memory)
}

func applyMask(mask string, value int) int {
	binary := fmt.Sprintf("%036b", value)
	result := make([]byte, len(mask))
	for i := range mask {
		if string(mask[i]) == "X" {
			result[i] = binary[i]
			continue
		}
		result[i] = mask[i]
	}

	return binary2Int(string(result))
}

func binary2Int(binary string) int {
	i, _ := strconv.ParseInt(binary, 2, 64)

	return int(i)
}

func applyMaskPt2(mask string, value int) string {
	binary := fmt.Sprintf("%036b", value)
	result := make([]byte, len(mask))
	for i := range mask {
		if string(mask[i]) == "1" || string(mask[i]) == "X" {
			result[i] = mask[i]
			continue
		}
		result[i] = binary[i]
	}

	return string(result)
}

func getAllPossibleValues(mask string, pos int) []string {
	if pos == len(mask) {
		return []string{mask}
	}

	var allMasks []string

	if string(mask[pos]) == "X" {
		for _, b := range []string{"0", "1"} {
			allMasks = append(allMasks, getAllPossibleValues(replaceValueInPos(mask, b, pos), pos+1)...)
		}
	} else {
		allMasks = append(allMasks, getAllPossibleValues(mask, pos+1)...)
	}

	return allMasks
}

func replaceValueInPos(subject string, value string, pos int) string {
	result := make([]byte, len(subject))
	for i := range subject {
		if i == pos {
			result[i] = []byte(value)[0]
		} else {
			result[i] = subject[i]
		}
	}

	return string(result)
}

func getSum(values map[int]int) int {
	x := 0
	for _, i := range values {
		x = x + i
	}

	return x
}

func parse(input string) []Instruction {
	var instructions []Instruction
	for _, row := range splitByNewLine(input) {
		var instruction Instruction
		if isRowMask(row) {
			instruction = Instruction{
				splitByEquals(row)[1],
				Memory{},
			}
		} else {
			memoryRaw := splitByEquals(row)
			re := regexp.MustCompile("[0-9]+")
			memoryAddress, _ := strconv.Atoi(re.FindAllString(memoryRaw[0], 1)[0])
			memoryValue, _ := strconv.Atoi(memoryRaw[1])

			instruction = Instruction{
				"",
				Memory{
					memoryAddress,
					memoryValue,
				},
			}
		}

		instructions = append(instructions, instruction)
	}

	return instructions
}

func isRowMask(row string) bool {
	return string(row[1]) == "a"
}

func splitByNewLine(input string) []string {
	return strings.Split(input, "\n")
}

func splitByEquals(input string) []string {
	return strings.Split(input, " = ")
}

func loadFile() string {
	data, err := ioutil.ReadFile("14_input.txt")
	if err != nil {
		panic(err)
	}

	return string(data)
}

func main() {
	fmt.Println("Pt1")
	fmt.Println(FindSumOfValuesLeftInMemoryPt1(parse(loadFile())))
	fmt.Println("Pt2")
	fmt.Println(FindSumOfValuesLeftInMemoryPt2(parse(loadFile())))
}
