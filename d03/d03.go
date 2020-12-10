package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func countTrees(grid []string, right int, down int) int {
	col := 0
	trees := 0
	// Would have been easier to treat grid as a 2d array
	for row, line := range grid {
		if row%down != 0 {
			continue
		}
		if len(line) == 0 {
			break
		}
		if line[col] == '#' {
			trees++
		}
		col += right
		// or col %= len(line)
		if col >= len(line) {
			col -= len(line)
		}
	}
	return trees
}

func main() {
	fileBytes, _ := ioutil.ReadFile("d03/input")
	fileString := string(fileBytes)
	grid := strings.Split(fileString, "\n")

	fmt.Println(countTrees(grid, 3, 1))

	fmt.Println(countTrees(grid, 1, 1) *
		countTrees(grid, 3, 1) *
		countTrees(grid, 5, 1) *
		countTrees(grid, 7, 1) *
		countTrees(grid, 1, 2))
}
