package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	data := []struct {
		passcode string
		expected string
	}{
		{passcode: "ihgpwlah", expected: "DDRRRD"},
		{passcode: "kglvqrro", expected: "DDUDRLRRUDRD"},
		{passcode: "ulqzkmiv", expected: "DRURDRUDDLLDLUURRDULRLDUUDDDRR"},
	}

	for _, d := range data {
		got := Part1(d.passcode)
		if got != d.expected {
			t.Errorf("Day 17, part1 = %s; want: %s", got, d.expected)
		}
	}
}

func TestPart2(t *testing.T) {
	data := []struct {
		passcode string
		expected int
	}{
		{passcode: "ihgpwlah", expected: 370},
		{passcode: "kglvqrro", expected: 492},
		{passcode: "ulqzkmiv", expected: 830},
	}

	for _, d := range data {
		got := Part2(d.passcode)
		if got != d.expected {
			t.Errorf("Day 17, part2 = %d; want: %d", got, d.expected)
		}
	}
}
