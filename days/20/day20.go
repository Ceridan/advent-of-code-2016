package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct{ left, right int64 }

func parseInput(blacklist []string) ([]Range, error) {
	ranges := make([]Range, len(blacklist))
	for i, b := range blacklist {
		split := strings.Split(b, "-")
		l, err := strconv.Atoi(split[0])
		if err != nil {
			return nil, err
		}
		r, err := strconv.Atoi(split[1])
		if err != nil {
			return nil, err
		}
		ranges[i] = Range{int64(l), int64(r)}
	}
	return ranges, nil
}

func mergeBlacklist(bl []Range) []Range {
	sortFn := func(i, j int) bool {
		if bl[i].left == bl[j].left {
			return bl[i].right < bl[j].right
		}
		return bl[i].left < bl[j].left
	}
	sort.Slice(bl, sortFn)

	var merged []Range
	prev := bl[0]
	for i := 1; i < len(bl); i++ {
		if bl[i].left > prev.right+1 {
			merged = append(merged, prev)
			prev = bl[i]
			continue
		}

		if bl[i].right <= prev.right {
			continue
		}

		prev.right = bl[i].right
	}
	merged = append(merged, prev)
	return merged
}

func Part1(blacklist []string, _ int64) int64 {
	bl, err := parseInput(blacklist)
	if err != nil {
		panic(err)
	}
	merged := mergeBlacklist(bl)
	return merged[0].right + 1
}

func Part2(blacklist []string, maxAddress int64) int64 {
	bl, err := parseInput(blacklist)
	if err != nil {
		panic(err)
	}
	merged := mergeBlacklist(bl)
	allowed := merged[0].left
	for i := 1; i < len(merged); i++ {
		allowed += merged[i].left - merged[i-1].right - 1

	}
	allowed += maxAddress - merged[len(merged)-1].right
	return allowed
}

func main() {
	input, err := os.ReadFile("days/20/input.txt")
	if err != nil {
		panic(err)
	}
	blacklist := strings.Split(string(input), "\n")
	blacklist = blacklist[:len(blacklist)-1]

	fmt.Printf("Day 20, part 1: %v\n", Part1(blacklist, 4294967295))
	fmt.Printf("Day 20, part 2: %v\n", Part2(blacklist, 4294967295))
}
