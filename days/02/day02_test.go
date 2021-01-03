package main

import (
	"strings"
	"testing"
)

func TestPart1(t *testing.T) {
	input := []string{
		"ULL",
		"RRDDD",
		"LURDL",
		"UUUUD",
	}
	expected := "1985"

	got := Part1(input)

	if got != expected {
		t.Errorf("Day 02, part1([%s]) = %s; want: %s", strings.Join(input, ", "), got, expected)
	}
}

func TestPart2(t *testing.T) {
	input := []string{
		"ULL",
		"RRDDD",
		"LURDL",
		"UUUUD",
	}
	expected := "5DB3"

	got := Part2(input)

	if got != expected {
		t.Errorf("Day 02, part2([%s]) = %s; want: %s", strings.Join(input, ", "), got, expected)
	}
}
