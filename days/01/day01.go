package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type location = struct {
	x int
	y int
}

var directionsToCompass = map[int]string{
	0: "N",
	1: "E",
	2: "S",
	3: "W",
}

func Part1(route []string) int {
	dir, x, y := 0, 0, 0

	for _, r := range route {
		turn := 1
		if r[0] == 'L' {
			turn = -1
		}

		dir = (dir + turn + 4) % 4
		val, _ := strconv.Atoi(r[1:])

		switch c := directionsToCompass[dir]; c {
		case "N":
			y += val
		case "E":
			x += val
		case "S":
			y -= val
		case "W":
			x -= val
		}
	}

	return abs(x) + abs(y)
}

func Part2(route []string) int {
	dir, x, y := 0, 0, 0
	locs := map[location]bool{
		location{x: 0, y: 0}: true,
	}

	for _, r := range route {
		turn := 1
		if r[0] == 'L' {
			turn = -1
		}

		dir = (dir + turn + 4) % 4
		c := directionsToCompass[dir]
		val, _ := strconv.Atoi(r[1:])

		for ; val > 0; val-- {
			switch c {
			case "N":
				y += 1
			case "E":
				x += 1
			case "S":
				y -= 1
			case "W":
				x -= 1
			}

			l := location{x: x, y: y}
			_, ok := locs[l]

			if ok {
				return abs(x) + abs(y)
			} else {
				locs[l] = true
			}
		}
	}

	return 0
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	input, err := os.ReadFile("days/01/input.txt")
	if err != nil {
		panic(err)
	}

	var data []string
	for _, line := range strings.Split(strings.Trim(string(input), "\n"), ", ") {
		if line != "" {
			data = append(data, line)
		}
	}

	fmt.Printf("Day 01, part 1: %v\n", Part1(data))
	fmt.Printf("Day 01, part 2: %v\n", Part2(data))
}
