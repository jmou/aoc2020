package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	name     string
	lo1, hi1 int
	lo2, hi2 int
}

func (r Rule) Valid(v int) bool {
	return (v >= r.lo1 && v <= r.hi1) || (v >= r.lo2 && v <= r.hi2)
}

type Ticket []int

type Notes struct {
	rules  []Rule
	your   Ticket
	nearby []Ticket
}

func parseTicket(line string) (Ticket, error) {
	var parsed []int
	fields := strings.Split(line, ",")
	for _, field := range fields {
		num, err := strconv.Atoi(field)
		if err != nil {
			return nil, err
		}
		parsed = append(parsed, num)
	}
	return parsed, nil
}

func Parse(r io.Reader) (*Notes, error) {
	scanner := bufio.NewScanner(r)
	var rules []Rule
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		pieces := strings.SplitN(scanner.Text(), ": ", 2)
		if len(pieces) != 2 {
			return nil, errors.New("rule parse failed")
		}
		var lo1, hi1, lo2, hi2 int
		_, err := fmt.Sscanf(pieces[1], "%d-%d or %d-%d", &lo1, &hi1, &lo2, &hi2)
		if err != nil {
			return nil, err
		}
		rules = append(rules, Rule{pieces[0], lo1, hi1, lo2, hi2})
	}

	if !scanner.Scan() || scanner.Text() != "your ticket:" || !scanner.Scan() || scanner.Err() != nil {
		if scanner.Err() != nil {
			return nil, scanner.Err()
		}
		return nil, errors.New("your ticket parse failed")
	}
	your, err := parseTicket(scanner.Text())
	if err != nil {
		return nil, err
	}

	var nearby []Ticket
	if !scanner.Scan() || scanner.Text() != "" ||
		!scanner.Scan() || scanner.Text() != "nearby tickets:" || scanner.Err() != nil {
		if scanner.Err() != nil {
			return nil, scanner.Err()
		}
		return nil, errors.New("nearby tickets parse failed")
	}
	for scanner.Scan() {
		ticket, err := parseTicket(scanner.Text())
		if err != nil {
			return nil, err
		}
		nearby = append(nearby, ticket)
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return &Notes{rules, your, nearby}, nil
}

func P1(notes *Notes) int {
	sum := 0
	var invalidNearby []int
	for i, ticket := range notes.nearby {
		for _, field := range ticket {
			valid := false
			for _, rule := range notes.rules {
				if rule.Valid(field) {
					valid = true
				}
			}
			if !valid {
				invalidNearby = append(invalidNearby, i)
				sum += field
			}
		}
	}
	i, j, k := 0, 0, 0
	for ; i < len(notes.nearby); i++ {
		if j < len(invalidNearby) && i == invalidNearby[j] {
			j++
		} else {
			notes.nearby[k] = notes.nearby[i]
			k++
		}
	}
	notes.nearby = notes.nearby[:k]
	return sum
}

func P2(notes Notes) int {
	invalidity := make([][]bool, len(notes.rules))
	for i := 0; i < len(invalidity); i++ {
		invalidity[i] = make([]bool, len(notes.your))
	}
	for r, rule := range notes.rules {
		for _, ticket := range notes.nearby {
			for f, field := range ticket {
				if !invalidity[r][f] && !rule.Valid(field) {
					invalidity[r][f] = true
				}
			}
		}
	}

	mapping := make(map[int]int)
	for len(mapping) < len(notes.rules) {
	rule:
		for r := 0; r < len(invalidity); r++ {
			slot := -1
			for f := 0; f < len(invalidity[r]); f++ {
				if !invalidity[r][f] {
					if slot != -1 {
						continue rule
					}
					slot = f
				}
			}
			if slot != -1 {
				mapping[r] = slot
				for i := 0; i < len(invalidity); i++ {
					invalidity[i][slot] = true
				}
			}
		}
	}

	product := 1
	for r, rule := range notes.rules {
		if strings.HasPrefix(rule.name, "departure ") {
			product *= notes.your[mapping[r]]
		}
	}
	return product
}

func main() {
	notes, err := Parse(os.Stdin)
	if err != nil {
		panic(err)
	}
	fmt.Println(P1(notes))
	fmt.Println(P2(*notes))
}
