package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func parse() ([2][]int, error) {
	if _, err := fmt.Scanf("Player 1:\n"); err != nil {
		return [2][]int{nil, nil}, err
	}
	var decks [2][]int
	for {
		var d int
		_, err := fmt.Scanf("%d\n", &d)
		// expected on empty line. We seem to consume the newline
		if err != nil {
			break
		}
		decks[0] = append(decks[0], d)
	}
	if _, err := fmt.Scanf("Player 2:\n"); err != nil {
		return [2][]int{nil, nil}, err
	}
	for {
		var d int
		_, err := fmt.Scanf("%d\n", &d)
		if err == io.EOF {
			break
		} else if err != nil {
			return [2][]int{nil, nil}, err
		}
		decks[1] = append(decks[1], d)
	}
	return decks, nil
}

type RecursiveCombat struct {
	decks [2][]int
	seen  map[string]bool
}

func copySlice(in []int) []int {
	out := make([]int, len(in))
	copy(out, in)
	return out
}

func (r *RecursiveCombat) Round() {
	var winner int
	if len(r.decks[0]) > r.decks[0][0] && len(r.decks[1]) > r.decks[1][0] {
		recursive := [2][]int{copySlice(r.decks[0][1 : r.decks[0][0]+1]),
			copySlice(r.decks[1][1 : r.decks[1][0]+1])}
		combat := RecursiveCombat{recursive, make(map[string]bool)}
		winner = combat.Game()
	} else if r.decks[0][0] > r.decks[1][0] {
		winner = 0
	} else {
		winner = 1
	}
	r.decks[winner] = append(r.decks[winner], r.decks[winner][0])
	r.decks[winner] = append(r.decks[winner], r.decks[1-winner][0])
	r.decks[0] = r.decks[0][1:]
	r.decks[1] = r.decks[1][1:]
}

func (r *RecursiveCombat) Game() int {
	for {
		var serialized strings.Builder
		for _, deck := range r.decks {
			for _, c := range deck {
				fmt.Fprintf(&serialized, "%d,", c)
			}
			serialized.WriteString("|")
		}
		if r.seen[serialized.String()] {
			return 0
		}
		r.seen[serialized.String()] = true

		r.Round()
		if len(r.decks[0]) == 0 {
			return 1
		} else if len(r.decks[1]) == 0 {
			return 0
		}
	}
}

func (r RecursiveCombat) Score(player int) int {
	score := 0
	for i, c := range r.decks[player] {
		score += (len(r.decks[player]) - i) * c
	}
	return score
}

func main() {
	decks, err := parse()
	if err != nil {
		log.Fatal(err)
	}
	// does this actually copy?
	originalDecks := decks

	for len(decks[0]) > 0 && len(decks[1]) > 0 {
		if decks[0][0] > decks[1][0] {
			decks[0] = append(decks[0], decks[0][0])
			decks[0] = append(decks[0], decks[1][0])
		} else {
			decks[1] = append(decks[1], decks[1][0])
			decks[1] = append(decks[1], decks[0][0])
		}
		decks[0] = decks[0][1:]
		decks[1] = decks[1][1:]
	}

	winningDeck := decks[0]
	if len(decks[0]) == 0 {
		winningDeck = decks[1]
	}
	score := 0
	for i, c := range winningDeck {
		score += (len(winningDeck) - i) * c
	}
	fmt.Println(score)

	decks = originalDecks
	combat := RecursiveCombat{decks, make(map[string]bool)}
	winner := combat.Game()
	fmt.Println(combat.Score(winner))
}
