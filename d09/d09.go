package main

import "fmt"

func main() {
	var input []int
	for {
		var d int
		if _, err := fmt.Scanf("%d\n", &d); err != nil {
			break
		}
		input = append(input, d)
	}

	var invalid int
	for i := 25; i < len(input); i++ {
		window := input[i-25 : i]
		num := input[i]
		valid := false
		for _, i := range window {
			// oops j should start after i
			for _, j := range window {
				if i != j && i+j == num {
					valid = true
					break
				}
			}
		}
		if !valid {
			fmt.Println(num)
			invalid = num
		}
	}

	lo, hi, sum := 0, 0, 0
	for hi < len(input) {
		if sum == invalid {
			// no min/max builtin
			min, max := input[lo], input[lo]
			for i := lo + 1; i < hi; i++ {
				if input[i] < min {
					min = input[i]
				}
				if input[i] > max {
					max = input[i]
				}
			}
			fmt.Println(min + max)
			break
		} else if sum < invalid {
			sum += input[hi]
			hi++
		} else {
			sum -= input[lo]
			lo++
		}
	}
}
