package main

import (
	"fmt"
	"os"
	"strings"
)

func Part1(instructions []string) string {
	var keyToNextKey = map[rune]map[rune]rune{
		'1': {'U': '1', 'R': '2', 'D': '4', 'L': '1'},
		'2': {'U': '2', 'R': '3', 'D': '5', 'L': '1'},
		'3': {'U': '3', 'R': '3', 'D': '6', 'L': '2'},
		'4': {'U': '1', 'R': '5', 'D': '7', 'L': '4'},
		'5': {'U': '2', 'R': '6', 'D': '8', 'L': '4'},
		'6': {'U': '3', 'R': '6', 'D': '9', 'L': '5'},
		'7': {'U': '4', 'R': '8', 'D': '7', 'L': '7'},
		'8': {'U': '5', 'R': '9', 'D': '8', 'L': '7'},
		'9': {'U': '6', 'R': '9', 'D': '9', 'L': '8'},
	}

	res := calculate(instructions, keyToNextKey)
	return res
}

func Part2(instructions []string) string {
	var keyToNextKey = map[rune]map[rune]rune{
		'1': {'U': '1', 'R': '1', 'D': '3', 'L': '1'},
		'2': {'U': '2', 'R': '3', 'D': '6', 'L': '2'},
		'3': {'U': '1', 'R': '4', 'D': '7', 'L': '2'},
		'4': {'U': '4', 'R': '4', 'D': '8', 'L': '3'},
		'5': {'U': '5', 'R': '6', 'D': '5', 'L': '5'},
		'6': {'U': '2', 'R': '7', 'D': 'A', 'L': '5'},
		'7': {'U': '3', 'R': '8', 'D': 'B', 'L': '6'},
		'8': {'U': '4', 'R': '9', 'D': 'C', 'L': '7'},
		'9': {'U': '9', 'R': '9', 'D': '9', 'L': '8'},
		'A': {'U': '6', 'R': 'B', 'D': 'A', 'L': 'A'},
		'B': {'U': '7', 'R': 'C', 'D': 'D', 'L': 'A'},
		'C': {'U': '8', 'R': 'C', 'D': 'C', 'L': 'B'},
		'D': {'U': 'B', 'R': 'D', 'D': 'D', 'L': 'D'},
	}

	res := calculate(instructions, keyToNextKey)
	return res
}

func calculate(instructions []string, keyToNextKey map[rune]map[rune]rune) string {
	res := ""
	cur := '5'

	for _, instr := range instructions {
		for _, ch := range instr {
			cur = keyToNextKey[cur][ch]
		}

		res += string(cur)
	}

	return res
}

func main() {
	input, err := os.ReadFile("days/02/input.txt")
	if err != nil {
		panic(err)
	}

	var data []string
	for _, line := range strings.Split(string(input), "\n") {
		if line != "" {
			data = append(data, line)
		}
	}

	fmt.Printf("Day 02, part 1: %v\n", Part1(data))
	fmt.Printf("Day 02, part 2: %v\n", Part2(data))
}
