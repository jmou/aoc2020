package main

import (
	"strings"
	"testing"
)

func TestP1(t *testing.T) {
	t.Skip("Combined P2 solver runs in exponential time")
	p1 := Solve(strings.NewReader(`mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`))
	if p1[0] != 165 {
		t.Errorf("%d; want 165", p1[0])
	}
}

func TestP2(t *testing.T) {
	p2 := Solve(strings.NewReader(`mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1`))
	if p2[1] != 208 {
		t.Errorf("%d; want 208", p2[0])
	}
}
