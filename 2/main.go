package main

import (
	aoc "aoc2020"
	"log"
	"regexp"
	"strings"
)

var rPolicy = regexp.MustCompile(`(\d+)-(\d+) (\w)`)

type policy struct {
	min int
	max int
	char string
}

func (p policy) ValidPosition(password string) bool {
	first := password[p.min-1:p.min]
	second := password[p.max-1:p.max]
	return (first == p.char) != (second == p.char)
}

func (p policy) ValidCount(password string) bool {
	cnt := strings.Count(password, p.char)
	return cnt >= p.min && cnt <= p.max
}

func main() {
	validByPosition()
}

func validByCount() {
	lines := aoc.ReadInput("./2/input.txt")
	countValid := 0
	for _, line := range lines {
		subs := strings.Split(line, ":")
		policy := parsePolicy(subs[0])
		password := strings.TrimSpace(subs[1])
		if policy.ValidCount(password) {
			countValid++
		}
	}

	log.Printf("found %d valid passwords\n", countValid)
}

func validByPosition() {
	lines := aoc.ReadInput("./2/input.txt")
	countValid := 0
	for _, line := range lines {
		subs := strings.Split(line, ":")
		policy := parsePolicy(subs[0])
		password := strings.TrimSpace(subs[1])
		if policy.ValidPosition(password) {
			countValid++
		}
	}

	log.Printf("found %d valid passwords\n", countValid)
}

func parsePolicy(line string) policy {
	matches := rPolicy.FindStringSubmatch(line)
	return policy{
		min:  aoc.Atoi(matches[1]),
		max:  aoc.Atoi(matches[2]),
		char: matches[3],
	}
}