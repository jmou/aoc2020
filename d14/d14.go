package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func sum(mem map[uint64]uint64) uint64 {
	var sum uint64
	for _, value := range mem {
		sum += value
	}
	return sum
}

func Solve(r io.Reader) [2]uint64 {
	var maskMask, maskValue uint64
	var floatBits []int
	memP1 := make(map[uint64]uint64)
	memP2 := make(map[uint64]uint64)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		pieces := strings.Split(scanner.Text(), " = ")
		if pieces[0] == "mask" {
			maskMask, _ = strconv.ParseUint(strings.Map(func(r rune) rune {
				if r == 'X' {
					return '0'
				}
				return '1'
			}, pieces[1]), 2, 36)
			maskValue, _ = strconv.ParseUint(strings.Map(func(r rune) rune {
				if r == '1' {
					return '1'
				}
				return '0'
			}, pieces[1]), 2, 36)
			floatBits = nil
			for i := 0; i < 36; i++ {
				if (1<<i)&^maskMask != 0 {
					floatBits = append(floatBits, i)
				}
			}
		} else {
			addr, _ := strconv.ParseUint(pieces[0][4:len(pieces[0])-1], 10, 36)
			value, _ := strconv.ParseUint(pieces[1], 10, 36)
			// &^ is bitwise AND NOT. ^ also used for unary NOT
			memP1[addr] = value&^maskMask | maskValue

			// mask so 0 is unchanged
			addr &= maskMask &^ maskValue
			addr |= maskValue
			// floatBits powerset
			for set := 0; set < (1 << len(floatBits)); set++ {
				for i := 0; i < len(floatBits); i++ {
					floatMask := uint64(1) << floatBits[i]
					if (1<<i)&set == 0 {
						addr &^= floatMask
					} else {
						addr |= floatMask
					}
				}
				memP2[addr] = value
			}
		}
	}
	return [2]uint64{sum(memP1), sum(memP2)}
}

func main() {
	s := Solve(os.Stdin)
	fmt.Println(s[0])
	fmt.Println(s[1])
}
