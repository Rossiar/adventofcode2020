package main

import (
	aoc "aoc2020"
	"log"
	"sort"
)

func main() {
	partTwo()
}

func partOne() {
	lines := aoc.ReadInput("./9/input.txt")
	result := findFirstInvalid(25, aoc.ToIntSlice(lines))
	log.Printf("%d was the first number", result)
}

func partTwo() {
	lines := aoc.ToIntSlice(aoc.ReadInput("./9/input.txt"))
	oddNumber := findFirstInvalid(25, lines)
	log.Printf("%d was the odd number", oddNumber)
	result := findContiguousComponents(oddNumber, lines)
	log.Printf("sequence: %+v", result)
	sort.Slice(result, func(i, j int) bool { return result[i] < result[j] })
	log.Printf("sorted: %+v", result)
	start := result[0]
	end := result[len(result)-1]
	log.Printf("%d + %d = %d", start, end, start+end)
}

func findFirstInvalid(preamble int, lines []int) int {
	for i := preamble; i < len(lines); i++ {
		goal := lines[i]
		slide := lines[i-preamble : i]

		// put elements into hashmap
		lookup := map[int]bool{}
		for _, val := range slide {
			lookup[val] = true
		}

		// check if the components of goal exist in the lookup
		if !hasComponents(goal, lookup) {
			return goal
		}
	}

	return -1
}

func findContiguousComponents(goal int, lines []int) []int {
	components := make([]int, 0)
	for i := 0; i < len(lines); i++ {
		components = []int{lines[i]}
		total := lines[i]
		for j := i + 1; j < len(lines); j++ {
			if total+lines[j] > goal {
				continue
			}

			total += lines[j]
			components = append(components, lines[j])

			if total == goal {
				return components
			}
		}
	}

	return components
}

func hasComponents(goal int, slide map[int]bool) bool {
	for a := range slide {
		if slide[goal-a] {
			return true
		}
	}

	return false
}
