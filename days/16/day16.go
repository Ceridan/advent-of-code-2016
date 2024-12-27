
package main

import (
	"fmt"
	"os"
	"strings"
)

func Part1(data []string) int {
	return 0
}

func Part2(data []string) int {
	return 0
}

func main() {
	input, err := os.ReadFile("days/16/input.txt")
	if err != nil {
		panic(err)
	}
	data := strings.Split(string(input), "\n")
	data = data[:len(data)-1]

	fmt.Printf("Day 16, part 1: %v\n", Part1(data))
	fmt.Printf("Day 16, part 2: %v\n", Part2(data))
}
