package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type disk struct {
	mod int
	pos int
}

func parseInput(data []string) []disk {
	disks := make([]disk, len(data))
	r := regexp.MustCompile("(\\d+)")
	for i, line := range data {
		matches := r.FindAllStringSubmatch(line, -1)
		mod, err := strconv.Atoi(matches[1][1])
		if err != nil {
			panic(err)
		}
		pos, err := strconv.Atoi(matches[3][1])
		if err != nil {
			panic(err)
		}
		d := disk{mod: mod, pos: pos}
		disks[i] = d
	}
	return disks
}

// https://en.wikipedia.org/wiki/Chinese_remainder_theorem
func solveReminders(disks []disk) int {
	m := make([]int, len(disks))
	M0 := 1
	for i, d := range disks {
		m[i] = d.mod
		M0 *= d.mod
	}

	M := make([]int, len(disks))
	for i, d := range disks {
		M[i] = M0 / d.mod
	}

	y := make([]int, len(disks))
	for i, d := range disks {
		mi := M[i] % d.mod
		// Rotate to the position it should be to have 0 position at the correct time.
		pos := (2*d.mod - d.pos - i - 1) % d.mod
		for j := 0; j < d.mod; j++ {
			if (mi*j)%d.mod == pos {
				y[i] = j
				break
			}
		}
	}

	x := 0
	for i := 0; i < len(disks); i++ {
		x = (x + M[i]*y[i]) % M0
	}
	return x
}

func Part1(data []string) int {
	disks := parseInput(data)
	return solveReminders(disks)
}

func Part2(data []string) int {
	disks := parseInput(data)
	disks = append(disks, disk{mod: 11, pos: 0})
	return solveReminders(disks)
}

func main() {
	input, err := os.ReadFile("days/15/input.txt")
	if err != nil {
		panic(err)
	}
	data := strings.Split(string(input), "\n")
	data = data[:len(data)-1]

	fmt.Printf("Day 15, part 1: %v\n", Part1(data))
	fmt.Printf("Day 15, part 2: %v\n", Part2(data))
}
