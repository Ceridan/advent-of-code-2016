package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	expected := "01100"
	state := "10000"
	diskSize := 20

	got := Part1(state, diskSize)

	if got != expected {
		t.Errorf("Day 16, part1 = %s; want: %s", got, expected)
	}
}
