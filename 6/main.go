package main

import (
	aoc "aoc2020"
	"log"
)

func main() {
	partTwo()
}


func partOne() {
	lines := aoc.ReadInput("./6/input.txt")
	total := 0
	for _, g := range parseGroups(lines) {
		uniqueAnswers := len(g.CountUniqueAnswers())
		total += uniqueAnswers
		log.Printf("%+v had %d unique answers\n", g, uniqueAnswers)
	}

	log.Printf("there were %d total unique answers\n", total)
}

func partTwo() {
	lines := aoc.ReadInput("./6/input.txt")
	total := 0
	for _, g := range parseGroups(lines) {
		fullAnswers := g.CountFullAnswers()
		total += fullAnswers
		log.Printf("%+v had %d full answers\n", g, fullAnswers)
	}

	log.Printf("there were %d total full answers\n", total)
}

type group struct {
	personAnswers []string
}

func (g group) CountUniqueAnswers() map[string]int {
	unique := make(map[string]int)
	for _, answers := range g.personAnswers {
		for _, answer := range answers {
			cnt, ok := unique[string(answer)]
			if !ok {
				unique[string(answer)] = 1
			} else {
				unique[string(answer)] = cnt+1
			}
		}
	}

	return unique
}

func (g group) CountFullAnswers() int {
	count := 0
	for _, a := range g.CountUniqueAnswers() {
		if a == len(g.personAnswers) {
			count++
		}
	}

	return count
}

func parseGroups(lines []string) []group {
	groups := make([]group, 0)
	var current group
	for _, line := range lines {
		if line == "" {
			groups = append(groups, current)
			current = group{}
			continue
		}

		current.personAnswers = append(current.personAnswers, line)
	}
	groups = append(groups, current)

	return groups
}