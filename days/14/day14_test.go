package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	expected := 22728

	got := Part1("abc")

	if got != expected {
		t.Errorf("Day 14, part1 = %d; want: %d", got, expected)
	}
}

func TestPart2(t *testing.T) {
	expected := 22551

	got := Part2("abc")

	if got != expected {
		t.Errorf("Day 14, part2 = %d; want: %d", got, expected)
	}
}
