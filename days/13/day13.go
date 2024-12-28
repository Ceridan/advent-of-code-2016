package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var directions = []*Point{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

type Point struct {
	x int
	y int
}

func (p *Point) add(other *Point) *Point {
	return &Point{x: p.x + other.x, y: p.y + other.y}
}

func (p *Point) equals(other *Point) bool {
	return p.x == other.x && p.y == other.y
}

type QueueItem struct {
	p     *Point
	moves int
}

func isOpenCell(p *Point, num int) bool {
	n := p.x*p.x + 3*p.x + 2*p.x*p.y + p.y + p.y*p.y + num

	ones := 0
	for n > 0 {
		ones += n % 2
		n /= 2
	}

	return ones%2 == 0
}

func targetBfs(number int, start, target *Point) int {
	queue := []QueueItem{{p: start, moves: 0}}
	visited := map[Point]bool{}
	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		if item.p.equals(target) {
			return item.moves
		}

		if item.p.x < 0 || item.p.y < 0 {
			continue
		}
		if _, ok := visited[*item.p]; ok {
			continue
		}
		visited[*item.p] = true

		if !isOpenCell(item.p, number) {
			continue
		}

		for _, d := range directions {
			queue = append(queue, QueueItem{item.p.add(d), item.moves + 1})
		}
	}
	return -1
}

func stepsBfs(number int, start *Point, steps int) int {
	queue := []QueueItem{{p: start, moves: 0}}
	visited := map[Point]bool{}
	locations := 0
	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		if item.p.x < 0 || item.p.y < 0 || item.moves > steps {
			continue
		}
		if _, ok := visited[*item.p]; ok {
			continue
		}
		visited[*item.p] = true

		if !isOpenCell(item.p, number) {
			continue
		}

		locations += 1

		for _, d := range directions {
			queue = append(queue, QueueItem{item.p.add(d), item.moves + 1})
		}
	}
	return locations
}

func Part1(number int, start *Point, target *Point) int {
	return targetBfs(number, start, target)
}

func Part2(number int, start *Point, steps int) int {
	return stepsBfs(number, start, steps)
}

func main() {
	input, err := os.ReadFile("days/13/input.txt")
	if err != nil {
		panic(err)
	}
	str := strings.Trim(string(input), "\n")
	number, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Day 13, part 1: %v\n", Part1(number, &Point{1, 1}, &Point{31, 39}))
	fmt.Printf("Day 13, part 2: %v\n", Part2(number, &Point{1, 1}, 50))
}
