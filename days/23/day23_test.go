package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	expected := int64(3)
	instructions := []string{
		"cpy 2 a",
		"tgl a",
		"tgl a",
		"tgl a",
		"cpy 1 a",
		"dec a",
		"dec a",
	}
	a := 0

	got := Part1(instructions, a)

	if got != expected {
		t.Errorf("Day 23, part1 = %d; want: %d", got, expected)
	}
}
