package main

import (
	aoc "aoc2020"
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
)

var (
	memLine = regexp.MustCompile(`^mem\[(\d+)\] = (\d+)$`)
)

func main() {
	partTwo()
}

func partOne() {
	lines := aoc.ReadInput("./14/input.txt")
	mem := make(map[int]int)
	var m mask
	for _, line := range lines {
		if !memLine.MatchString(line) {
			m = parseMask(line)
			continue
		}
		matches := memLine.FindStringSubmatch(line)
		addr := aoc.Atoi(matches[1])
		val := uint64(aoc.Atoi(matches[2]))

		log.Printf("%036b (%d) mem[%d]", val, int(val), addr)
		log.Printf("%s mask", m)
		applied := m.apply(val)
		log.Printf("%036b (%d) mem[%d]", applied, int(applied), addr)
		log.Println()
		mem[addr] = int(applied)
	}

	total := 0
	for _, val := range mem {
		total += val
	}
	log.Printf("total mem is %d", total)
}

type mask struct {
	orig     string
	zeroMask uint64
	oneMask  uint64
}

func parseMask(maskLine string) mask {
	m := strings.Split(maskLine, " = ")
	zerosForXs := strings.ReplaceAll(m[1], "0", "X")
	onesForXs := strings.ReplaceAll(m[1], "1", "X")
	zeroesToOnes := strings.ReplaceAll(onesForXs, "0", "1")

	zeroMask, err := strconv.ParseUint(strings.ReplaceAll(zeroesToOnes, "X", "0"), 2, 36)
	if err != nil {
		panic(err)
	}

	oneMask, err := strconv.ParseUint(strings.ReplaceAll(zerosForXs, "X", "0"), 2, 36)
	if err != nil {
		panic(err)
	}

	return mask{orig: m[1], zeroMask: zeroMask, oneMask: oneMask}
}

func (m mask) apply(num uint64) uint64 {
	// add any ones
	withOnes := num | m.oneMask
	// clear the zeroes
	zeroed := withOnes &^ m.zeroMask
	return zeroed
}

func (m mask) String() string {
	return m.orig
}

type maskV2 struct {
	orig string
}

func parseMaskV2(maskLine string) maskV2 {
	m := strings.Split(maskLine, " = ")
	masks := make([]mask, 0)
	masks = append(masks, parseMask(maskLine))
	return maskV2{orig: m[1]}
}

func (m maskV2) String() string {
	return m.orig
}

func (m maskV2) apply(num uint64) []uint64 {
	mask := m.orig
	result := ""
	addr := fmt.Sprintf("%036b", num)
	exes := make([]int, 0)
	for i := 0; i < len(mask); i++ {
		switch string(mask[i]) {
		case "1":
			result += "1"
		case "0":
			result += string(addr[i])
		case "X":
			result += "X"
			exes = append(exes, i)
		}
	}

	stringResults := rotatingReplace(result, exes)
	results := make([]uint64, len(stringResults))
	for i, r := range stringResults {
		res, err := strconv.ParseUint(r, 2, 36)
		if err != nil {
			panic(err)
		}

		results[i] = res
	}

	return results
}

func rotatingReplace(orig string, exes []int) []string {
	length := len(exes)
	rs := make([]string, 0)
	for i := 0; i < int(math.Pow(2, float64(length))); i++ {
		r := []rune(orig)
		replacements := fmt.Sprintf("%0*b", length, i)
		for i, target := range exes {
			r[target] = rune(replacements[i])
		}
		rs = append(rs, string(r))
	}
	return rs
}

func partTwo() {
	lines := aoc.ReadInput("./14/input.txt")
	mem := make(map[int]int)
	var m maskV2
	for _, line := range lines {
		if !memLine.MatchString(line) {
			m = parseMaskV2(line)
			continue
		}
		matches := memLine.FindStringSubmatch(line)
		addr := aoc.Atoi(matches[1])
		val := aoc.Atoi(matches[2])

		log.Printf("%036b mem[%d] = %d", addr, addr, val)
		log.Printf("%s mask", m)
		for _, applied := range m.apply(uint64(addr)) {
			log.Printf("%036b mem[%d] = %d", applied, int(applied), val)
			mem[int(applied)] = val
		}
		log.Println()
	}

	total := 0
	for _, val := range mem {
		total += val
	}
	log.Printf("total mem is %d", total)
}
