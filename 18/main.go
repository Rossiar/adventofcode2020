package main

import (
	aoc "aoc2020"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	partOne()

	partTwo()
}

func partOne() {
	lines := aoc.ReadInput("./18/input.txt")
	total := 0
	for _, line := range lines {
		total += parse(line)
	}
	log.Printf("%d", total)
}

var simpleExpression = regexp.MustCompile(`\([^(;]+?\)`)

func parse(expression string) int {
	result := -1
	operation := ""

	for strings.Contains(expression, "(") {
		expression = simpleExpression.ReplaceAllStringFunc(expression, func(s string) string {
			trimParenthesis := s[1 : len(s)-1]
			return strconv.Itoa(parse(trimParenthesis))
		})
	}

	for _, char := range strings.Split(expression, " ") {
		if char == "+" {
			operation = "+"
		} else if char == "*" {
			operation = "*"
		} else {
			val := aoc.Atoi(char)
			if result == -1 {
				result = val
				continue
			}

			switch operation {
			case "+":
				result += val
			case "*":
				result *= val
			}
		}
	}
	return result
}

func partTwo() {
	lines := aoc.ReadInput("./18/input.txt")
	total := 0
	for _, line := range lines {
		total += parseAdvanced(line)
	}
	log.Printf("%d", total)
}

var (
	additionExpression       = regexp.MustCompile(`(\d+) \+ (\d+)`)
	multiplicationExpression = regexp.MustCompile(`(\d+) \* (\d+)`)
)

func parseAdvanced(expression string) int {
	for strings.Contains(expression, "(") {
		expression = simpleExpression.ReplaceAllStringFunc(expression, func(s string) string {
			trimParenthesis := s[1 : len(s)-1]
			return strconv.Itoa(parseAdvanced(trimParenthesis))
		})
	}

	for strings.Contains(expression, "+") {
		expression = additionExpression.ReplaceAllStringFunc(expression, func(s string) string {
			match := additionExpression.FindStringSubmatch(s)
			sum := aoc.Atoi(match[1]) + aoc.Atoi(match[2])
			return strconv.Itoa(sum)
		})
	}

	for strings.Contains(expression, "*") {
		expression = multiplicationExpression.ReplaceAllStringFunc(expression, func(s string) string {
			match := multiplicationExpression.FindStringSubmatch(s)
			product := aoc.Atoi(match[1]) * aoc.Atoi(match[2])
			return strconv.Itoa(product)
		})
	}

	return aoc.Atoi(expression)
}
