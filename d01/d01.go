package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("d01/input")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	numbers := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		numbers = append(numbers, n)
	}

outer1:
	for _, i := range numbers {
		for _, j := range numbers {
			if i+j == 2020 {
				fmt.Println(i * j)
				break outer1
			}
		}
	}

outer2:
	for _, i := range numbers {
		for _, j := range numbers {
			for _, k := range numbers {
				if i+j+k == 2020 {
					fmt.Println(i * j * k)
					break outer2
				}
			}
		}
	}
}
