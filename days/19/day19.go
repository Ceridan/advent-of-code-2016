package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://en.wikipedia.org/wiki/Josephus_problem
func solveJosephus2(n int) int {
	b := 1
	for b <= n {
		b = b << 1
	}
	l := n - (b >> 1)
	return 2*l + 1
}

func solveJosephusX(n int) int {
	b := 1
	for b <= n {
		b = b * 3
	}
	l := b / 3

	if n == l {
		return n
	}
	return n - l
}

func Part1(num int) int {
	return solveJosephus2(num)
}

func Part2(num int) int {
	return solveJosephusX(num)
}

func main() {
	input, err := os.ReadFile("days/19/input.txt")
	if err != nil {
		panic(err)
	}
	num, err := strconv.Atoi(strings.Trim(string(input), "\n"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("Day 19, part 1: %v\n", Part1(num))
	fmt.Printf("Day 19, part 2: %v\n", Part2(num))
}
