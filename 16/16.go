package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Range struct {
	l int
	u int
}

type Rule struct {
	name   string
	rangeA Range
	rangeB Range
}

type Notes struct {
	rules         []Rule
	yourTicket    []int
	nearbyTickets [][]int
}

func FindInvalidTicketScanningErrorRatePt1(notes Notes) int {
	validEntries := getValidEntries(notes.rules)
	var invalidTicketEntries []int
	for _, nearbyTicket := range notes.nearbyTickets {
		for _, entry := range nearbyTicket {
			if validEntries[entry] != true {
				invalidTicketEntries = append(invalidTicketEntries, entry)
			}
		}
	}

	return getSum(invalidTicketEntries)
}

func getSum(values []int) int {
	x := 0
	for _, i := range values {
		x = x + i
	}

	return x
}

func getValidEntries(rules []Rule) map[int]bool {
	validEntries := make(map[int]bool)
	for _, rule := range rules {
		for i := rule.rangeA.l; i <= rule.rangeA.u; i++ {
			validEntries[i] = true
		}
		for i := rule.rangeB.l; i <= rule.rangeB.u; i++ {
			validEntries[i] = true
		}
	}

	return validEntries
}

func parse(input string) Notes {
	sections := strings.Split(input, "\n\n")
	rules := parseRules(sections[0])
	yourTicket := parseTicket(strings.Split(sections[1], "\n")[1])
	nearbyTickets := parseNearbyTickets(sections[2])

	return Notes{
		rules,
		yourTicket,
		nearbyTickets,
	}
}

func parseTicket(input string) []int {
	var ticket []int
	for _, value := range strings.Split(input, ",") {
		number, _ := strconv.Atoi(value)
		ticket = append(ticket, number)
	}

	return ticket
}

func parseNearbyTickets(nearbyTicketsSection string) [][]int {
	var nearbyTickets [][]int
	for _, rawTicketRow := range strings.Split(nearbyTicketsSection, "\n") {
		if rawTicketRow == "nearby tickets:" {
			continue
		}
		nearbyTickets = append(nearbyTickets, parseTicket(rawTicketRow))
	}

	return nearbyTickets
}

func parseRules(input string) []Rule {
	var rules []Rule
	for _, row := range strings.Split(input, "\n") {
		rawRules := strings.Split(row, ": ")

		rawRanges := strings.Split(rawRules[1], " or ")

		rule := Rule{
			rawRules[0],
			parseRange(rawRanges[0]),
			parseRange(rawRanges[1]),
		}

		rules = append(rules, rule)
	}

	return rules
}

func parseRange(input string) Range {
	numbers := strings.Split(input, "-")
	l, _ := strconv.Atoi(numbers[0])
	u, _ := strconv.Atoi(numbers[1])

	return Range{l, u}
}

func loadFile() string {
	data, err := ioutil.ReadFile("16_input.txt")
	if err != nil {
		panic(err)
	}

	return string(data)
}

func main() {
	fmt.Println("Pt1")
	fmt.Println(FindInvalidTicketScanningErrorRatePt1(parse(loadFile())))
}
