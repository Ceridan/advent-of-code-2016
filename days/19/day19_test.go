package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	expected := 3
	num := 5

	got := Part1(num)

	if got != expected {
		t.Errorf("Day 19, part1 = %d; want: %d", got, expected)
	}
}

func TestPart2(t *testing.T) {
	expected := 2
	num := 5

	got := Part2(num)

	if got != expected {
		t.Errorf("Day 19, part2 = %d; want: %d", got, expected)
	}
}
