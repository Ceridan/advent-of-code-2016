package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func Part1(triangles [][]int) int {
	return calculate(triangles)
}

func Part2(triangles [][]int) int {
	return calculate(triangles)
}

func calculate(triangles [][]int) int {
	possible := 0

	for _, sides := range triangles {
		if checkTriangle(sides) {
			possible += 1
		}
	}

	return possible
}

func checkTriangle(sides []int) bool {
	a := sides[0]
	b := sides[1]
	c := sides[2]

	switch {
	case a <= 0 || b <= 0 || c <= 0:
		return false
	case a >= b && a >= c:
		return a < b+c
	case b >= a && b >= c:
		return b < a+c
	default:
		return c < a+b
	}
}

func parseByRows(input string) [][]int {
	var data [][]int
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			break
		}

		var sides []int
		for _, s := range strings.Split(line, " ") {
			if s == "" {
				continue
			}

			side, _ := strconv.Atoi(s)
			sides = append(sides, side)
		}
		data = append(data, sides)
	}
	return data
}

func parseByColumns(input string) [][]int {
	var data [][]int
	triangles := [][]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0},
	}

	for i, line := range strings.Split(input, "\n") {
		if line == "" {
			break
		}

		j := 0
		for _, s := range strings.Split(line, " ") {
			if s == "" {
				continue
			}

			side, _ := strconv.Atoi(s)
			triangles[j][i%3] = side
			j += 1
		}

		if i%3 == 2 {
			for _, triangle := range triangles {
				tr := make([]int, len(triangle))
				copy(tr, triangle)
				data = append(data, tr)
			}
		}
	}
	return data
}

func main() {
	input, err := ioutil.ReadFile("days/03/input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Day 03, part 1: %v\n", Part1(parseByRows(string(input))))
	fmt.Printf("Day 03, part 2: %v\n", Part2(parseByColumns(string(input))))
}
