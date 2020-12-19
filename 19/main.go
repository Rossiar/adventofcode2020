package main

import (
	aoc "aoc2020"
	"fmt"
	"log"
	"regexp"
	"strings"
)

func main() {
	partOne()
}

func partOne() {
	lines := aoc.ReadInput("./19/input.txt")
	rules := Rules{}
	messageStart := -1
	for i, line := range lines {
		if line == "" {
			messageStart = i + 1
			break
		}

		raw := strings.Split(line, ":")
		rules[raw[0]] = parseRule(raw[1])
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

type Rules = map[string]*rule

type rule struct {
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

func partTwo() {

}
