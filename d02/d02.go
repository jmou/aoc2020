package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("d02/input")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	valid_p1 := 0
	valid_p2 := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		pieces := strings.SplitN(scanner.Text(), ": ", 2)
		if len(pieces) != 2 {
			panic(nil)
		}
		password := pieces[1]
		pieces = strings.SplitN(pieces[0], " ", 2)
		if len(pieces) != 2 {
			panic(nil)
		}
		if len(pieces[1]) != 1 {
			panic(nil)
		}
		char := pieces[1][0]
		pieces = strings.SplitN(pieces[0], "-", 2)
		if len(pieces) != 2 {
			panic(nil)
		}
		low, err := strconv.Atoi(pieces[0])
		if err != nil {
			panic(err)
		}
		high, err := strconv.Atoi(pieces[1])
		if err != nil {
			panic(err)
		}

		// count := 0
		// for i := 0; i < len(password); i++ {
		// 	if password[i] == char {
		// 		count++
		// 	}
		// }
		if count := strings.Count(password, string(char)); count >= low && count <= high {
			valid_p1++
		}

		if (password[low-1] == char) != (password[high-1] == char) {
			valid_p2++
		}
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	fmt.Println(valid_p1)
	fmt.Println(valid_p2)
}
