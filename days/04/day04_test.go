package main

import (
	"strings"
	"testing"
)

func TestCalculateChecksum(t *testing.T) {
	fixtures := []struct {
		input    []string
		expected string
	}{
		{
			input:    []string{"aaaaa", "bbb", "z", "y", "x"},
			expected: "abxyz",
		},
		{
			input:    []string{"aaaaa", "bbbbbbbb", "z", "y", "x"},
			expected: "baxyz",
		},
		{
			input:    []string{"not", "a", "real", "room"},
			expected: "oarel",
		},
	}

	for _, fxt := range fixtures {
		in := strings.Join(fxt.input, "-")
		t.Run(in, func(t *testing.T) {
			got := calculateChecksum(fxt.input)
			if got != fxt.expected {
				t.Errorf("Day 01, calculateChecksum([%s]) = %s; want: %s", in, got, fxt.expected)
			}
		})
	}
}

func TestPart1(t *testing.T) {
	input := `
aaaaa-bbb-z-y-x-123[abxyz]
a-b-c-d-e-f-g-h-987[abcde]
a-b-c-d-e-f-g-h-987[bcdef]
not-a-real-room-404[oarel]
totally-real-room-200[decoy]
`
	expected := 1514

	rooms := parseInput(input)
	got := Part1(rooms)

	if got != expected {
		t.Errorf("Day 04, part1 = %d; want: %d", got, expected)
	}
}

func TestCalculateDecrypt(t *testing.T) {
	fixture := struct {
		encrypted string
		shift     int
	}{
		encrypted: "qzmt-zixmtkozy-ivhz",
		shift:     343,
	}
	expected := "very encrypted name"

	got := decrypt(fixture.encrypted, fixture.shift)

	if got != expected {
		t.Errorf("Day 04, decrypt(%s-%d) = %s; want: %s", fixture.encrypted, fixture.shift, got, expected)
	}
}
