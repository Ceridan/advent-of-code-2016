package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var directions = []*Point{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

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

func (p *Point) add(other *Point) *Point {
	return &Point{x: p.x + other.x, y: p.y + other.y}
}

func (p *Point) equals(other *Point) bool {
	return p.x == other.x && p.y == other.y
}

func parseInput(data []string) []Node {
	nodes := make([]Node, len(data)-2)
	r := regexp.MustCompile("(\\d+)")
	for i := 2; i < len(data); i++ {
		matches := r.FindAllStringSubmatch(data[i], -1)
		x, _ := strconv.Atoi(matches[0][1])
		y, _ := strconv.Atoi(matches[1][1])
		size, _ := strconv.Atoi(matches[2][1])
		used, _ := strconv.Atoi(matches[3][1])
		available, _ := strconv.Atoi(matches[4][1])
		nodes[i-2] = Node{&Point{x, y}, size, used, available}
	}

	sort.Slice(nodes, func(i, j int) bool {
		if nodes[i].point.y == nodes[j].point.y {
			return nodes[i].point.x < nodes[j].point.x
		}
		return nodes[i].point.y < nodes[j].point.y
	})

	return nodes
}

func buildGrid(nodes []Node) [][]Node {
	hx, hy := nodes[len(nodes)-1].point.x, nodes[len(nodes)-1].point.y
	grid := make([][]Node, hy+1)
	for y := 0; y < hy+1; y++ {
		grid[y] = make([]Node, hx+1)
		for x := 0; x < hx+1; x++ {
			grid[y][x] = nodes[y*(hx+1)+x]
		}
	}
	return grid
}

func bfs(grid [][]Node, startNode *Node, targetNode *Node) int {
	queue := []struct {
		pos   *Point
		steps int
	}{{pos: startNode.point, steps: 0}}
	visited := make(map[Point]bool)

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]
		if item.pos.equals(targetNode.point) {
			return item.steps
		}
		if item.pos.y < 0 || item.pos.y >= len(grid) || item.pos.x < 0 || item.pos.x >= len(grid[0]) {
			continue
		}
		if grid[item.pos.y][item.pos.x].used > 2*startNode.size {
			continue
		}
		if _, ok := visited[*item.pos]; ok {
			continue
		}
		visited[*item.pos] = true

		for _, dir := range directions {
			queue = append(queue, struct {
				pos   *Point
				steps int
			}{pos: item.pos.add(dir), steps: item.steps + 1})
		}
	}
	return -1
}

func findEmptyNode(nodes []Node) *Node {
	for _, node := range nodes {
		if node.used == 0 {
			return &node
		}
	}
	return nil
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
	nodes := parseInput(data)
	emptyNode := findEmptyNode(nodes)
	if emptyNode == nil {
		fmt.Println("Empty node not found")
		os.Exit(1)
	}
	grid := buildGrid(nodes)
	leftFromTargetNode := grid[0][len(grid[0])-2]

	steps := bfs(grid, emptyNode, &leftFromTargetNode)
	if steps == -1 {
		fmt.Println("Impossible to move data from target node")
		os.Exit(1)
	}
	steps = steps + leftFromTargetNode.point.x*5 + 1
	return steps
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
