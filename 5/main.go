package main

import (
	aoc "aoc2020"
	"log"
	"regexp"
)

func main() {
	partTwo()
}

func partOne() {
	lines := aoc.ReadInput("./5/input.txt")
	max := 0
	for _, line := range lines {
		num := parseSeatId(line)
		if num > max {
			max = num
			log.Printf("new highest seat %d from %s\n", max, line)
		}
	}
}

func partTwo() {
	lines := aoc.ReadInput("./5/input.txt")
	seats := make([]bool, 1023)
	for _, line := range lines {
		num := parseSeatId(line)
		seats[num] = true
	}

	for i, seat := range seats[0:len(seats)-1] {
		if !seat && seats[i+1] && seats[i-1] {
			log.Printf("seat %d is empty\n", i)
		}
	}
}

var seatNumberPattern = regexp.MustCompile(`^(F|B){7}(L|R){3}$`)

func parseSeatId(seatNumber string) int {
	if !seatNumberPattern.MatchString(seatNumber) {
		panic("bad seat number")
	}

	row := parse(seatNumber[:7], "B", "F", makeRange(0, 127))
	column := parse(seatNumber[7:], "R", "L", makeRange(0, 7))
	return row * 8 + column
}

func parse(directions string, upper, lower string, space []int) int {
	if len(directions) == 0 {
		return space[0]
	}

	switch string(directions[0]) {
	case upper:
		return parse(directions[1:], upper, lower, space[len(space)/2:])
	case lower:
		return parse(directions[1:], upper, lower, space[:len(space)/2])
	default:
		panic("unexpected value")
	}
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}