package main

import (
	aoc "aoc2020"
	"log"
)

func main() {
	lines := aoc.ReadInput("./1/input.txt")
	input := aoc.ToIntSlice(lines)

	first, second := twoSumEquals2020(input)
	log.Printf("first: %d\n", first)
	log.Printf("second: %d\n", second)
	log.Printf("result: %d\n", first*second)

	nums := threeSumEquals2020(input)
	log.Printf("%+v\n", nums)
	multiplied := 1
	for _, num := range nums {
		log.Printf("%d\n", num)
		multiplied *= num
	}
	log.Printf("result: %d\n", multiplied)
}

func twoSumEquals2020(input []int) (int, int) {
	for _, first := range input {
		for _, second := range input {
			if first+second == 2020 {
				return first, second
			}
		}
	}

	return 0,0
}

func threeSumEquals2020(input []int) []int {
	for _, first := range input {
		for _, second := range input {
			for _, third := range input {
				if first+second+third == 2020 {
					return []int{first, second, third}
				}
			}
		}
	}

	return nil
}