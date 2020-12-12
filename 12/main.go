package main

import (
	aoc "aoc2020"
	"container/ring"
	"log"
	"math"
	"regexp"
)

func main() {
	partTwo()
}

var instructionPattern = regexp.MustCompile(`^(\w)(\d+)`)
var facings = []string{"N", "E", "S", "W"}

func partOne() {
	lines := aoc.ReadInput("./12/input.txt")
	compass := ring.New(4)
	for i := 0; i < len(facings); i++ {
		compass.Value = facings[i]
		compass = compass.Next()
	}
	// face east
	compass = compass.Move(1)
	ship := &ship{compass: compass}

	for _, line := range lines {
		matches := instructionPattern.FindStringSubmatch(line)
		dir := matches[1]
		val := aoc.Atoi(matches[2])
		switch dir {
		case "N", "S", "E", "W", "F":
			oldX, oldY := ship.Position()
			posX, posY := ship.Move(dir, val)
			log.Printf("moving %s%s from (%d,%d) to (%d,%d)", matches[1], matches[2], oldX, oldY, posX, posY)
		case "L":
			old := ship.Heading()
			ship.Rotate(val * -1)
			log.Printf("turned %s%s from %s to %s", matches[1], matches[2], old, ship.Heading())
		case "R":
			old := ship.Heading()
			ship.Rotate(val)
			log.Printf("turned %s%s from %s to %s", matches[1], matches[2], old, ship.Heading())
		}
	}

	log.Printf("%d", ship.Manhattan())
}

type ship struct {
	x, y    int
	compass *ring.Ring
}

func (s *ship) Heading() string {
	return s.compass.Value.(string)
}

func (s *ship) Manhattan() int {
	return int(math.Abs(float64(s.x)) + math.Abs(float64(s.y)))
}

func (s *ship) Position() (int, int) {
	return s.x, s.y
}

func (s *ship) Rotate(angle int) {
	s.compass = s.compass.Move(angle / 90)
}

func (s *ship) Move(direction string, distance int) (int, int) {
	switch direction {
	case "N":
		s.y += distance
	case "S":
		s.y -= distance
	case "E":
		s.x += distance
	case "W":
		s.x -= distance
	case "F":
		s.Move(s.Heading(), distance)
	}
	return s.Position()
}

func partTwo() {
	lines := aoc.ReadInput("./12/input.txt")
	compass := ring.New(4)
	for i := 0; i < len(facings); i++ {
		compass.Value = facings[i]
		compass = compass.Next()
	}
	// face east
	compass = compass.Move(1)
	ship := &shipWithWaypoint{ship: &ship{compass: compass}, w: &waypoint{x: 10, y: 1}}

	for _, line := range lines {
		matches := instructionPattern.FindStringSubmatch(line)
		dir := matches[1]
		val := aoc.Atoi(matches[2])
		switch dir {
		case "N", "S", "E", "W":
			oldX, oldY := ship.Waypoint()
			ship.Move(dir, val)
			posX, posY := ship.Waypoint()
			log.Printf("waypoint moving %s%d from (%d,%d) to (%d,%d)", dir, val, oldX, oldY, posX, posY)
		case "F":
			oldX, oldY := ship.Position()
			posX, posY := ship.Move(dir, val)
			log.Printf("ship moving %s%d from (%d,%d) to (%d,%d)", dir, val, oldX, oldY, posX, posY)
		case "L":
			oldX, oldY := ship.Waypoint()
			ship.Rotate(val * -1)
			posX, posY := ship.Waypoint()
			log.Printf("waypoint turning %s%d from (%d,%d) to (%d,%d)", dir, val, oldX, oldY, posX, posY)
		case "R":
			oldX, oldY := ship.Waypoint()
			ship.Rotate(val)
			posX, posY := ship.Waypoint()
			log.Printf("turned %s%d from (%d,%d) to (%d,%d)", dir, val, oldX, oldY, posX, posY)
		}
	}

	log.Printf("%d", ship.Manhattan())
}

type waypoint struct {
	x, y int
}

type shipWithWaypoint struct {
	*ship
	w *waypoint
}

func (s *shipWithWaypoint) Move(direction string, distance int) (int, int) {
	switch direction {
	case "N":
		s.w.y += distance
	case "S":
		s.w.y -= distance
	case "E":
		s.w.x += distance
	case "W":
		s.w.x -= distance
	case "F":
		s.x += s.w.x * distance
		s.y += s.w.y * distance
	}

	return s.Position()
}

func (s *shipWithWaypoint) Waypoint() (int, int) {
	return s.w.x, s.w.y
}

func (s *shipWithWaypoint) Rotate(degrees int) {
	radians := float64(degrees) * (math.Pi / 180)

	// shamelessly copied from https://stackoverflow.com/questions/2259476/rotating-a-point-about-another-point-2d
	sin := math.Sin(radians)
	cos := math.Cos(radians)

	wx := float64(s.w.x)
	wy := float64(s.w.y)

	// rotate point CLOCKWISE (https://stackoverflow.com/a/25196651/1392312)
	newX := wx*cos + wy*sin
	newY := -wx*sin + wy*cos

	s.w.x = int(math.Round(newX))
	s.w.y = int(math.Round(newY))
}
