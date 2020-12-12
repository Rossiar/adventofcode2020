package main

import (
	aoc "aoc2020"
	"log"
	"regexp"
	"strings"
)

var bagRulePattern = regexp.MustCompile(`^(\d) (\w+ \w+) (bag|bags)(|\.)$`)

func main() {
	partTwo()
}

func partOne() {
	lines := aoc.ReadInput("./7/input.txt")
	allRules := make(BagSpace)
	for _, line := range lines {
		bag, rules := parseRule(line)
		allRules[bag] = rules
	}


	canHoldShinyGold := allRules.BagsCanHold("shiny gold")
	log.Printf("%d bags can hold a shiny gold bag\n", len(canHoldShinyGold))
}

func partTwo() {
	lines := aoc.ReadInput("./7/input.txt")
	allRules := make(BagSpace)
	for _, line := range lines {
		bag, rules := parseRule(line)
		allRules[bag] = rules
	}

	log.Printf("shiny gold has %d bags\n", allRules.CountSubBags("shiny gold"))
}

type BagSpace map[string]map[string]int

func (bs BagSpace) CountSubBags(bag string) int {
	cnt := 0
	for childBag, num := range bs[bag] {
		childCnt := bs.CountSubBags(childBag)
		cnt += num+(childCnt*num)
	}

	return cnt
}

func (bs BagSpace) BagsCanHold(bag string) map[string]int {
	canHold := make(map[string]int, 0)
	for parentBag, rules := range bs {
		_, ok := rules[bag]
		if ok {
			canHold[parentBag] = 1
			for grandParentBag, _ := range bs.BagsCanHold(parentBag) {
				canHold[grandParentBag] = 1
			}
		}
	}

	return canHold
}

func parseRule(line string) (string, map[string]int) {
	rules := make(map[string]int)
	split := strings.Split(line, " bags contain ")
	bag := split[0]
	contains := split[1]
	if contains == "no other bags." {
		return bag, rules
	}

	for _, childBag := range strings.Split(contains, ", ") {
		matches := bagRulePattern.FindStringSubmatch(childBag)
		rules[matches[2]] = aoc.Atoi(matches[1])
	}
	return bag, rules
}