package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	params := []struct {
		input    string
		expected int64
	}{
		{"ADVENT", 6},
		{"A(1x5)BC", 7},
		{"(3x3)XYZ", 9},
		{"A(2x2)BCD(2x2)EFG", 11},
		{"(6x1)(1x3)A", 6},
		{"X(8x2)(3x3)ABCY", 18},
	}

	for _, param := range params {
		got := Part1(param.input)

		if got != param.expected {
			t.Errorf("Day 09, part1 = %d; want: %d", got, param.expected)
		}
	}
}

func TestPart2(t *testing.T) {
	params := []struct {
		input    string
		expected int64
	}{
		{"(3x3)XYZ", 9},
		{"X(8x2)(3x3)ABCY", 20},
		{"(27x12)(20x12)(13x14)(7x10)(1x12)A", 241920},
		{"(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN", 445},
	}

	for _, param := range params {
		got := Part2(param.input)

		if got != param.expected {
			t.Errorf("Day 09, part2 = %d; want: %d", got, param.expected)
		}
	}
}
