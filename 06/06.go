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
		total = total + getCountOfUniqueAnswers(answers)
	}

	return total
}

func FindSumOfGroupCountsPt2(allAnswers []string) int {
	total := 0
	for _, answers := range allAnswers {
		total = total + getCountOfCustomersAllWithSameAnswer(answers)
	}

	return total
}

func getCountOfCustomersAllWithSameAnswer(customerAnswers string) int {
	uniqueAnswers := getUniqueAnswers(customerAnswers)
	count := 0
	for answer := range uniqueAnswers {
		if doAllCustomersHaveAnswer(answer, customerAnswers) {
			count++
		}
	}

	return count
}

func doAllCustomersHaveAnswer(answer string, customersAnswers string) bool {
	for _, customer := range strings.Split(customersAnswers, "\n") {
		if strings.Contains(customer, answer) == false {
			return false
		}
	}

	return true
}

func getCountOfUniqueAnswers(answers string) int {
	return len(getUniqueAnswers(answers))
}

func getUniqueAnswers(answers string) map[string]int {
	m := make(map[string]int)
	for _, c := range answers {
		matched, _ := regexp.MatchString(`\w`, string(c))
		if matched {
			m[string(c)] = 1
		}
	}

	return m
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
	fmt.Println("Pt2 Sum of Group Counts")
	fmt.Println(FindSumOfGroupCountsPt2(parse(loadFile())))
}
