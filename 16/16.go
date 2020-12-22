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
	name         string
	validEntries map[int]bool
}

type Notes struct {
	rules         []Rule
	yourTicket    []int
	nearbyTickets [][]int
}

func FindInvalidTicketScanningErrorRatePt1(notes Notes) int {
	validEntries := getAllValidEntries(notes.rules)
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

func FindMultipleOfDepartedPt2(notes Notes) int {
	mapped := getRuleColumnMapping(notes)

	var yourTicketDepartedValues []int
	for ruleI, colI := range mapped {
		if strings.HasPrefix(notes.rules[ruleI].name, "departure") {
			yourTicketDepartedValues = append(yourTicketDepartedValues, notes.yourTicket[colI])
		}
	}

	return getMultiple(yourTicketDepartedValues)
}

func getRuleColumnMapping(notes Notes) map[int]int {
	bpGraph := getBPGraph(notes)
	return findMatchingSolution(bpGraph)
}

// credit to https://www.geeksforgeeks.org/maximum-bipartite-matching/
func findMatchingSolution(bpGraph [][]bool) map[int]int {
	N := len(bpGraph)
	M := len(bpGraph[0])
	matchR := make(map[int]int, N)
	for i := 0; i < N; i++ {
		matchR[i] = -1
	}

	totalMatched := 0
	for u := 0; u < M; u++ {
		seen := make(map[int]bool, N)
		for i := 0; i < N; i++ {
			seen[i] = false
		}

		if bpm(bpGraph, u, seen, matchR) {
			totalMatched++
		}

	}
	if totalMatched != N {
		panic("Solution not found")
	}

	return matchR
}

func bpm(bpGraph [][]bool, u int, seen map[int]bool, matchR map[int]int) bool {
	for v := 0; v < len(bpGraph); v++ {
		if bpGraph[u][v] && seen[v] == false {
			seen[v] = true
			if matchR[v] < 0 || bpm(bpGraph, matchR[v], seen, matchR) {
				matchR[v] = u
				return true
			}

		}
	}
	return false
}

func getBPGraph(notes Notes) [][]bool {
	allTickets := append(getValidNearbyTickets(notes), notes.yourTicket)

	var bpGraph [][]bool
	for columnI := range allTickets[0] {
		var rulesSatisfy []bool
		for ruleI := range notes.rules {
			value := false
			if doesRuleSatisfyCol(notes.rules[ruleI], allTickets, columnI) {
				value = true
			}
			rulesSatisfy = append(rulesSatisfy, value)
		}
		bpGraph = append(bpGraph, rulesSatisfy)
	}

	return bpGraph
}

func getMultiple(values []int) int {
	x := 1
	for _, i := range values {
		x = x * i
	}

	return x
}

func getValidNearbyTickets(notes Notes) [][]int {
	validEntries := getAllValidEntries(notes.rules)
	var nearbyTickets [][]int
	for _, nearbyTicket := range notes.nearbyTickets {
		if isTicketValid(nearbyTicket, validEntries) {
			nearbyTickets = append(nearbyTickets, nearbyTicket)
		}
	}

	return nearbyTickets
}

func isTicketValid(nearbyTicket []int, validEntries map[int]bool) bool {
	for _, entry := range nearbyTicket {
		if validEntries[entry] != true {
			return false
		}
	}

	return true
}

func doesRuleSatisfyCol(rule Rule, nearbyTickets [][]int, x int) bool {
	for y := range nearbyTickets {
		if rule.validEntries[nearbyTickets[y][x]] != true {
			//fmt.Printf("Rule: %v, fail for %v at %v, number not found: %v \n", rule.name, x, y, nearbyTickets[y][x])
			return false
		}
	}

	return true
}

func getAllValidEntries(rules []Rule) map[int]bool {
	validEntries := make(map[int]bool)
	for _, rule := range rules {
		for validEntry := range rule.validEntries {
			validEntries[validEntry] = true
		}
	}

	return validEntries
}

func getSum(values []int) int {
	x := 0
	for _, i := range values {
		x = x + i
	}

	return x
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
			getRuleValidEntries(parseRange(rawRanges[0]), parseRange(rawRanges[1])),
		}

		rules = append(rules, rule)
	}

	return rules
}

func getRuleValidEntries(numberRangeA Range, numberRangeB Range) map[int]bool {
	validEntries := make(map[int]bool)
	for i := numberRangeA.l; i <= numberRangeA.u; i++ {
		validEntries[i] = true
	}
	for i := numberRangeB.l; i <= numberRangeB.u; i++ {
		validEntries[i] = true
	}

	return validEntries
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
	fmt.Println("Pt2")
	fmt.Println(FindMultipleOfDepartedPt2(parse(loadFile())))
}
