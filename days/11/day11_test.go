package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	arrangement := []string{
		"The first floor contains a hydrogen-compatible microchip and a lithium-compatible microchip.",
		"The second floor contains a hydrogen generator.",
		"The third floor contains a lithium generator.",
		"The fourth floor contains nothing relevant.",
	}
	expected := 11

	got := Part1(arrangement)

	if got != expected {
		t.Errorf("Day 11, part1 = %d; want: %d", got, expected)
	}
}
