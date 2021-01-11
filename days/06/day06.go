package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func Part1(messages []string) string {
	tmessages := transpose(messages)
	ecv := make([]rune, len(messages[0]), len(messages[0]))

	for i, msg := range tmessages {
		_, most := common_chars(msg)
		ecv[i] = most
	}

	return string(ecv)
}

func Part2(messages []string) string {
	tmessages := transpose(messages)
	ecv := make([]rune, len(messages[0]), len(messages[0]))

	for i, msg := range tmessages {
		less, _ := common_chars(msg)
		ecv[i] = less
	}

	return string(ecv)
}

func common_chars(message string) (less rune, most rune) {
	chars := make(map[rune]int)
	runes := []rune(message)

	for _, r := range runes {
		chars[r] += 1
	}

	lc, mc := len(message)+1, 0
	for k, v := range chars {
		if v < lc {
			less = k
			lc = v
		}
		if v > mc {
			most = k
			mc = v
		}
	}

	return less, most
}

func transpose(messages []string) []string {
	trunes := make([][]rune, len(messages[0]), len(messages[0]))

	for _, message := range messages {
		runes := []rune(message)
		for j := 0; j < len(messages[0]); j++ {
			trunes[j] = append(trunes[j], runes[j])
		}
	}

	tmessages := make([]string, len(messages[0]), len(messages[0]))
	for i, runes := range trunes {
		tmessages[i] = string(runes)
	}

	return tmessages
}

func main() {
	input, err := ioutil.ReadFile("days/06/input.txt")
	if err != nil {
		panic(err)
	}

	var messages []string
	for _, in := range strings.Split(string(input), "\n") {
		if in == "" {
			continue
		}
		messages = append(messages, strings.Trim(in, " "))
	}

	fmt.Printf("Day 06, part 1: %v\n", Part1(messages))
	fmt.Printf("Day 06, part 2: %v\n", Part2(messages))
}
