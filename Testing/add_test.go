package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	// test cases
	cases := []struct {
		a, b, want int
	}{
		{1, 2, 3},
		{2, 3, 5},
		{3, 4, 7},
		{-1, -1, 0},
		{0, 0, 0},
		{0, 1, 1},
	}
	// iterate over tetst cases
	for _, c := range cases {
		got := add(c.a, c.b)
		if got != c.want {
			t.Errorf("add(%d, %d) == %d, want %d", c.a, c.b, got, c.want)
		}
	}
}
