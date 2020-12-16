package main

import (
	aoc "aoc2020"
	"log"
	"regexp"
	"strings"
)

var (
	fieldPattern  = regexp.MustCompile(`^.*?: (\d+-\d+) or (\d+-\d+)$`)
	ticketPattern = regexp.MustCompile(`(\d(,|))+`)
)

func main() {
	partOne()
}

type puzzle struct {
	fields  []field
	tickets [][]int
}

func (p puzzle) TicketScanningErrorRate() int {
	invalidTicketFields := make([]int, 0)
	for _, ticket := range p.tickets {
		for _, val := range ticket {
			valid := false
			for _, f := range p.fields {
				if f.ValidForAnyRule(val) {
					valid = true
					break
				}
			}
			if !valid {
				invalidTicketFields = append(invalidTicketFields, val)
			}
		}
	}

	ticketScanningErrorRate := 0
	for _, val := range invalidTicketFields {
		ticketScanningErrorRate += val
	}
	return ticketScanningErrorRate
}

type field struct {
	Rules []rule
}

func parseField(line string) field {
	matches := fieldPattern.FindStringSubmatch(line)
	return field{
		Rules: []rule{
			parseRule(matches[1]),
			parseRule(matches[2]),
		},
	}
}

func (f field) ValidForAnyRule(num int) bool {
	for _, rule := range f.Rules {
		if rule.Validate(num) {
			return true
		}
	}
	return false
}

type rule struct {
	Max int
	Min int
}

func (r rule) Validate(num int) bool {
	return num >= r.Min && num <= r.Max
}

func parseRule(raw string) rule {
	components := strings.Split(raw, "-")
	return rule{Min: aoc.Atoi(components[0]), Max: aoc.Atoi(components[1])}
}

func parseTicket(raw string) []int {
	components := strings.Split(raw, ",")
	numbers := make([]int, len(components))
	for i, r := range components {
		numbers[i] = aoc.Atoi(r)
	}
	return numbers
}

func parsePuzzle(file string) puzzle {
	lines := aoc.ReadInput(file)
	fields := make([]field, 0)
	tickets := make([][]int, 0)
	for _, line := range lines {
		if line == "" {
			continue
		}

		if fieldPattern.MatchString(line) {
			fields = append(fields, parseField(line))
			continue
		}

		if ticketPattern.MatchString(line) {
			tickets = append(tickets, parseTicket(line))
			continue
		}
	}

	return puzzle{fields: fields, tickets: tickets}
}

func partOne() {
	puzzle := parsePuzzle("./16/input.txt")
	log.Printf("ticket scanning error rate %d", puzzle.TicketScanningErrorRate())
}

func partTwo() {

}
