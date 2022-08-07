package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Part1(compressed string) int64 {
	return calc(compressed, false)
}

func Part2(compressed string) int64 {
	return calc(compressed, true)
}

func calc(compressed string, expand bool) int64 {
	r := regexp.MustCompile("\\((\\d+)x(\\d+)\\)")
	var dlen int64
	var i = 0
	for i < len(compressed) {
		if compressed[i] == '(' {
			matches := r.FindStringSubmatch(compressed[i:])
			count, _ := strconv.Atoi(matches[1])
			repeat, _ := strconv.Atoi(matches[2])
			if expand {
				markerLen := i + len(matches[0])
				dlen += calc(compressed[markerLen:markerLen+count], expand) * int64(repeat)
			} else {
				dlen += int64(count * repeat)
			}
			i += count + len(matches[0])
			continue
		}
		dlen += 1
		i += 1
	}
	return dlen
}

func main() {
	input, err := os.ReadFile("days/09/input.txt")
	if err != nil {
		panic(err)
	}
	compressed := strings.Trim(string(input), "\n")

	fmt.Printf("Day 09, part 1: %v\n", Part1(compressed))
	fmt.Printf("Day 09, part 2: %v\n", Part2(compressed))
}
