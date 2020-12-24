package main

import (
	"bufio"
	"fmt"
	"os"
)

type Coord struct {
	x, y int
}

func (c Coord) Add(o Coord) Coord {
	return Coord{c.x + o.x, c.y + o.y}
}

type Conway struct {
	tiles map[Coord]bool
}

func (c Conway) CountBlack() int {
	black := 0
	for _, color := range c.tiles {
		if color {
			black++
		}
	}
	return black
}

func (c Conway) neighbors(from Coord) []Coord {
	var neighbors []Coord
	neighbors = append(neighbors, from.Add(Coord{1, 0}))
	neighbors = append(neighbors, from.Add(Coord{-1, 0}))
	neighbors = append(neighbors, from.Add(Coord{1, -1}))
	neighbors = append(neighbors, from.Add(Coord{0, -1}))
	neighbors = append(neighbors, from.Add(Coord{-1, 1}))
	neighbors = append(neighbors, from.Add(Coord{0, 1}))
	return neighbors
}

func (c Conway) blackNeighbors(from Coord) int {
	black := 0
	for _, coord := range c.neighbors(from) {
		if c.tiles[coord] {
			black++
		}
	}
	return black
}

func (c Conway) maybeSetBlack(coord Coord, next map[Coord]bool) {
	if _, ok := next[coord]; !ok {
		black := c.blackNeighbors(coord)
		if c.tiles[coord] && (black == 0 || black > 2) {
			next[coord] = false
		} else if !c.tiles[coord] && black == 2 {
			next[coord] = true
		} else if c.tiles[coord] {
			next[coord] = true
		}
	}
}

func (c *Conway) Step() {
	next := make(map[Coord]bool)
	for coord, color := range c.tiles {
		if color {
			c.maybeSetBlack(coord, next)
			for _, neighbor := range c.neighbors(coord) {
				c.maybeSetBlack(neighbor, next)
			}
		}
	}
	c.tiles = next
}

func main() {
	// white = false
	tiles := make(map[Coord]bool)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		pos := Coord{0, 0}
		for i := 0; i < len(line); i++ {
			var delta Coord
			switch line[i] {
			case 'e':
				delta = Coord{1, 0}
			case 'w':
				delta = Coord{-1, 0}
			default:
				switch line[i : i+2] {
				case "se":
					delta = Coord{1, -1}
				case "sw":
					delta = Coord{0, -1}
				case "nw":
					delta = Coord{-1, 1}
				case "ne":
					delta = Coord{0, 1}
				default:
					panic("invalid direction")
				}
				i++
			}
			pos = pos.Add(delta)
		}
		tiles[pos] = !tiles[pos]
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	conway := Conway{tiles}
	fmt.Println(conway.CountBlack())

	for i := 0; i < 100; i++ {
		conway.Step()
	}
	fmt.Println(conway.CountBlack())
}
