package main

import (
	"fmt"
	"io"
	"log"

	"github.com/ef-ds/deque"
)

func parse() ([]deque.Deque, error) {
	if _, err := fmt.Scanf("Player 1:\n"); err != nil {
		return nil, err
	}
	var decks []deque.Deque
	var deck deque.Deque
	for {
		var d int
		_, err := fmt.Scanf("%d\n", &d)
		// expected on empty line. We seem to consume the newline
		if err != nil {
			break
		}
		deck.PushFront(d)
	}
	decks = append(decks, deck)
	if _, err := fmt.Scanf("Player 2:\n"); err != nil {
		return nil, err
	}
	deck.Init()
	for {
		var d int
		_, err := fmt.Scanf("%d\n", &d)
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		deck.PushFront(d)
	}
	decks = append(decks, deck)
	return decks, nil
}

func main() {
	decks, err := parse()
	if err != nil {
		log.Fatal(err)
	}
	for decks[0].Len() > 0 && decks[1].Len() > 0 {
		c0, _ := decks[0].PopBack()
		c1, _ := decks[1].PopBack()
		if c0.(int) > c1.(int) {
			decks[0].PushFront(c0)
			decks[0].PushFront(c1)
		} else {
			decks[1].PushFront(c1)
			decks[1].PushFront(c0)
		}
	}

	winningDeck := &decks[0]
	if decks[0].Len() == 0 {
		winningDeck = &decks[1]
	}

	// can't use range with Deque
	i, score := 1, 0
	for winningDeck.Len() > 0 {
		c, _ := winningDeck.PopFront()
		score += i * c.(int)
		i++
	}
	fmt.Println(score)
}
