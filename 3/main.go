package main

import (
	aoc "aoc2020"
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	partTwo()
}

func partOne() {
	lines := aoc.ReadInput("./3/input.txt")
	trees := howManyTreesHit(lines, 3, 1)
	log.Printf("hit %d trees", trees)
}

func partTwo() {
	slopes := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	total := 1
	for _, slope := range slopes {
		lines := aoc.ReadInput("./3/input.txt")
		trees := howManyTreesHit(lines, slope[0], slope[1])
		total *= trees
		log.Printf("hit %d trees\n", trees)
	}

	log.Printf("total is %d\n", total)
}

func howManyTreesHit(lines []string, xStep, yStep int) int {
	lines = enlarge(lines, xStep)

	x := xStep
	for y := yStep; y < len(lines); y+=yStep {
		marker := "0"
		hit := lines[y][x:x+1]
		if hit == "#" {
			marker = "X"
		}

		lines[y] = lines[y][:x] + marker + lines[y][x:]
		x += xStep
	}

	trees := 0
	for _, line := range lines {
		if strings.Contains(line, "X") {
			trees++
		}
	}

	return trees
}

func enlarge(slope []string, stepSize int) []string {
	maxX := len(slope)*stepSize
	for y := 0; y < len(slope); y++ {
		for len(slope[y]) < maxX  {
			slope[y] += slope[y]
		}
	}
	return slope
}

func writeOut(path string, lines []string) {
	outFile, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	writer := bufio.NewWriter(outFile)
	for _, line := range lines {
		if _, err := writer.WriteString(line + "\n"); err != nil {
			panic(err)
		}
	}
}
