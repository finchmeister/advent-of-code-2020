package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

type Input struct {
	rules    map[string]string
	messages []string
}

func FindMessageCountMatchRule0Pt1(input Input) int {
	return getCountOfMatchingMessages(input.rules, input.messages)
}

func FindMessageCountMatchRule0Pt2(input Input) int {
	var matchingMessages []string
	// Solution obtained with 10, these values compute same answer
	maxRecursionRule8 := 5
	maxRecursionRule11 := 4
	for i := 1; i < maxRecursionRule8; i++ {
		input.rules["8"] = getRule8AfterNRecursion(i)
		for j := 1; j < maxRecursionRule11; j++ {
			input.rules["11"] = getRule11AfterNRecursion(j)
			matchingMessages = append(
				matchingMessages,
				getMatchingMessages(input.rules, input.messages)...,
			)
		}
	}

	return getUniqueMessageCount(matchingMessages)
}

func getUniqueMessageCount(messages []string) int {
	keys := make(map[string]bool)
	uniqueMessages := 0
	for _, entry := range messages {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			uniqueMessages++
		}
	}

	return uniqueMessages
}

func getRule8AfterNRecursion(n int) string {
	rule := "42 | 42"
	for i := 0; i < n; i++ {
		rule = rule + " 42"
	}

	return rule
}

func getRule11AfterNRecursion(n int) string {
	return "42 31 |" + strings.Repeat(" 42", n+1) + strings.Repeat(" 31", n+1)
}

func getCountOfMatchingMessages(rules map[string]string, messages []string) int {
	return len(getMatchingMessages(rules, messages))
}

func getMatchingMessages(rules map[string]string, messages []string) []string {
	regex := "^" + convertRuleToRegex(rules, "0") + "$"
	r := regexp.MustCompile(regex)
	var matchingMessages []string
	for _, message := range messages {
		matches := r.FindStringSubmatch(message)
		if len(matches) != 0 {
			matchingMessages = append(matchingMessages, message)
		}
	}

	return matchingMessages
}

func convertRuleToRegex(rules map[string]string, ruleNo string) string {
	currentRule := strings.Split(rules[ruleNo], " ")
	expandedRule := make([]string, len(currentRule))
	for i, value := range currentRule {
		ruleReplacement := rules[value]
		if ruleReplacement == "a" || ruleReplacement == "b" {
			expandedRule[i] = ruleReplacement
			continue
		}
		if value == "|" {
			expandedRule[i] = value
			continue
		}

		expandedRule[i] = "(" + convertRuleToRegex(rules, value) + ")"
	}

	return strings.Join(expandedRule, "")
}

func parse(input string) Input {
	sections := strings.Split(input, "\n\n")
	rules := parseRules(sections[0])

	return Input{
		rules,
		strings.Split(sections[1], "\n"),
	}
}

func parseRules(input string) map[string]string {
	rules := make(map[string]string, len(input))
	for _, row := range strings.Split(input, "\n") {
		rawRules := strings.Split(row, ": ")

		rules[rawRules[0]] = strings.Replace(rawRules[1], `"`, "", -1)
	}

	return rules
}

func loadFile() string {
	data, err := ioutil.ReadFile("19_input.txt")
	if err != nil {
		panic(err)
	}

	return string(data)
}

func main() {
	fmt.Println("Pt1")
	fmt.Println(FindMessageCountMatchRule0Pt1(parse(loadFile())))
	fmt.Println("Pt2")
	fmt.Println(FindMessageCountMatchRule0Pt2(parse(loadFile())))
}
