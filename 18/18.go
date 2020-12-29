package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func FindSumOfExpressionsPt1(expressions []string) int {
	sum := 0
	for _, expression := range expressions {
		sum = sum + evaluateExpression(expression)
	}

	return sum
}

func evaluateExpression(expression string) int {
	operator := "+"
	result := 0

	for i := 0; i < len(expression); i++ {
		v := string(expression[i])
		if isSpace(v) {
			continue
		}

		if isLeftParenthesis(v) {
			result = computeValue(
				result,
				evaluateExpression(expression[i+1:]),
				operator,
			)

			i = i + getMatchingRightParenthesisPosition(expression[i+1:]) + 1
			continue
		}

		if isRightParenthesis(v) {
			return result
		}

		if isAddition(v) || isMultiplication(v) {
			operator = v
		}

		if isDigit(v) {
			result = computeValue(result, getDigit(v), operator)
		}
	}

	return result
}

func computeValue(a int, b int, operator string) int {
	if operator == "+" {
		return a + b
	}

	if operator == "*" {
		return a * b
	}

	panic("Unexpected Operator")
}

func isDigit(v string) bool {
	if _, err := strconv.Atoi(v); err != nil {
		return false
	}

	return true
}

func getDigit(v string) int {
	digit, _ := strconv.Atoi(v)

	return digit
}

func isSpace(v string) bool {
	return v == " "
}

func isLeftParenthesis(v string) bool {
	return v == "("
}

func isRightParenthesis(v string) bool {
	return v == ")"
}

func isAddition(v string) bool {
	return v == "+"
}

func isMultiplication(v string) bool {
	return v == "*"
}

// e.g., 2 * 3) => 5
func getMatchingRightParenthesisPosition(expression string) int {
	parenthesisStack := 1
	for i := range expression {
		v := string(expression[i])
		if isLeftParenthesis(v) {
			parenthesisStack++
		}
		if isRightParenthesis(v) {
			if parenthesisStack == 1 {
				return i
			}
			parenthesisStack--
		}
	}
	panic("Mismatched parenthesis")
}

func parse(input string) []string {
	return strings.Split(input, "\n")
}

func loadFile() string {
	data, err := ioutil.ReadFile("18_input.txt")
	if err != nil {
		panic(err)
	}

	return string(data)
}

func main() {
	fmt.Println("Pt1")
	fmt.Println(FindSumOfExpressionsPt1(parse(loadFile())))
}
