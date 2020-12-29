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
		sum = sum + evaluateExpressionPt1(expression)
	}

	return sum
}

func FindSumOfExpressionsPt2(expressions []string) int {
	sum := 0
	for _, expression := range expressions {
		sum = sum + evaluateExpressionPt2(expression)
	}

	return sum
}

func evaluateExpressionPt1(expression string) int {
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
				evaluateExpressionPt1(expression[i+1:]),
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

func evaluateExpressionPt2(expression string) int {
	return evaluateExpressionPt1(transformExpressionPt2(expression))
}

func transformExpressionPt2(expression string) string {
	for i := 0; i < len(expression); i++ {
		v := string(expression[i])
		if isSpace(v) {
			continue
		}

		if isAddition(v) {
			expression = surroundWithParenthesis(expression, i)
			i++
		}
	}

	return expression
}

func surroundWithParenthesis(expression string, pos int) string {
	// add left
	for i := pos - 1; i >= 0; i-- {
		v := string(expression[i])
		if isSpace(v) {
			continue
		}

		if isDigit(v) {
			expression = insertIntoString(expression, "(", i)
			pos++
			break
		}

		if isRightParenthesis(v) {
			insertPosition := getMatchingLeftParenthesisPosition(expression[:i])
			expression = insertIntoString(
				expression,
				"(",
				insertPosition,
			)
			pos++
			break
		}
	}

	// add right
	for i := pos + 1; i < len(expression); i++ {
		v := string(expression[i])
		if isSpace(v) {
			continue
		}

		if isDigit(v) {
			expression = insertIntoString(expression, ")", i+1)
			break
		}

		if isLeftParenthesis(v) {
			insertPosition := i + getMatchingRightParenthesisPosition(expression[i+1:]) + 1
			expression = insertIntoString(
				expression,
				")",
				insertPosition,
			)
			pos++
			break
		}
	}

	return expression
}

func insertIntoString(subject string, value string, pos int) string {
	if pos < 0 {
		pos = 0
	}
	if pos > len(subject) {
		pos = len(subject)
	}

	return subject[:pos] + value + subject[pos:]
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

// e.g., (2 * 3 => 0
func getMatchingLeftParenthesisPosition(expression string) int {
	parenthesisStack := 1
	for i := len(expression) - 1; i >= 0; i-- {
		v := string(expression[i])
		if isRightParenthesis(v) {
			parenthesisStack++
		}
		if isLeftParenthesis(v) {
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
	fmt.Println("Pt2")
	fmt.Println(FindSumOfExpressionsPt2(parse(loadFile())))
}
