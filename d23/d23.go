package main

import "fmt"

type Game struct {
	cups []int
}

func (g Game) get(delta int) int {
	return g.cups[(delta)%len(g.cups)]
}

func (g Game) decrementLabel(label int) int {
	return (label-1+len(g.cups)-1)%len(g.cups) + 1
}

func indexOf(haystack []int, needle int) int {
	for i, candidate := range haystack {
		if candidate == needle {
			return i
		}
	}
	return -1
}

func (g *Game) Move() {
	destLabel := g.decrementLabel(g.cups[0])
	splice := g.cups[1:4]
	for indexOf(splice, destLabel) >= 0 {
		destLabel = g.decrementLabel(destLabel)
	}
	dest := indexOf(g.cups, destLabel)
	next := make([]int, len(g.cups))
	n := copy(next, g.cups[4:dest+1])
	n += copy(next[n:], g.cups[1:4])
	n += copy(next[n:], g.cups[dest+1:9])
	next[n] = g.cups[0]
	g.cups = next
}

type FastGame struct {
	next []int
	curr int
}

func NewFastGame(cups []int, fill int) FastGame {
	if fill <= len(cups) {
		fill = len(cups)
	}
	g := FastGame{[]int{}, cups[0]}
	g.next = make([]int, fill+1)
	for i := 1; i < len(cups); i++ {
		g.next[cups[i-1]] = cups[i]
	}
	if fill > len(cups)+1 {
		g.next[cups[len(cups)-1]] = len(cups) + 1
		for i := len(cups) + 2; i <= fill; i++ {
			g.next[i-1] = i
		}
		g.next[fill] = cups[0]
	} else {
		g.next[cups[len(cups)-1]] = cups[0]
	}
	return g
}

func (g *FastGame) Move() {
	spliceFirst := g.next[g.curr]
	spliceLast := g.next[g.next[spliceFirst]]
	g.next[g.curr] = g.next[spliceLast]
	dest := g.curr
	for {
		dest--
		if dest == 0 {
			dest = len(g.next) - 1
		}
		if dest != spliceFirst && dest != g.next[spliceFirst] && dest != spliceLast {
			break
		}
	}
	g.next[dest], g.next[spliceLast] = spliceFirst, g.next[dest]
	g.curr = g.next[g.curr]
}

func (g FastGame) FirstCups() [9]int {
	var cups [9]int
	cups[0] = 1
	for i := 1; i < 9; i++ {
		cups[i] = g.next[cups[i-1]]
	}
	return cups
}

func main() {
	// game := Game{[]int{3, 8, 9, 1, 2, 5, 4, 6, 7}}
	game := Game{[]int{4, 6, 9, 2, 1, 7, 5, 3, 8}}
	for i := 0; i < 100; i++ {
		game.Move()
	}
	fmt.Println(game)

	fastGame := NewFastGame([]int{4, 6, 9, 2, 1, 7, 5, 3, 8}, 1000000)
	for i := 0; i < 10_000_000; i++ {
		fastGame.Move()
	}
	firstCups := fastGame.FirstCups()
	fmt.Println(firstCups)
	fmt.Println(firstCups[1] * firstCups[2])
}
