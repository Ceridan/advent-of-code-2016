
package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	expected := 0
	data := []string{}

	got := Part1(data)

	if got != expected {
		t.Errorf("Day 16, part1 = %d; want: %d", got, expected)
	}
}
