package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	operations := []string{
		"rect 3x2",
		"rotate column x=1 by 1",
		"rotate row y=0 by 4",
		"rotate column x=1 by 1",
	}
	expected := 6

	got := Part1(operations, 7, 3)

	if got != expected {
		t.Errorf("Day 08, part1 = %d; want: %d", got, expected)
	}
}

func TestPart2(t *testing.T) {
	operations := []string{
		"rect 3x2",
		"rotate column x=1 by 1",
		"rotate row y=0 by 4",
		"rotate column x=1 by 1",
	}
	expected := `
--------------------------------------------------
.#..#.#
#.#....
.#.....
--------------------------------------------------
`

	got := Part2(operations, 7, 3)

	if got != expected {
		t.Errorf("Day 08, part2 = %s; want: %s", got, expected)
	}
}
