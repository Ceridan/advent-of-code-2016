package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	expected := 38
	row := ".^^.^.^^^^"
	depth := 10

	got := Part1(row, depth)

	if got != expected {
		t.Errorf("Day 18, part1 = %d; want: %d", got, expected)
	}
}
