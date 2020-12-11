package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// go doesn't like anonymous structs / tuples
	contains := make(map[string][]struct {
		int
		string
	})
	containedBy := make(map[string][]string)
	for scanner.Scan() {
		pieces := strings.Split(scanner.Text(), "s contain ")
		outer := pieces[0]
		if pieces[1] == "no other bags." {
			// noop?
			continue
		}
		contents := strings.Split(pieces[1], ", ")
		for _, content := range contents {
			re := regexp.MustCompile(`^(\d+) (.*?)s?\.?$`)
			match := re.FindStringSubmatch(content)
			count, _ := strconv.Atoi(match[1])
			inner := match[2]
			contains[outer] = append(contains[outer], struct {
				int
				string
			}{count, inner})
			containedBy[inner] = append(containedBy[inner], outer)
		}
	}

	seen := make(map[string]bool)
	// I think this is sloppy because it mutates the same underlying array
	frontier := containedBy["shiny gold bag"]
	for len(frontier) > 0 {
		inner := frontier[0]
		frontier = frontier[1:]
		if seen[inner] {
			continue
		}
		seen[inner] = true
		for _, outer := range containedBy[inner] {
			frontier = append(frontier, outer)
		}
	}
	fmt.Println(len(seen))

	fmt.Println(bagSize(contains, make(map[string]int), "shiny gold bag") - 1)
}

func bagSize(contains map[string][]struct {
	int
	string
}, sizes map[string]int, bag string) int {
	if size, ok := sizes[bag]; ok {
		return size
	}

	size := 1
	for _, inner := range contains[bag] {
		size += inner.int * bagSize(contains, sizes, inner.string)
	}

	sizes[bag] = size
	return size
}
