package main

import (
	"fmt"
)

func IntAbs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

type Coord struct {
	x, y int
}

func (c Coord) Manhattan() int {
	return IntAbs(c.x) + IntAbs(c.y)
}

func (c *Coord) Rotate(n int) {
	for ; n > 0; n-- {
		c.x, c.y = c.y, -c.x
	}
}

type P1Ship struct {
	Coord
	dir int
}

func (s *P1Ship) Step(action byte, value int) {
	cardinals := []byte{'E', 'S', 'W', 'N'}
	switch action {
	case 'N':
		s.y += value
	case 'S':
		s.y -= value
	case 'E':
		s.x += value
	case 'W':
		s.x -= value
	case 'L':
		s.dir += 4 - value/90
		s.dir %= 4
	case 'R':
		s.dir += value / 90
		s.dir %= 4
	case 'F':
		s.Step(cardinals[s.dir], value)
	}
}

type P2Ship struct {
	Coord
	waypoint Coord
}

func (s *P2Ship) Step(action byte, value int) {
	switch action {
	case 'N':
		s.waypoint.y += value
	case 'S':
		s.waypoint.y -= value
	case 'E':
		s.waypoint.x += value
	case 'W':
		s.waypoint.x -= value
	case 'L':
		s.waypoint.Rotate(4 - value/90)
	case 'R':
		s.waypoint.Rotate(value / 90)
	case 'F':
		s.x += value * s.waypoint.x
		s.y += value * s.waypoint.y
	}
}

func main() {
	p1Ship := P1Ship{Coord{0, 0}, 0}
	p2Ship := P2Ship{Coord{0, 0}, Coord{10, 1}}
	for {
		var action byte
		var value int
		if _, err := fmt.Scanf("%c%d\n", &action, &value); err != nil {
			break
		}
		p1Ship.Step(action, value)
		p2Ship.Step(action, value)
	}
	fmt.Println(p1Ship.Manhattan())
	fmt.Println(p2Ship.Manhattan())
}
