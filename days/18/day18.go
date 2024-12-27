package main

import (
	"fmt"
	"os"
	"strings"
)

func generateNextRow(row []byte) ([]byte, int) {
	nextRow := make([]byte, len(row))
	safe := 0

	for i := 1; i < len(row)-1; i++ {
		if row[i-1]^row[i+1] == 1 {
			nextRow[i] = 1
		} else {
			nextRow[i] = 0
			safe++
		}
	}
	return nextRow, safe
}

func convertRow(row string) ([]byte, int) {
	bytes := make([]byte, len(row)+2)
	safe := 0
	for i := range row {
		if row[i] == '.' {
			bytes[i+1] = 0
			safe++
		} else {
			bytes[i+1] = 1
		}
	}
	return bytes, safe
}

func calculateSafeTiles(row string, depth int) int {
	bytes, safe := convertRow(row)
	for i := 1; i < depth; i++ {
		b, s := generateNextRow(bytes)
		bytes = b
		safe += s
	}
	return safe
}

func Part1(row string, depth int) int {
	return calculateSafeTiles(row, depth)
}

func Part2(row string, depth int) int {
	return calculateSafeTiles(row, depth)
}

func main() {
	input, err := os.ReadFile("days/18/input.txt")
	if err != nil {
		panic(err)
	}
	row := strings.Trim(string(input), "\n")

	fmt.Printf("Day 18, part 1: %v\n", Part1(row, 40))
	fmt.Printf("Day 18, part 2: %v\n", Part2(row, 400000))
}
