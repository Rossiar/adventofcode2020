package main

import (
	aoc "aoc2020"
	"log"
	"regexp"
	"strings"
)

var instructionPattern = regexp.MustCompile(`^(nop|acc|jmp) (\+|-)(\d+)$`)

func main() {
	partTwo()
}

func partOne() {
	lines := aoc.ReadInput("./8/input.txt")
	_, acc, _ := didDuplicate(lines)
	log.Printf("accumulator is %d", acc)
}

func partTwo() {
	lines := aoc.ReadInput("./8/input.txt")
	for i := 0; i < len(lines); i++ {
		lines[i] = swapPrefix(lines[i])

		stack, acc, duplicate := didDuplicate(lines)
		if !duplicate {
			for _, s := range stack {
				log.Printf("%s", s)
			}
			log.Printf("final acc is %d", acc)
			break
		}

		// if not then reset the instruction and keep trying
		lines[i] = swapPrefix(lines[i])
	}
}

func didDuplicate(lines []string) ([]string, int, bool) {
	acc := 0
	used := make(map[int]int, 0)
	stack := make([]string, 0)
	step := 1
	for i := 0; i < len(lines); i += step {
		line := lines[i]
		instruction := instructionPattern.FindStringSubmatch(line)
		if _, ok := used[i]; ok {
			return stack, acc, true
		}
		used[i] = 1
		stack = append(stack, line)

		val := aoc.Atoi(instruction[3])
		if instruction[2] == "-" {
			val *= -1
		}

		switch instruction[1] {
		case "acc":
			acc += val
			step = 1
		case "jmp":
			step = val
		case "nop":
			step = 1
		}
	}

	return stack, acc, false
}

func swapPrefix(line string) string {
	if strings.HasPrefix(line, "jmp") {
		return strings.ReplaceAll(line, "jmp", "nop")
	}

	if strings.HasPrefix(line, "nop") {
		return strings.ReplaceAll(line, "nop", "jmp")
	}

	return line
}
