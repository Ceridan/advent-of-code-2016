package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	expected := 14
	grid := []string{
		"###########",
		"#0.1.....2#",
		"#.#######.#",
		"#4.......3#",
		"###########",
	}

	got := Part1(grid)

	if got != expected {
		t.Errorf("Day 24, part1 = %d; want: %d", got, expected)
	}
}
