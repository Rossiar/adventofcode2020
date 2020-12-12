package adventofcode2020

import (
	"bufio"
	"os"
	"strconv"
)

func ReadInput(path string) []string {
	inFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	input := make([]string, 0)
	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	return input
}

func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func Atof(s string) float64 {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return float64(i)
}

func ToIntSlice(strings []string) []int {
	numbers := make([]int, len(strings))
	for i, s := range strings {
		numbers[i] = Atoi(s)
	}
	return numbers
}
