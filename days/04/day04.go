package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type room struct {
	names    []string
	sectorId int
	checksum string
}

func Part1(rooms []room) int {
	valid := 0

	for _, r := range rooms {
		checksum := calculateChecksum(r.names)
		if checksum == r.checksum {
			valid += r.sectorId
		}
	}

	return valid
}

func Part2(rooms []room) int {
	for _, r := range rooms {
		checksum := calculateChecksum(r.names)
		if checksum != r.checksum {
			continue
		}
		decrypted := decrypt(strings.Join(r.names, "-"), r.sectorId)
		if decrypted == "northpole object storage" {
			return r.sectorId
		}
	}
	return 0
}

func decrypt(name string, sectorId int) string {
	encrypted := []rune(name)
	decrypted := make([]rune, len(encrypted))

	for i, r := range encrypted {
		if r == '-' {
			decrypted[i] = ' '
		} else {
			decrypted[i] = rune(((int(r)-97)+sectorId)%26 + 97)
		}
	}

	return string(decrypted)
}

func calculateChecksum(names []string) string {
	nameToCount := make(map[rune]int)
	for _, w := range names {
		for _, r := range []rune(w) {
			_, ok := nameToCount[r]
			if ok {
				nameToCount[r] += 1
			} else {
				nameToCount[r] = 1
			}
		}
	}

	countToName := make(map[int][]string)
	countSet := make(map[int]bool)
	for k, v := range nameToCount {
		_, ok := countToName[v]
		if ok {
			countToName[v] = append(countToName[v], string(k))
		} else {
			countToName[v] = []string{string(k)}
		}
		countSet[v] = true
	}

	counts := make([]int, 0, len(countSet))
	for k := range countSet {
		counts = append(counts, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(counts)))

	checksum := make([]string, 0, 5)
	for _, v := range counts {
		sort.Strings(countToName[v])
		for _, ch := range countToName[v] {
			checksum = append(checksum, ch)

			if len(checksum) == 5 {
				return strings.Join(checksum, "")
			}
		}
	}
	return ""
}

func parseInput(input string) []room {
	var rooms []room
	re := regexp.MustCompile(`([a-z]+)|(\d)+`)

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		parsed := re.FindAll([]byte(line), -1)
		names := make([]string, len(parsed))
		for i, w := range parsed {
			names[i] = string(w)
		}

		sectorId, _ := strconv.Atoi(names[len(parsed)-2])

		r := room{
			names:    names[:len(names)-2],
			sectorId: sectorId,
			checksum: names[len(names)-1],
		}
		rooms = append(rooms, r)
	}
	return rooms
}

func main() {
	input, err := ioutil.ReadFile("days/04/input.txt")
	if err != nil {
		panic(err)
	}

	rooms := parseInput(string(input))

	fmt.Printf("Day 04, part 1: %v\n", Part1(rooms))
	fmt.Printf("Day 04, part 2: %v\n", Part2(rooms))
}
