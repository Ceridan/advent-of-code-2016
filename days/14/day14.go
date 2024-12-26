package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func md5hash(s string) string {
	hash := md5.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}

func checkTriplet(hash string) uint8 {
	for i := 0; i < len(hash)-2; i++ {
		if hash[i] == hash[i+1] && hash[i] == hash[i+2] {
			return hash[i]
		}
	}
	return 0
}

func checkQuintet(hash string, char uint8) bool {
	for i := 0; i < len(hash)-4; i++ {
		found := true
		for j := i; j < i+5; j++ {
			if hash[j] != char {
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

func searchKeys(possible map[int]uint8, hash string, idx int) []int {
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

func generateKeys(salt string, hashFn func(s string) string) int {
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
	hashFn := func(s string) string {
		return md5hash(s)
	}
	return generateKeys(salt, hashFn)
}

func Part2(salt string) int {
	hashFn := func(s string) string {
		hash := s
		for i := 0; i <= 2016; i++ {
			hash = md5hash(hash)
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

	fmt.Printf("Day 14, part 1: %v\n", Part1(salt))
	fmt.Printf("Day 14, part 2: %v\n", Part2(salt))
}
