package main

import (
	"fmt"
	"math"
	"os"
	"strings"
	"unicode"
)

var directions = []*Point{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

type Point struct {
	x int
	y int
}

func (p *Point) add(other *Point) *Point {
	return &Point{x: p.x + other.x, y: p.y + other.y}
}

func parseInput(data []string) ([][]rune, []*Point) {
	pointMap := make(map[int]*Point)
	grid := make([][]rune, len(data))
	for y := 0; y < len(data); y++ {
		grid[y] = []rune(data[y])
		for x := 0; x < len(grid[y]); x++ {
			if unicode.IsDigit(grid[y][x]) {
				pointMap[int(grid[y][x]-'0')] = &Point{x: x, y: y}
			}
		}
	}

	points := make([]*Point, len(pointMap))
	for r, p := range pointMap {
		points[r] = p
	}
	return grid, points
}

func bfs(grid [][]rune, start *Point, points []*Point) map[int]int {
	pointCost := make(map[int]int)
	queue := []struct {
		pos   *Point
		moves int
	}{{pos: start, moves: 0}}
	visited := make(map[Point]bool)

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]
		x, y := item.pos.x, item.pos.y
		if grid[y][x] == '#' {
			continue
		}
		if _, ok := visited[*item.pos]; ok {
			continue
		}
		visited[*item.pos] = true

		if unicode.IsDigit(grid[y][x]) {
			pointCost[int(grid[y][x]-'0')] = item.moves
			if len(pointCost) == len(points) {
				return pointCost
			}
		}

		for _, dir := range directions {
			queue = append(queue, struct {
				pos   *Point
				moves int
			}{pos: item.pos.add(dir), moves: item.moves + 1})
		}
	}
	return pointCost
}

func calculateTotalMoves(costs map[int]map[int]int, mask int, idx int, n int, returnIdx int, cache [][]int) int {
	if mask == (1<<n)-1 {
		if returnIdx >= 0 {
			return costs[idx][returnIdx]
		}
		return 0
	}
	if cache[idx][mask] != -1 {
		return cache[idx][mask]
	}

	moves := math.MaxInt
	for i := 0; i < n; i++ {
		if mask&(1<<i) == 0 {
			mv := costs[idx][i] + calculateTotalMoves(costs, mask|(1<<i), i, n, returnIdx, cache)
			if mv < moves {
				moves = mv
			}
		}
	}
	cache[idx][mask] = moves
	return moves
}

// https://en.wikipedia.org/wiki/Travelling_salesman_problem
func tsp(costs map[int]map[int]int, startIdx int, returnIdx int) int {
	n := len(costs)
	cache := make([][]int, n)
	for i := 0; i < n; i++ {
		cache[i] = make([]int, 1<<n)
		for j := 0; j < len(cache[i]); j++ {
			cache[i][j] = -1
		}
	}
	return calculateTotalMoves(costs, 1<<startIdx, startIdx, n, returnIdx, cache)
}

func calculateCosts(grid [][]rune, points []*Point) map[int]map[int]int {
	costs := make(map[int]map[int]int, len(points))
	for i, p := range points {
		costs[i] = bfs(grid, p, points)
	}
	return costs
}

func Part1(data []string) int {
	grid, points := parseInput(data)
	costs := calculateCosts(grid, points)
	return tsp(costs, 0, -1)
}

func Part2(data []string) int {
	grid, points := parseInput(data)
	costs := calculateCosts(grid, points)
	return tsp(costs, 0, 0)
}

func main() {
	input, err := os.ReadFile("days/24/input.txt")
	if err != nil {
		panic(err)
	}
	data := strings.Split(string(input), "\n")
	data = data[:len(data)-1]

	fmt.Printf("Day 24, part 1: %v\n", Part1(data))
	fmt.Printf("Day 24, part 2: %v\n", Part2(data))
}
