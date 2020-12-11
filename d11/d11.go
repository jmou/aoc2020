package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

type Ruleset interface {
	Neighbors(g [][]byte, i int, j int) int
	MaxNeighbors() int
}

type P1Ruleset struct {
}

func (r P1Ruleset) Neighbors(g [][]byte, i int, j int) int {
	neighbors := 0
	for di := -1; di <= 1; di++ {
		for dj := -1; dj <= 1; dj++ {
			if di == 0 && dj == 0 {
				continue
			}
			in, jn := i+di, j+dj
			if in >= 0 && in < len(g) && jn >= 0 && jn < len(g[i]) &&
				g[in][jn] == '#' {
				neighbors++
			}
		}
	}
	return neighbors
}

func (r P1Ruleset) MaxNeighbors() int {
	return 4
}

type P2Ruleset struct {
}

func (r P2Ruleset) Neighbors(g [][]byte, i int, j int) int {
	neighbors := 0
	for di := -1; di <= 1; di++ {
		for dj := -1; dj <= 1; dj++ {
			if di == 0 && dj == 0 {
				continue
			}
			in, jn := i, j
		loop:
			for {
				in += di
				jn += dj
				if in >= 0 && in < len(g) && jn >= 0 && jn < len(g[i]) {
					// break exits switch, continue does not
					switch g[in][jn] {
					case '#':
						neighbors++
						break loop
					case 'L':
						break loop
					case '.':
						continue
					}
				} else {
					break
				}
			}
		}
	}
	return neighbors
}

func (r P2Ruleset) MaxNeighbors() int {
	return 5
}

type Gol [][]byte

func (g Gol) Duplicate() Gol {
	next := make([][]byte, len(g))
	for i := 0; i < len(g); i++ {
		next[i] = make([]byte, len(g[i]))
		copy(next[i], g[i])
	}
	return next
}

func (g Gol) Occupied() int {
	count := 0
	for i := 0; i < len(g); i++ {
		for j := 0; j < len(g[i]); j++ {
			if g[i][j] == '#' {
				count++
			}
		}
	}
	return count
}

func (g Gol) Print() {
	for i := 0; i < len(g); i++ {
		os.Stdout.Write(g[i])
		os.Stdout.Write([]byte("\n"))
	}
}

func (g *Gol) Step(rules Ruleset) bool {
	changed := false
	next := g.Duplicate()
	for i := 0; i < len(*g); i++ {
		for j := 0; j < len((*g)[i]); j++ {
			switch (*g)[i][j] {
			case '.': // nop
			case 'L':
				if rules.Neighbors(*g, i, j) == 0 {
					next[i][j] = '#'
					changed = true
				}
			case '#':
				if rules.Neighbors(*g, i, j) >= rules.MaxNeighbors() {
					next[i][j] = 'L'
					changed = true
				}
			}
		}
	}
	*g = next
	return changed
}

func main() {
	input, _ := ioutil.ReadAll(os.Stdin)
	grid := Gol(bytes.Split(input, []byte("\n")))
	if len(grid[len(grid)-1]) == 0 {
		grid = grid[:len(grid)-1]
	}
	gridp2 := grid.Duplicate()

	var p1 P1Ruleset
	for grid.Step(p1) {
	}
	fmt.Println(grid.Occupied())

	var p2 P2Ruleset
	for gridp2.Step(p2) {
	}
	fmt.Println(gridp2.Occupied())
}
