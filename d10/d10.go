package main

import (
	"fmt"
	"sort"
)

func main() {
	var input []int
	var d int
	for {
		if _, err := fmt.Scanln(&d); err != nil {
			break
		}
		input = append(input, d)
	}
	sort.Ints(input)

	diff1, diff3 := 0, 1
	dp := make([]int, input[len(input)-1]+1)
	dp[0] = 1
	for i := 0; i < len(input); i++ {
		diff := input[i]
		if i > 0 {
			diff -= input[i-1]
		}
		if diff == 1 {
			diff1++
		} else if diff == 3 {
			diff3++
		}

		j := input[i]
		dp[j] = dp[j-1]
		if j >= 2 {
			dp[j] += dp[j-2]
		}
		if j >= 3 {
			dp[j] += dp[j-3]
		}
	}

	fmt.Println(diff1 * diff3)
	fmt.Println(dp[len(dp)-1])
}
