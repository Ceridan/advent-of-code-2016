package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// https://en.wikipedia.org/wiki/Josephus_problem
func solveJosephus2(n int) int {
	var l int
	b := 1
	for {
		if b > n {
			l = n - (b >> 1)
			break
		}
		b = b << 1
	}
	return 2*l + 1
}

func solveJosephusX(n int) int {
	p := math.Floor(math.Log(float64(n)) / math.Log(3.0))
	b := math.Pow(3.0, p)

	if n == int(b) {
		return n
	}

	return n - int(b)
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
