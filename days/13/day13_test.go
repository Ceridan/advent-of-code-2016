package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	expected := 11

	got := Part1(10, &Point{1, 1}, &Point{7, 4})

	if got != expected {
		t.Errorf("Day 13, part1 = %d; want: %d", got, expected)
	}
}
