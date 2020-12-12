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
	lines := aoc.ReadInput("./10/sample2.txt")
	jolts := aoc.ToIntSlice(lines)
	sort.Slice(jolts, func(i, j int) bool {
		return jolts[i] < jolts[j]
	})

	permutation := make([]int, 0)
	permutations := gen(jolts, permutation)
	for _, p := range permutations {
		if len(p) == 29 {
			log.Printf("%+v", p)
		}
	}
	log.Printf("%d total permutations", len(permutations))
}

func gen(jolts, permu []int) [][]int {
	if len(jolts) == 0 {
		return [][]int{permu}
	}

	if len(permu) == 0 {
		return gen(jolts[1:], append(permu, jolts[0]))
	}

	head := permu[len(permu)-1]

	// count all available branches
	branches := make([]int, 0)
	for i := 0; i < len(jolts); i++ {
		if jolts[i] > head+3 {
			break
		}
		branches = append(branches, i)
	}

	permutations := make([][]int, 0)
	for _, i := range branches {
		branch := make([]int, 0)
		branch = append(permu, jolts[i])
		permutations = append(permutations, gen(jolts[i+1:], branch)...)
	}

	return permutations
}
