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

	regex := "^" + convertRuleToRegex(input.rules, "0") + "$"
	r := regexp.MustCompile(regex)

	matchingMessages := 0
	for _, message := range input.messages {
		matches := r.FindStringSubmatch(message)
		if len(matches) != 0 {
			matchingMessages++
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
}
