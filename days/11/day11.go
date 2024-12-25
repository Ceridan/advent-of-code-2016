package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

const numberOfFloors = 4

func Part1(arrangement []string) int {
	generators := parseInput(arrangement, regexp.MustCompile("a (\\w+) generator"))
	microchips := parseInput(arrangement, regexp.MustCompile("a (\\w+)-compatible microchip"))
	return calculateMoves(generators, microchips)
}

func Part2(arrangement []string) int {
	generators := parseInput(arrangement, regexp.MustCompile("a (\\w+) generator"))
	microchips := parseInput(arrangement, regexp.MustCompile("a (\\w+)-compatible microchip"))
	generators[0] += 2
	microchips[0] += 2
	return calculateMoves(generators, microchips)
}
func getFloorItems(generators []int, microchips []int) []int {
	items := make([]int, numberOfFloors)
	for i := 0; i < numberOfFloors; i++ {
		items[i] = generators[i] + microchips[i]
	}
	return items
}

// Very hacky way to solve it. We do not respect types of the items, just counts.
// It requires 2 * (n - 1) - 1 steps to move all items from one floor to the next floor.
// The tricky part is an example when we have only microchips on the first floor.
// My idea is to add a heuristic - add additional moves to reach first generator.
func calculateMoves(generators []int, microchips []int) int {
	moves := 0
	for i := 0; i < numberOfFloors; i++ {
		if generators[i] > 0 {
			moves += 2 * i
			break
		}
	}

	items := getFloorItems(generators, microchips)
	for i := 0; i < numberOfFloors-1; i++ {
		if items[i] == 1 {
			moves += 1
		} else {
			moves += 2*(items[i]-1) - 1
		}
		items[i+1] += items[i]
	}
	return moves
}

func parseInput(arrangement []string, r *regexp.Regexp) []int {
	floors := make([]int, numberOfFloors)
	for i := 0; i < numberOfFloors; i++ {
		matches := r.FindAllStringSubmatch(arrangement[i], -1)
		floors[i] += len(matches)
	}
	return floors
}

func main() {
	input, err := os.ReadFile("days/11/input.txt")
	if err != nil {
		panic(err)
	}

	var arrangement []string
	for _, in := range strings.Split(string(input), "\n") {
		if in == "" {
			continue
		}
		arrangement = append(arrangement, strings.Trim(in, " "))
	}

	fmt.Printf("Day 11, part 1: %v\n", Part1(arrangement))
	fmt.Printf("Day 11, part 2: %v\n", Part2(arrangement))
}
