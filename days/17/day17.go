package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

const vaultSize = 4

var directions = map[string]*Point{"R": {1, 0}, "D": {0, 1}, "L": {-1, 0}, "U": {0, -1}}

type Point struct {
	x int
	y int
}

type QueueItem struct {
	pos  *Point
	path string
}

func (p *Point) add(other *Point) *Point {
	return &Point{x: p.x + other.x, y: p.y + other.y}
}

func (p *Point) equals(other *Point) bool {
	return p.x == other.x && p.y == other.y
}

func md5hash(s string) string {
	hash := md5.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}

func isOpen(r string) bool {
	switch r {
	case "b":
		fallthrough
	case "c":
		fallthrough
	case "d":
		fallthrough
	case "e":
		fallthrough
	case "f":
		return true
	default:
		return false
	}
}

func listOpenDirections(hash string) []string {
	var dirs []string
	if isOpen(hash[0:1]) {
		dirs = append(dirs, "U")
	}
	if isOpen(hash[1:2]) {
		dirs = append(dirs, "D")
	}
	if isOpen(hash[2:3]) {
		dirs = append(dirs, "L")
	}
	if isOpen(hash[3:4]) {
		dirs = append(dirs, "R")
	}
	return dirs
}

func shortestBfs(passcode string) string {
	vault := &Point{vaultSize - 1, vaultSize - 1}
	queue := []QueueItem{{pos: &Point{0, 0}, path: ""}}

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		if item.pos.equals(vault) {
			return item.path
		}

		if item.pos.x < 0 || item.pos.x >= vaultSize || item.pos.y < 0 || item.pos.y >= vaultSize {
			continue
		}

		hash := md5hash(passcode + item.path)
		dirs := listOpenDirections(hash)
		for _, dir := range dirs {
			q := QueueItem{pos: item.pos.add(directions[dir]), path: item.path + dir}
			queue = append(queue, q)
		}
	}
	return ""
}

func longestBfs(passcode string) int {
	vault := &Point{vaultSize - 1, vaultSize - 1}
	queue := []QueueItem{{pos: &Point{0, 0}, path: ""}}
	longest := 0

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		if item.pos.equals(vault) {
			if longest < len(item.path) {
				longest = len(item.path)
			}
			continue
		}

		if item.pos.x < 0 || item.pos.x >= vaultSize || item.pos.y < 0 || item.pos.y >= vaultSize {
			continue
		}

		hash := md5hash(passcode + item.path)
		dirs := listOpenDirections(hash)
		for _, dir := range dirs {
			q := QueueItem{pos: item.pos.add(directions[dir]), path: item.path + dir}
			queue = append(queue, q)
		}
	}
	return longest
}

func Part1(passcode string) string {
	return shortestBfs(passcode)
}

func Part2(passcode string) int {
	return longestBfs(passcode)
}

func main() {
	input, err := os.ReadFile("days/17/input.txt")
	if err != nil {
		panic(err)
	}
	passcode := strings.Trim(string(input), "\n")

	fmt.Printf("Day 17, part 1: %v\n", Part1(passcode))
	fmt.Printf("Day 17, part 2: %v\n", Part2(passcode))
}
