package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	expected := int64(3)
	blacklist := []string{
		"5-8",
		"0-2",
		"4-7",
	}
	maxAddress := int64(9)

	got := Part1(blacklist, maxAddress)

	if got != expected {
		t.Errorf("Day 20, part1 = %d; want: %d", got, expected)
	}
}

func TestPart2(t *testing.T) {
	expected := int64(2)
	blacklist := []string{
		"5-8",
		"0-2",
		"4-7",
	}
	maxAddress := int64(9)

	got := Part2(blacklist, maxAddress)

	if got != expected {
		t.Errorf("Day 20, part2 = %d; want: %d", got, expected)
	}
}
