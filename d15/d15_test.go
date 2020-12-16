package main

import (
	"testing"
)

func TestP1(t *testing.T) {
	a := Solve([]int{0, 3, 6}, 2020)
	if a != 436 {
		t.Errorf("%d; want 436", a)
	}

	a = Solve([]int{1, 3, 2}, 2020)
	if a != 1 {
		t.Errorf("%d; want 1", a)
	}

	a = Solve([]int{2, 1, 3}, 2020)
	if a != 10 {
		t.Errorf("%d; want 10", a)
	}

	a = Solve([]int{1, 2, 3}, 2020)
	if a != 27 {
		t.Errorf("%d; want 27", a)
	}

	a = Solve([]int{2, 3, 1}, 2020)
	if a != 78 {
		t.Errorf("%d; want 78", a)
	}

	a = Solve([]int{3, 2, 1}, 2020)
	if a != 438 {
		t.Errorf("%d; want 438", a)
	}

	a = Solve([]int{3, 1, 2}, 2020)
	if a != 1836 {
		t.Errorf("%d; want 1836", a)
	}
}

func TestP2(t *testing.T) {
	a := Solve([]int{0, 3, 6}, 30000000)
	if a != 175594 {
		t.Errorf("%d; want 175594", a)
	}

	a = Solve([]int{1, 3, 2}, 30000000)
	if a != 2578 {
		t.Errorf("%d; want 2578", a)
	}

	a = Solve([]int{2, 1, 3}, 30000000)
	if a != 3544142 {
		t.Errorf("%d; want 3544142", a)
	}

	a = Solve([]int{1, 2, 3}, 30000000)
	if a != 261214 {
		t.Errorf("%d; want 261214", a)
	}

	a = Solve([]int{2, 3, 1}, 30000000)
	if a != 6895259 {
		t.Errorf("%d; want 6895259", a)
	}

	a = Solve([]int{3, 2, 1}, 30000000)
	if a != 18 {
		t.Errorf("%d; want 18", a)
	}

	a = Solve([]int{3, 1, 2}, 30000000)
	if a != 362 {
		t.Errorf("%d; want 362", a)
	}
}
