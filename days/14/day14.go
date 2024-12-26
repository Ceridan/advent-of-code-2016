package main

import (
	"crypto/md5"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkTriplet(hash [16]byte) uint8 {
	s := fmt.Sprintf("%x", hash)
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+1] && s[i] == s[i+2] {
			return s[i]
		}
	}
	return 0
}

func checkQuintet(hash [16]byte, char uint8) bool {
	s := fmt.Sprintf("%x", hash)
	for i := 0; i < len(s)-4; i++ {
		found := true
		for j := i; j < i+5; j++ {
			if s[j] != char {
				found = false
				break
			}
		}
		if found {
			return true
		}
	}
	return false
}

func searchKeys(possible map[int]uint8, hash [16]byte, idx int) []int {
	mx := 0
	if idx-1000 > mx {
		mx = idx - 1000
	}
	var keys []int
	for i := mx; i < idx; i++ {
		if char, ok := possible[i]; ok && checkQuintet(hash, char) {
			keys = append(keys, i)
		}
	}
	return keys
}

func generateKeys(salt string, hashFn func(s string) [16]byte) int {
	possible := make(map[int]uint8)
	keys := 0
	i := 0
	for {
		s := salt + strconv.Itoa(i)
		hash := hashFn(s)

		char := checkTriplet(hash)
		if char > 0 {
			possible[i] = char
		}
		ks := searchKeys(possible, hash, i)

		for _, k := range ks {
			keys += 1
			if keys == 64 {
				return k
			}
		}
		i++
	}
}

func Part1(salt string) int {
	hashFn := func(s string) [16]byte {
		return md5.Sum([]byte(s))
	}
	return generateKeys(salt, hashFn)
}

func Part2(salt string) int {
	hashFn := func(s string) [16]byte {
		hash := md5.Sum([]byte(s))
		for i := 1; i <= 2016; i++ {
			hash = md5.Sum([]byte(fmt.Sprintf("%x", hash)))
		}
		return hash
	}
	return generateKeys(salt, hashFn)
}

func main() {
	input, err := os.ReadFile("days/14/input.txt")
	if err != nil {
		panic(err)
	}
	salt := strings.Trim(string(input), "\n")

	fmt.Printf("Day 13, part 1: %v\n", Part1(salt))
	fmt.Printf("Day 13, part 2: %v\n", Part2(salt))
}
