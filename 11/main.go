package main

import (
	aoc "aoc2020"
	"log"
)

func main() {
	partTwo()
}

func partOne() {
	lines := aoc.ReadInput("./11/input.txt")
	seating := lines
	for i := 0; i < 1000; i++ {
		newSeating := beSeated(seating)
		printSeats(newSeating)
		log.Println()

		if equalSeats(newSeating, seating) {
			seating = newSeating
			log.Println("found stable seating plan")
			break
		}
		seating = newSeating
	}

	log.Printf("%d occupied seats", countOccupiedSeats(seating))
}

func partTwo() {
	lines := aoc.ReadInput("./11/input.txt")
	seating := lines
	for i := 0; i < 1000; i++ {
		newSeating := beSeatedTwo(seating)

		if equalSeats(newSeating, seating) {
			seating = newSeating
			log.Printf("found stable seating plan in after %d iterations\n", i)
			break
		}
		seating = newSeating
	}

	log.Printf("%d occupied seats", countOccupiedSeats(seating))
}

func equalSeats(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for ai, aLine := range a {
		if aLine != b[ai] {
			return false
		}
	}

	return true
}

func printSeats(lines []string) {
	for _, line := range lines {
		log.Printf("%s", line)
	}
}

func countOccupiedSeats(lines []string) int {
	count := 0
	for _, line := range lines {
		for _, c := range line {
			if string(c) == "#" {
				count++
			}
		}
	}
	return count
}

func beSeated(lines []string) []string {
	after := make([]string, len(lines))
	for y, line := range lines {
		for x, c := range line {
			result := "."
			switch string(c) {
			case "L":
				result = "L"
				if count(getAdjacent(lines, x, y), "#") == 0 {
					result = "#"
				}

			case "#":
				result = "#"
				if count(getAdjacent(lines, x, y), "#") >= 4 {
					result = "L"
				}
			}

			after[y] = after[y] + result
		}
	}

	return after
}

func beSeatedTwo(lines []string) []string {
	after := make([]string, len(lines))
	for y, line := range lines {
		for x, c := range line {
			result := "."
			switch string(c) {
			case "L":
				result = "L"
				if count(getVisible(lines, x, y), "#") == 0 {
					result = "#"
				}

			case "#":
				result = "#"
				if count(getVisible(lines, x, y), "#") >= 5 {
					result = "L"
				}
			}

			after[y] = after[y] + result
		}
	}

	return after
}

func getAdjacent(lines []string, x, y int) []string {
	adjacent := make([]string, 0)
	// far right
	if x != len(lines[y])-1 {
		adjacent = append(adjacent, string(lines[y][x+1]))
	}

	// bottom right corner
	if y != len(lines)-1 && x != len(lines[y])-1 {
		adjacent = append(adjacent, string(lines[y+1][x+1]))
	}

	// bottom
	if y != len(lines)-1 {
		adjacent = append(adjacent, string(lines[y+1][x]))
	}

	// bottom left corner
	if y != len(lines)-1 && x != 0 {
		adjacent = append(adjacent, string(lines[y+1][x-1]))
	}

	// far left
	if x != 0 {
		adjacent = append(adjacent, string(lines[y][x-1]))
	}

	// top left corner
	if y != 0 && x != 0 {
		adjacent = append(adjacent, string(lines[y-1][x-1]))
	}

	// top
	if y != 0 {
		adjacent = append(adjacent, string(lines[y-1][x]))
	}

	// top right corner
	if y != 0 && x != len(lines[y])-1 {
		adjacent = append(adjacent, string(lines[y-1][x+1]))
	}

	return adjacent
}

func getVisible(lines []string, x, y int) []string {
	visible := make([]string, 0)
	// far right
	for xx := x + 1; xx < len(lines[y]); xx++ {
		if string(lines[y][xx]) != "." {
			visible = append(visible, string(lines[y][xx]))
			break
		}
	}

	// bottom right corner
	for yy, xx := y+1, x+1; yy < len(lines) && xx < len(lines[y]); yy, xx = yy+1, xx+1 {
		if string(lines[yy][xx]) != "." {
			visible = append(visible, string(lines[yy][xx]))
			break
		}
	}

	// bottom
	for yy := y + 1; yy < len(lines[y]); yy++ {
		if string(lines[yy][x]) != "." {
			visible = append(visible, string(lines[yy][x]))
			break
		}
	}

	// bottom left corner
	for yy, xx := y+1, x-1; yy < len(lines) && xx >= 0; yy, xx = yy+1, xx-1 {
		if string(lines[yy][xx]) != "." {
			visible = append(visible, string(lines[yy][xx]))
			break
		}
	}

	// far left
	for xx := x - 1; xx >= 0; xx-- {
		if string(lines[y][xx]) != "." {
			visible = append(visible, string(lines[y][xx]))
			break
		}
	}

	// top left corner
	for yy, xx := y-1, x-1; yy >= 0 && xx >= 0; yy, xx = yy-1, xx-1 {
		if string(lines[yy][xx]) != "." {
			visible = append(visible, string(lines[yy][xx]))
			break
		}
	}

	// top
	for yy := y - 1; yy >= 0; yy-- {
		if string(lines[yy][x]) != "." {
			visible = append(visible, string(lines[yy][x]))
			break
		}
	}

	// top right corner
	for yy, xx := y-1, x+1; yy >= 0 && xx < len(lines[y]); yy, xx = yy-1, xx+1 {
		if string(lines[yy][xx]) != "." {
			visible = append(visible, string(lines[yy][xx]))
			break
		}
	}

	return visible
}

func count(elements []string, elem string) int {
	count := 0
	for _, e := range elements {
		if e == elem {
			count++
		}
	}
	return count
}
