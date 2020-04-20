package main

import "testing"

func check(t *testing.T, got []int, idx0 int, idx1 int) {
	fatal := func() {
		t.Fatalf("expected %v,%v, got %v\n", idx0, idx1, got)
	}
	if len(got) != 2 {
		fatal()
	}
	if got[0] != idx0 || got[1] != idx1 {
		fatal()
	}
}

func TestSolve1(t *testing.T) {
	idx := twoSum([]int{2, 7, 11, 15}, 9)
	check(t, idx, 0, 1)
}
