package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	inputBytes, _ := ioutil.ReadFile("d06/input")
	inputString := strings.Trim(string(inputBytes), "\n")
	// chomp newline (otherwise off-by-one in group count)
	// if inputString[len(inputString)-1] == '\n' {
	// 	inputString = inputString[:len(inputString)-1]
	// }
	groups := strings.Split(inputString, "\n\n")

	p1 := 0
	p2 := 0
	for _, group := range groups {
		yesses := make(map[rune]int)
		// iterating over a range gives runes
		// indexing gives bytes
		for _, question := range group {
			if question == '\n' {
				continue
			}
			yesses[question]++
		}
		p1 += len(yesses)

		group_size := strings.Count(group, "\n") + 1
		for _, count := range yesses {
			if count == group_size {
				p2++
			}
		}
	}

	fmt.Println(p1)
	fmt.Println(p2)
}
