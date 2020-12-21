package main

import (
	aoc "aoc2020"
	"log"
	"regexp"
)

func main() {
	partOne()
}

type tile struct {
	id                       string
	left, top, right, bottom string
	neighbours               []string
}

func (t *tile) Match(other *tile) {
	for _, side := range []string{t.top, t.bottom, t.right, t.left} {
		for _, oSide := range []string{other.top, other.bottom, other.right, other.left} {
			if side == oSide || aoc.Reverse(side) == oSide {
				t.neighbours = append(t.neighbours, other.id)
				return
			}
		}
	}
}

var tileIdPattern = regexp.MustCompile(`^Tile (\d+):$`)

func parseTiles(lines []string) []*tile {
	tiles := make([]*tile, 0)
	current := &tile{}
	for _, line := range lines {
		if tileIdPattern.MatchString(line) {
			matches := tileIdPattern.FindStringSubmatch(line)
			current = &tile{id: matches[1]}
		} else if line == "" {
			tiles = append(tiles, current)
		} else {
			if current.top == "" {
				current.top = line
			} else if len(current.right) == 9 {
				current.bottom = line
			}
			current.right += string(line[len(line)-1])
			current.left += string(line[0])
		}
	}
	return tiles
}

func partOne() {
	lines := aoc.ReadInput("./20/input.txt")
	tiles := parseTiles(lines)
	for _, tile := range tiles {
		for _, other := range tiles {
			if tile.id == other.id {
				continue
			}

			tile.Match(other)
		}
	}

	total := 1
	middle := 0
	edge := 0
	corners := 0
	for _, tile := range tiles {
		switch len(tile.neighbours) {
		case 2:
			corners++
		case 3:
			edge++
		case 4:
			middle++
		default:
			log.Printf("tile %s had %+v neighbours??", tile.id, tile.neighbours)
		}
		if len(tile.neighbours) == 2 {
			log.Printf("corner: %s (%+v))", tile.id, tile.neighbours)
			total *= aoc.Atoi(tile.id)
		}
	}

	log.Printf("edges: %d, corners: %d, middle: %d", edge+corners, corners, middle)
	log.Printf("total: %d", total)
}

func partTwo() {

}
