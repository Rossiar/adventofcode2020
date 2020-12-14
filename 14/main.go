package main

import (
	aoc "aoc2020"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var (
	memLine = regexp.MustCompile(`^mem\[(\d+)\] = (\d+)$`)
)

func main() {
	partOne()
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

func partTwo() {

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
