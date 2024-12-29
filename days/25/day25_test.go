package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	expected := int64(1)
	instructions := []string{
		"out 0",
		"out 1",
		"jnz a -2",
	}

	got := Part1(instructions)

	if got != expected {
		t.Errorf("Day 25, part1 = %d; want: %d", got, expected)
	}
}
