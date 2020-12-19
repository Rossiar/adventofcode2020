package main

import (
	aoc "aoc2020"
	"log"
	"sort"
)

func main() {
	partTwo()
}

func partOne() {
	lines := aoc.ReadInput("./10/input.txt")
	jolts := aoc.ToIntSlice(lines)
	sort.Slice(jolts, func(i, j int) bool {
		return jolts[i] < jolts[j]
	})

	jc := &joltCounter{}
	for i := 0; i < len(jolts); i++ {
		var diff int
		if i != 0 {
			diff = jolts[i] - jolts[i-1]
		} else {
			diff = jolts[i]
		}

		jc.Count(diff)
	}
	// device jolt diff from highest adapter
	jc.Count(3)
	log.Printf("%+v gives result %d", jc, jc.Result())
}

type joltCounter struct {
	oneJolt   int
	threeJolt int
}

func (jc *joltCounter) Count(diff int) {
	switch diff {
	case 1:
		jc.oneJolt++
	case 3:
		jc.threeJolt++
	default:
		log.Printf("jolt of %d", diff)
	}
}

func (jc joltCounter) Result() int {
	return jc.oneJolt * jc.threeJolt
}

func partTwo() {
	lines := aoc.ReadInput("./10/input.txt")
	jolts := aoc.ToIntSlice(lines)
	sort.Ints(jolts)

	jolts = append([]int{0}, jolts...)
	num := mathsWay(jolts, 0, make(map[int]int))
	log.Printf("%d total permutations", num)
}

// https://github.com/chigley/advent2020/blob/master/day10/day10.go
func mathsWay(jolts []int, i int, cache map[int]int) int {
	if n, ok := cache[i]; ok {
		return n
	}

	if i == len(jolts)-1 {
		return 1
	}

	branches := 0
	for j := i + 1; j < len(jolts); j++ {
		if jolts[j] > jolts[i]+3 {
			break
		}

		branches += mathsWay(jolts, j, cache)
	}

	cache[i] = branches
	return branches
}
