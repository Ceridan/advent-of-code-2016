package main

import (
	"crypto/md5"
	"fmt"
	"os"
	"strings"
)

func Part1(doorId string) string {
	password := make([]rune, 0, 8)
	i := 0

	for len(password) < cap(password) {
		current := fmt.Sprintf("%s%d", doorId, i)
		hash := md5.Sum([]byte(current))

		if hash[0] == 0x00 && hash[1] == 0x00 && hash[2] <= 0x0f {
			hexstring := []rune(fmt.Sprintf("%x", hash))
			password = append(password, hexstring[5])
		}

		i += 1
	}

	return string(password)
}

func Part2(doorId string) string {
	password := make([]rune, 8)
	i, placed := 0, 0

	for placed < 8 {
		current := fmt.Sprintf("%s%d", doorId, i)
		hash := md5.Sum([]byte(current))

		if hash[0] == 0x00 && hash[1] == 0x00 && hash[2] <= 0x0f {
			pos := int(hash[2])
			if pos < 8 && password[pos] == 0 {
				hexstring := []rune(fmt.Sprintf("%x", hash))
				password[pos] = hexstring[6]
				placed += 1
			}
		}

		i += 1
	}

	return string(password)
}

func main() {
	input, err := os.ReadFile("days/05/input.txt")
	if err != nil {
		panic(err)
	}

	doorId := strings.Trim(string(input), "\n")

	fmt.Printf("Day 05, part 1: %v\n", Part1(doorId))
	fmt.Printf("Day 05, part 2: %v\n", Part2(doorId))
}
