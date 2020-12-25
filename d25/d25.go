package main

import "fmt"

func transform(subject, loop int) int {
	value := 1
	for i := 0; i < loop; i++ {
		value *= subject
		value %= 20201227
	}
	return value
}

func crackLoop(subject, public int) int {
	value := 1
	loop := 0
	for value != public {
		value *= subject
		value %= 20201227
		loop++
	}
	return loop
}

func main() {
	publicKeys := make([]int, 2)
	for i := 0; i < 2; i++ {
		if _, err := fmt.Scanf("%d\n", &publicKeys[i]); err != nil {
			panic(err)
		}
	}
	loops := make([]int, 2)
	for i, public := range publicKeys {
		loops[i] = crackLoop(7, public)
	}
	keys := make([]int, 2)
	for i := 0; i < 2; i++ {
		keys[i] = transform(publicKeys[i], loops[1-i])
	}
	if keys[0] != keys[1] {
		panic("key mismatch")
	}
	fmt.Println(keys[0])
}
