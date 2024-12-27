package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	expected := 5
	data := []string{
		"Disc #1 has 5 positions; at time=0, it is at position 4.",
		"Disc #2 has 2 positions; at time=0, it is at position 1.",
	}

	got := Part1(data)

	if got != expected {
		t.Errorf("Day 15, part1 = %d; want: %d", got, expected)
	}
}
