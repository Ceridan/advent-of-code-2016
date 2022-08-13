package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	instructions := []string{
		"value 5 goes to bot 2",
		"bot 2 gives low to bot 1 and high to bot 0",
		"value 3 goes to bot 1",
		"bot 1 gives low to output 1 and high to bot 0",
		"bot 0 gives low to output 2 and high to output 0",
		"value 2 goes to bot 2",
	}
	expected := 2

	got := Part1(instructions, 2, 5)

	if got != expected {
		t.Errorf("Day 10, part1 = %d; want: %d", got, expected)
	}
}

func TestPart2(t *testing.T) {
	instructions := []string{
		"value 5 goes to bot 2",
		"bot 2 gives low to bot 1 and high to bot 0",
		"value 3 goes to bot 1",
		"bot 1 gives low to output 1 and high to bot 0",
		"bot 0 gives low to output 2 and high to output 0",
		"value 2 goes to bot 2",
	}
	expected := 30

	got := Part2(instructions)

	if got != expected {
		t.Errorf("Day 10, part2 = %d; want: %d", got, expected)
	}
}
