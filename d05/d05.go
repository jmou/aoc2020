package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func seatId(pass string) int {
	rowBin := strings.ReplaceAll(pass[:7], "B", "1")
	rowBin = strings.ReplaceAll(rowBin, "F", "0")
	row, _ := strconv.ParseInt(rowBin, 2, 8)
	colBin := strings.ReplaceAll(pass[7:10], "R", "1")
	colBin = strings.ReplaceAll(colBin, "L", "0")
	col, _ := strconv.ParseInt(colBin, 2, 4)
	return int(row*8 + col)
}

func main() {
	file, _ := os.Open("d05/input")
	scanner := bufio.NewScanner(file)

	seats := []int{}
	for scanner.Scan() {
		seats = append(seats, seatId(scanner.Text()))
	}
	sort.Ints(seats)

	fmt.Println(seats[len(seats)-1])

	for i, seat := range seats {
		if seats[i+1] == seat+2 {
			fmt.Println(seat)
			break
		}
	}
}
