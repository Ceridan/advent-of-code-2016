package main

import (
	"fmt"
	"os"
	"strings"
)

func generateNextRow(row []byte) ([]byte, int) {
	nextRow := make([]byte, len(row))
	safe := 0

	for i := range nextRow {
		var left byte = 0
		if i > 0 && row[i-1] == 1 {
			left = 1
		}

		var right byte = 0
		if i < len(row)-1 && row[i+1] == 1 {
			right = 1
		}

		center := row[i]

		if left == 1 && center == 1 && right == 0 {
			nextRow[i] = 1
		} else if left == 0 && center == 1 && right == 1 {
			nextRow[i] = 1
		} else if left == 1 && center == 0 && right == 0 {
			nextRow[i] = 1
		} else if left == 0 && center == 0 && right == 1 {
			nextRow[i] = 1
		} else {
			nextRow[i] = 0
			safe++
		}
	}
	return nextRow, safe
}

func convertRow(row string) ([]byte, int) {
	bytes := make([]byte, len(row))
	safe := 0
	for i := range row {
		if row[i] == '.' {
			bytes[i] = 0
			safe++
		} else {
			bytes[i] = 1
		}
	}
	return bytes, safe
}

func Part1(row string, depth int) int {
	bytes, safe := convertRow(row)
	for i := 1; i < depth; i++ {
		b, s := generateNextRow(bytes)
		bytes = b
		safe += s
	}
	return safe
}

func Part2(row string, depth int) int {
	return 0
}

func main() {
	input, err := os.ReadFile("days/18/input.txt")
	if err != nil {
		panic(err)
	}
	row := strings.Trim(string(input), "\n")

	fmt.Printf("Day 18, part 1: %v\n", Part1(row, 40))
	fmt.Printf("Day 18, part 2: %v\n", Part2(row, 40))
}
