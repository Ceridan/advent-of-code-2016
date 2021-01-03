package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	input := [][]int{
		{5, 10, 25},
		{3, 4, 5},
		{1, 2, 3},
	}
	expected := 1

	got := Part1(input)

	if got != expected {
		t.Errorf("Day 03, part1 = %d; want: %d", got, expected)
	}
}
