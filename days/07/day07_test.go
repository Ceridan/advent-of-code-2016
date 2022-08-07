package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	ips := []string{
		"abba[mnop]qrst",
		"abcd[bddb]xyyx",
		"aaaa[qwer]tyui",
		"ioxxoj[asdfgh]zxcvbn",
	}
	expected := 2

	got := Part1(ips)

	if got != expected {
		t.Errorf("Day 07, part1 = %d; want: %d", got, expected)
	}
}

func TestPart2(t *testing.T) {
	ips := []string{
		"aba[bab]xyz",
		"xyx[xyx]xyx",
		"aaa[kek]eke",
		"zazbz[bzb]cdb",
		"aba[bab]xyx[yxy]",
	}
	expected := 4

	got := Part2(ips)

	if got != expected {
		t.Errorf("Day 07, part2 = %d; want: %d", got, expected)
	}
}
