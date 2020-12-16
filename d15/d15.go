package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func Parse() ([]int, error) {
	inputBytes, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return nil, err
	}
	var input []int
	pieces := strings.Split(string(inputBytes), ",")
	for _, piece := range pieces {
		d, err := strconv.Atoi(strings.TrimSpace(piece))
		if err != nil {
			return nil, err
		}
		input = append(input, d)
	}
	return input, nil
}

func Solve(starts []int, iterations int) int {
	turn := 1
	saidAt := make(map[int]int)
	for i := 0; i < len(starts)-1; i++ {
		saidAt[starts[i]] = turn
		turn++
	}
	num := starts[len(starts)-1]
	for ; turn < iterations; turn++ {
		lastNum := num
		if numSaid, ok := saidAt[num]; ok {
			num = turn - numSaid
		} else {
			num = 0
		}
		saidAt[lastNum] = turn
	}
	return num
}

func main() {
	input, err := Parse()
	if err != nil {
		panic(err)
	}
	p1 := Solve(input, 2020)
	fmt.Println(p1)
	p2 := Solve(input, 30000000)
	fmt.Println(p2)
}
