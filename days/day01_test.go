package main

import (
	"strings"
	"testing"
)

type fixture struct {
	input    []string
	expected int
}

func TestPart1(t *testing.T) {
	fixtures := []fixture{
		{
			input:    []string{"R2", "L3"},
			expected: 5,
		},
		{
			input:    []string{"R2", "R2", "R2"},
			expected: 2,
		},
		{
			input:    []string{"R5", "L5", "R5", "R3"},
			expected: 12,
		},
	}

	for _, fxt := range fixtures {
		in := strings.Join(fxt.input, ", ")
		t.Run(in, func(t *testing.T) {
			got := Day01Part1(fxt.input)
			if got != fxt.expected {
				t.Errorf("Day 01, part1([%s]) = %d; want: %d", in, got, fxt.expected)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	fixtures := []fixture{
		{
			input:    []string{"R8", "R4", "R4", "R8"},
			expected: 4,
		},
	}

	for _, fxt := range fixtures {
		in := strings.Join(fxt.input, ", ")
		t.Run(in, func(t *testing.T) {
			got := Day01Part2(fxt.input)
			if got != fxt.expected {
				t.Errorf("Day 01, part2([%s]) = %d; want: %d", in, got, fxt.expected)
			}
		})
	}
}
