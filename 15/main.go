package main

import (
	aoc "aoc2020"
	"log"
	"strings"
)

func main() {
	partOne := playGame(2020, "17,1,3,16,19,0")
	log.Printf("part one: %d", partOne)

	partTwo := playGame(30000000, "17,1,3,16,19,0")
	log.Printf("part two: %d", partTwo)
}

type game struct {
	lookup  map[int]int
	numbers []int
}

func (g *game) Next() {
	i := len(g.numbers)
	last := g.numbers[i-1]

	result := 0
	if oldI, ok := g.lookup[last]; ok {
		result = i - oldI
	}
	g.lookup[last] = i
	g.numbers = append(g.numbers, result)
}

func startPuzzle(start ...int) *game {
	g := &game{lookup: make(map[int]int), numbers: make([]int, 1)}
	for _, s := range start {
		g.numbers = append(g.numbers, s)
		g.lookup[s] = len(g.numbers)
	}
	return g
}

func playGame(nth int, input string) int {
	start := aoc.ToIntSlice(strings.Split(input, ","))
	game := startPuzzle(start...)

	for i := 0; i < nth-len(start); i++ {
		game.Next()
	}

	return game.numbers[len(game.numbers)-1]

}
