package main

import (
	aoc "aoc2020"
	"log"
	"regexp"
	"sort"
	"strings"
)

var (
	fieldPattern  = regexp.MustCompile(`^(.*?): (\d+-\d+) or (\d+-\d+)$`)
	ticketPattern = regexp.MustCompile(`(\d(,|))+`)
)

func main() {
	partTwo()
}

type puzzle struct {
	fields  []field
	tickets [][]int
}

type fieldMapping struct {
	Name          string
	Possibilities []int
}

func (p *puzzle) FindFieldMapping() map[int]string {
	p.RemoveInvalidTickets()
	mapping := make([]fieldMapping, len(p.fields))
	for i, f := range p.fields {
		fm := fieldMapping{Name: f.Name}
		for i := 0; i < len(p.fields); i++ {
			allValid := true
			for _, ticket := range p.tickets {
				if !f.ValidForAnyRule(ticket[i]) {
					allValid = false
					break
				}
			}
			if allValid {
				fm.Possibilities = append(fm.Possibilities, i)
			}
		}
		mapping[i] = fm
	}

	// basically go from lowest possibilities to highest - so whiny fields like arrival location
	// that only have 1 possibility can be mapped properly - this break horribly if you have 2 fields that
	// have the same possibilities :)
	sort.Slice(mapping, func(i, j int) bool {
		return len(mapping[i].Possibilities) < len(mapping[j].Possibilities)
	})

	finalMapping := make(map[int]string)
	for _, fm := range mapping {
		for _, poss := range fm.Possibilities {
			if _, exists := finalMapping[poss]; !exists {
				finalMapping[poss] = fm.Name
				break
			}
		}
	}

	return finalMapping
}

func (p *puzzle) RemoveInvalidTickets() {
	validTickets := make([][]int, 0)
	for _, ticket := range p.tickets {
		allFieldsValid := true
		for _, val := range ticket {
			valid := false
			for _, f := range p.fields {
				if f.ValidForAnyRule(val) {
					valid = true
					break
				}
			}
			if !valid {
				allFieldsValid = false
			}
		}
		if allFieldsValid {
			validTickets = append(validTickets, ticket)
		}
	}
	p.tickets = validTickets
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
	Name  string
	Rules []rule
}

func parseField(line string) field {
	matches := fieldPattern.FindStringSubmatch(line)
	return field{
		Name: matches[1],
		Rules: []rule{
			parseRule(matches[2]),
			parseRule(matches[3]),
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

func parsePuzzle(file string) *puzzle {
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

	return &puzzle{fields: fields, tickets: tickets}
}

func partOne() {
	puzzle := parsePuzzle("./16/input.txt")
	log.Printf("ticket scanning error rate %d", puzzle.TicketScanningErrorRate())
}

func partTwo() {
	puzzle := parsePuzzle("./16/input.txt")
	mapping := puzzle.FindFieldMapping()
	log.Printf("%+v", mapping)
	myTicket := make(map[string]int)
	total := 1
	for i, name := range mapping {
		val := puzzle.tickets[0][i]
		myTicket[name] = puzzle.tickets[0][i]
		if strings.HasPrefix(name, "departure") {
			total *= val
		}
	}
	log.Printf("departure total: %d", total)
}
