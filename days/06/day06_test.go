package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	messages := []string{
		"eedadn",
		"drvtee",
		"eandsr",
		"raavrd",
		"atevrs",
		"tsrnev",
		"sdttsa",
		"rasrtv",
		"nssdts",
		"ntnada",
		"svetve",
		"tesnvt",
		"vntsnd",
		"vrdear",
		"dvrsen",
		"enarar",
	}
	expected := "easter"

	got := Part1(messages)

	if got != expected {
		t.Errorf("Day 06, part1 = %s; want: %s", got, expected)
	}
}
