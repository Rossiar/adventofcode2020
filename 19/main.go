package main

import (
	aoc "aoc2020"
	"fmt"
	"log"
	"regexp"
	"strings"
)

func main() {
	partTwo()
}

func partOne() {
	lines := aoc.ReadInput("./19/sample2.txt")
	rules := Rules{}
	messageStart := -1
	for i, line := range lines {
		if line == "" {
			messageStart = i + 1
			break
		}

		raw := strings.Split(line, ":")
		rule := parseRule(raw[1])
		rule.id = raw[0]
		rules[rule.id] = rule
	}

	resolved := strings.ReplaceAll(rules["0"].Resolve(rules), " ", "")
	toMatch := regexp.MustCompile(fmt.Sprintf("^%s$", resolved))
	log.Printf(toMatch.String())
	matches := 0
	for i := messageStart; i < len(lines); i++ {
		if toMatch.MatchString(lines[i]) {
			matches++
		}
	}
	log.Printf("%d matches", matches)
}

func partTwo() {
	lines := aoc.ReadInput("./19/input2.txt")
	rules := Rules{}
	messageStart := -1
	for i, line := range lines {
		if line == "" {
			messageStart = i + 1
			break
		}

		raw := strings.Split(line, ":")
		rule := parseRule(raw[1])
		rule.id = raw[0]
		rules[rule.id] = rule
	}

	thirtyOneRaw := strings.ReplaceAll(rules["31"].Resolve(rules), " ", "")
	fortyTwoRaw := strings.ReplaceAll(rules["42"].Resolve(rules), " ", "")
	matches := 0
	for i := messageStart; i < len(lines); i++ {
		line := lines[i]

		// horror
		for i := 1; i < 10; i++ {
			// always one extra 42 match due to rule 8
			fortyTwoMatches := i + 1
			// there must be at least ix 32 matches (for rule 11
			thirtyOneMatches := i
			toMatch := regexp.MustCompile(fmt.Sprintf("^%s{%d,}%s{%d}$", fortyTwoRaw, fortyTwoMatches,
				thirtyOneRaw, thirtyOneMatches))
			if toMatch.MatchString(line) {
				matches++
				//log.Printf("matched %s with %s", line, toMatch.String())
				break
			}
		}
	}
	log.Printf("%d matches", matches)
}

type Rules = map[string]*rule

type rule struct {
	id      string
	matches string
}

func (r *rule) Resolve(rules Rules) string {
	resolved := idPattern.ReplaceAllStringFunc(r.matches, func(s string) string {
		sr := rules[s]
		return sr.Resolve(rules)
	})
	if strings.Contains(resolved, "|") {
		resolved = fmt.Sprintf("(%s)", resolved)
	}

	return resolved
}

var (
	simpleRulePattern = regexp.MustCompile(`"(\w)"`)
	idPattern         = regexp.MustCompile(`\d+`)
)

func parseRule(line string) *rule {
	if matches := simpleRulePattern.FindStringSubmatch(line); matches != nil {
		return &rule{
			matches: matches[1],
		}
	}

	return &rule{
		matches: line,
	}
}
