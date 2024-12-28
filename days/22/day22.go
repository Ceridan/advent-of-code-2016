package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type Node struct {
	point     *Point
	size      int
	used      int
	available int
}

func parseInput(data []string) []Node {
	nodes := make([]Node, len(data)-2)
	r := regexp.MustCompile("/dev/grid/node-x(\\d+)-y(\\d+)")
	for i := 2; i < len(data); i++ {
		s := strings.ReplaceAll(data[i], " ", "[]")
		s = strings.ReplaceAll(s, "][", "")
		s = strings.ReplaceAll(s, "[]", " ")
		split := strings.Split(s, " ")
		matches := r.FindAllStringSubmatch(split[0], -1)
		x, _ := strconv.Atoi(matches[0][1])
		y, _ := strconv.Atoi(matches[0][2])
		point := &Point{x, y}
		size, _ := strconv.Atoi(split[1][:len(split[1])-1])
		used, _ := strconv.Atoi(split[2][:len(split[2])-1])
		available, _ := strconv.Atoi(split[3][:len(split[3])-1])
		nodes[i-2] = Node{point, size, used, available}
	}
	return nodes
}

func Part1(data []string) int {
	nodes := parseInput(data)
	viable := 0

	for i := 0; i < len(nodes)-1; i++ {
		for j := i + 1; j < len(nodes); j++ {
			if nodes[i].used > 0 && nodes[i].used <= nodes[j].available {
				viable++
			}
			if nodes[j].used > 0 && nodes[j].used <= nodes[i].available {
				viable++
			}
		}
	}
	return viable
}

func Part2(data []string) int {
	return 0
}

func main() {
	input, err := os.ReadFile("days/22/input.txt")
	if err != nil {
		panic(err)
	}
	data := strings.Split(string(input), "\n")
	data = data[:len(data)-1]

	fmt.Printf("Day 22, part 1: %v\n", Part1(data))
	fmt.Printf("Day 22, part 2: %v\n", Part2(data))
}
