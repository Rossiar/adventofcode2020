package main

import (
	aoc "aoc2020"
	"log"
	"strings"
)

func main() {
	partOne()
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

func partOne() {
	start := aoc.ToIntSlice(strings.Split("17,1,3,16,19,0", ","))
	game := startPuzzle(start...)

	for i := 0; i < 2020-len(start); i++ {
		game.Next()
	}

	log.Printf("%dth number was %d", len(game.numbers)-1, game.numbers[len(game.numbers)-1])
}

func partTwo() {

}
