package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	doorId := "abc"
	expected := "18f47a30"

	got := Part1(doorId)

	if got != expected {
		t.Errorf("Day 05, part1(%s) = %s; want: %s", doorId, got, expected)
	}
}

func TestPart2(t *testing.T) {
	doorId := "abc"
	expected := "05ace8e3"

	got := Part2(doorId)

	if got != expected {
		t.Errorf("Day 05, part2(%s) = %s; want: %s", doorId, got, expected)
	}
}
