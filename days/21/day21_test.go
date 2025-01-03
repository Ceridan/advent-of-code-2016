package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	expected := "decab"
	operations := []string{
		"swap position 4 with position 0",
		"swap letter d with letter b",
		"reverse positions 0 through 4",
		"rotate left 1 step",
		"move position 1 to position 4",
		"move position 3 to position 0",
		"rotate based on position of letter b",
		"rotate based on position of letter d",
	}
	password := "abcde"

	got := Part1(operations, password)

	if got != expected {
		t.Errorf("Day 21, part1 = %s; want: %s", got, expected)
	}
}

func TestPart2(t *testing.T) {
	expected := "abcde"
	operations := []string{
		"swap position 4 with position 0",
		"swap letter d with letter b",
		"reverse positions 0 through 4",
		"rotate left 1 step",
		"move position 1 to position 4",
		"move position 3 to position 0",
		"rotate based on position of letter b",
		"rotate based on position of letter d",
	}
	password := "decab"

	got := Part2(operations, password)

	if got != expected {
		t.Errorf("Day 21, part2 = %s; want: %s", got, expected)
	}
}
