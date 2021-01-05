package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"strings"
)

func Part1(doorId string) string {
	password := make([]string, 0, 8)
	i := 0

	for len(password) < cap(password) {
		current := fmt.Sprintf("%s%d", doorId, i)
		hash := md5.Sum([]byte(current))

		if hash[0] == 0x00 && hash[1] == 0x00 && hash[2] <= 0x0f {
			password = append(password, fmt.Sprintf("%x", hash[2]))
		}

		i += 1
	}

	return strings.Join(password, "")
}

func Part2(doorId string) string {
	return ""
}

func main() {
	input, err := ioutil.ReadFile("days/05/input.txt")
	if err != nil {
		panic(err)
	}

	doorId := strings.Trim(string(input), "\n")

	fmt.Printf("Day 05, part 1: %v\n", Part1(doorId))
	fmt.Printf("Day 05, part 2: %v\n", Part2(doorId))
}
