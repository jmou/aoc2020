package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Rule struct {
	char     byte
	subrules [][]string
}

type Ruleset struct {
	rules map[string]Rule
}

func (r Ruleset) PartialMatches(msg string, rulesuffix []string) []string {
	if len(rulesuffix) == 0 {
		return []string{msg}
	}
	rule := rulesuffix[0]
	if subrules := r.rules[rule].subrules; len(subrules) > 0 {
		var rests []string
		for _, subrule := range subrules {
			for _, tail := range r.PartialMatches(msg, subrule) {
				for _, rest := range r.PartialMatches(tail, rulesuffix[1:]) {
					rests = append(rests, rest)
				}
			}
		}
		return rests
	} else {
		if len(msg) >= 1 && msg[0] == r.rules[rule].char {
			return r.PartialMatches(msg[1:], rulesuffix[1:])
		} else {
			return nil
		}
	}
}

func main() {
	ruleset := Ruleset{make(map[string]Rule)}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		pieces := strings.SplitN(scanner.Text(), ": ", 2)
		var rule Rule
		if pieces[1][0] == '"' {
			rule.char = pieces[1][1]
		} else {
			subrules := strings.Split(pieces[1], " | ")
			for _, s := range subrules {
				rule.subrules = append(rule.subrules, strings.Split(s, " "))
			}
		}
		ruleset.rules[pieces[0]] = rule
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	var messages []string
	for scanner.Scan() {
		messages = append(messages, scanner.Text())
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	p1 := 0
	for _, msg := range messages {
		for _, matches := range ruleset.PartialMatches(msg, []string{"0"}) {
			if matches == "" {
				p1++
				break
			}
		}
	}
	fmt.Println(p1)

	p2 := 0
	// cannot modify struct values. could store pointers instead
	rule8 := ruleset.rules["8"]
	rule8.subrules = append(rule8.subrules, []string{"42", "8"})
	ruleset.rules["8"] = rule8
	rule11 := ruleset.rules["11"]
	rule11.subrules = append(rule11.subrules, []string{"42", "11", "31"})
	ruleset.rules["11"] = rule11
	for _, msg := range messages {
		for _, matches := range ruleset.PartialMatches(msg, []string{"0"}) {
			if matches == "" {
				p2++
				break
			}
		}
	}
	fmt.Println(p2)
}
