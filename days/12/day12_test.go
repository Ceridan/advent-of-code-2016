package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	instructions := []string{
		"cpy 41 a",
		"inc a",
		"inc a",
		"dec a",
		"jnz a 2",
		"dec a",
	}
	expected := int64(42)

	got := Part1(instructions)

	if got != expected {
		t.Errorf("Day 12, part1 = %d; want: %d", got, expected)
	}
}
