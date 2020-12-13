package main

import (
	aoc "aoc2020"
	"log"
	"strings"
)

func main() {
	partTwo()
}

func partOne() {
	lines := aoc.ReadInput("./13/input.txt")
	estimate := aoc.Atoi(lines[0])
	log.Printf("estimate is %d", estimate)
	routes := strings.Split(lines[1], ",")
	log.Printf("routes are %+v", routes)
	minWait := 1000
	selected := 0
	for _, route := range routes {
		if route == "x" {
			continue
		}

		log.Printf("route %s", route)
		freq := aoc.Atoi(route)
		nextBus := findClosestDivisor(estimate, freq)
		log.Printf("next bus at %d", nextBus)
		timeUntil := nextBus - estimate
		if timeUntil < minWait {
			minWait = timeUntil
			selected = freq
			log.Printf("selecting route %d that departs in %d minutes", freq, timeUntil)
		}
	}

	log.Printf("selected route %d that departs in %d minutes", selected, minWait)
	log.Printf("%d", minWait*selected)
}

func partTwo() {
	lines := aoc.ReadInput("./13/sample2.txt")
	routes := strings.Split(lines[1], ",")
	log.Printf("routes are %+v", routes)
	freq := make([]int, len(routes))
	for i, route := range routes {
		if route == "x" {
			continue
		}

		f := aoc.Atoi(route)
		freq[i] = f
	}

	t := freq[0]
	for {
		pos := check(t, freq)
		if pos == t {
			log.Printf("found %d", t)
			break
		}
		t = pos
	}
}

func check(t int, routes []int) int {
	lastReal := 0
	for i, route := range routes {
		if route == 0 {
			continue
		}

		if (t+i)%route != 0 {
			return findClosestDivisor(t, routes[lastReal])
		}
		lastReal = i
	}
	return t
}

func findClosestDivisor(num, divisor int) int {
	return ((num / divisor) * divisor) + divisor
}
