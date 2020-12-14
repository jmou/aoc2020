package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	var start int
	fmt.Scanln(&start)
	var line string
	fmt.Scanln(&line)
	pieces := strings.Split(line, ",")

	// no constant for MaxInt
	earliest := math.MaxInt32
	var p1 int
	rem, div := 0, 1
	for offset, bus := range pieces {
		if bus == "x" {
			continue
		}
		busId, err := strconv.Atoi(bus)
		if err != nil {
			panic(err)
		}

		arrival := ((start + busId - 1) / busId) * busId
		if arrival < earliest {
			earliest = arrival
			p1 = busId * (earliest - start)
		}

		// positive remainder
		target := ((-offset)%busId + busId) % busId
		// chinese remainder theorem sieve
		// I think this assumes bus IDs are coprime; a more proper
		// solution apparently would use the LCM.
		for rem%busId != target {
			rem += div
		}
		div *= busId
	}

	fmt.Println(p1)
	fmt.Println(rem)
}
