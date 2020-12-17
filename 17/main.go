package main

import (
	aoc "aoc2020"
	"log"
)

func main() {
	partOne()
}

type model struct {
	space      map[position]string
	lx, ly, lz int
}

type position struct {
	x, y, z int
}

func partOne() {
	lines := aoc.ReadInput("./17/input.txt")
	model := model{space: make(map[position]string)}
	for y, line := range lines {
		for x, c := range line {
			model.space[position{x: x, y: y}] = string(c)
		}
	}

	for i := 0; i < 6; i++ {
		model = model.Cycle()
		log.Printf("%d active cubes", model.CountActive())

	}
}

func (m model) CountActive() int {
	count := 0
	for _, state := range m.space {
		if state == "#" {
			count++
		}
	}
	return count
}

func (m *model) Cycle() model {
	m.InitNeighbours()
	after := model{space: make(map[position]string)}
	for position, state := range m.space {
		neighbours := m.Neighbours(position)
		activeNeighbours := 0
		for _, neighbourState := range neighbours {
			if neighbourState == "#" {
				activeNeighbours++
			}
		}

		if state == "#" && (activeNeighbours < 2 || activeNeighbours > 3) {
			after.space[position] = "."
		} else if state == "." && activeNeighbours == 3 {
			after.space[position] = "#"
		} else {
			after.space[position] = state
		}
	}
	return after
}

func (m *model) InitNeighbours() {
	for position := range m.space {
		neighbours := m.Neighbours(position)
		for neighbour, neighbourState := range neighbours {
			if _, exists := m.space[neighbour]; !exists {
				m.space[neighbour] = neighbourState
			}
		}
	}
}

func (m model) Neighbours(pos position) map[position]string {
	neighbours := make(map[position]string)

	for tz := pos.z - 1; tz <= pos.z+1; tz++ {
		for ty := pos.y - 1; ty <= pos.y+1; ty++ {
			for tx := pos.x - 1; tx <= pos.x+1; tx++ {
				neighbour := position{x: tx, y: ty, z: tz}
				if neighbour == pos {
					continue
				}

				val, exists := m.space[neighbour]
				if exists {
					neighbours[neighbour] = val
				} else {
					neighbours[neighbour] = "."
				}
			}
		}
	}

	return neighbours
}

func partTwo() {

}
