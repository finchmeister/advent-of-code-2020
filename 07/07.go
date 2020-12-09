package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func FindOuterBagCountPt1(allBags map[string]map[string]int) int {
	return len(findParentBagCount("shiny gold", allBags, make(map[string]bool)))
}

func findParentBagCount(bagColour string, allBags map[string]map[string]int, bagsContaining map[string]bool) map[string]bool {
	for parentBagColour, parentBagContents := range allBags {
		for childBagColour, _ := range parentBagContents {
			if childBagColour == bagColour {
				bagsContaining[parentBagColour] = true
				findParentBagCount(parentBagColour, allBags, bagsContaining)
			}
		}
	}

	return bagsContaining
}

func parse(input string) map[string]map[string]int {
	return parseRows(parseInput(input))
}

func parseRows(rows []string) map[string]map[string]int {
	allBags := make(map[string]map[string]int)
	re := regexp.MustCompile(`^(\w+ \w+)`)
	for _, row := range rows {
		match := re.FindString(row)
		allBags[match] = parseBagContents(row)
	}

	return allBags
}

func parseBagContents(row string) map[string]int {
	contents := make(map[string]int)
	re := regexp.MustCompile(`(\d+) (\w+ \w+)`)
	matches := re.FindAllString(row, -1)
	for _, match := range matches {
		colour := strings.SplitN(match, " ", 2)[1]
		count, _ := strconv.Atoi(strings.SplitN(match, " ", 2)[0])
		contents[colour] = count
	}

	return contents
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

func loadFile() string {
	data, err := ioutil.ReadFile("07_input.txt")
	if err != nil {
		panic(err)
	}

	return string(data)
}

func main() {
	fmt.Println("Pt1")
	fmt.Println(FindOuterBagCountPt1(parseRows(parseInput(loadFile()))))
}
