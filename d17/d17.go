package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

type Coord interface {
	Adjacencies(includeSelf bool) []Coord
}

type Coord3D struct {
	x, y, z int
}

func (c Coord3D) Adjacencies(includeSelf bool) []Coord {
	var points []Coord
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				if x != 0 || y != 0 || z != 0 || includeSelf {
					points = append(points, Coord3D{c.x + x, c.y + y, c.z + z})
				}
			}
		}
	}
	return points
}

type Coord4D struct {
	x, y, z, w int
}

func (c Coord4D) Adjacencies(includeSelf bool) []Coord {
	var points []Coord
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				for w := -1; w <= 1; w++ {
					if x != 0 || y != 0 || z != 0 || w != 0 || includeSelf {
						points = append(points, Coord4D{c.x + x, c.y + y, c.z + z, c.w + w})
					}
				}
			}
		}
	}
	return points
}

type grid map[Coord]bool

func NewGrid() grid {
	return make(grid)
}

func (g *grid) Set(point Coord) {
	(*g)[point] = true
}

func (g grid) Neighbors(point Coord) int {
	count := 0
	for _, adj := range point.Adjacencies(false) {
		if g[adj] {
			count++
		}
	}
	return count
}

func (g *grid) Step() {
	next := NewGrid()
	for point, _ := range *g {
		for _, adj := range point.Adjacencies(true) {
			if _, ok := next[adj]; !ok {
				neighbors := g.Neighbors(adj)
				if neighbors == 3 || (neighbors == 2 && (*g)[adj]) {
					next.Set(adj)
				}
			}
		}
	}
	*g = next
}

func (g grid) Active() int {
	count := 0
	for _, active := range g {
		if active {
			count++
		}
	}
	return count
}

func main() {
	input, err := ioutil.ReadAll(bufio.NewReader(os.Stdin))
	if err != nil {
		panic(err)
	}
	gridP1 := NewGrid()
	gridP2 := NewGrid()
	for i, x, y := 0, 0, 0; i < len(input); i++ {
		switch input[i] {
		case '\n':
			x = 0
			y++
		case '#':
			gridP1.Set(Coord3D{x, y, 0})
			gridP2.Set(Coord4D{x, y, 0, 0})
			x++
		default:
			x++
		}
	}
	for i := 0; i < 6; i++ {
		gridP1.Step()
	}
	fmt.Println(gridP1.Active())
	for i := 0; i < 6; i++ {
		gridP2.Step()
	}
	fmt.Println(gridP2.Active())
}
