package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func FindSumOfGroupCountsPt1(allAnswers []string) int {
	total := 0
	for _, answers := range allAnswers {
		total = total + getGroupCountPt1(answers)
	}

	return total
}

func getGroupCountPt1(answers string) int {
	m := make(map[string]int)
	for _, c := range answers {
		matched, _ := regexp.MatchString(`\w`, string(c))
		if matched {
			m[string(c)] = 1
		}
	}

	return len(m)
}

func parse(input string) []string {
	return strings.Split(input, "\n\n")
}

func loadFile() string {
	data, err := ioutil.ReadFile("06_input.txt")
	if err != nil {
		panic(err)
	}

	return string(data)
}

func main() {
	fmt.Println("Pt1 Sum of Group Counts")
	fmt.Println(FindSumOfGroupCountsPt1(parse(loadFile())))
}
