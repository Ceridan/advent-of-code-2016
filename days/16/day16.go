package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func convertState(state string) []byte {
	var data []byte
	for i := 0; i < len(state); i++ {
		var b byte = 0
		if state[i:i+1] == "1" {
			b = 1
		}
		data = append(data, b)
	}
	return data
}

func generateData(data []byte, diskSize int) []byte {
	if len(data) >= diskSize {
		return data[:diskSize]
	}

	newData := make([]byte, 2*len(data)+1)
	copy(newData, data)
	newData[len(data)] = 0

	for i := 0; i < len(data); i++ {
		newData[i+1+len(data)] = data[len(data)-1-i] ^ 1
	}
	return generateData(newData, diskSize)
}

func calculateChecksum(data []byte) string {
	if len(data)%2 == 1 {
		var result strings.Builder
		for _, b := range data {
			result.WriteString(strconv.Itoa(int(b)))
		}
		return result.String()
	}

	newData := make([]byte, len(data)/2)
	for i := 0; i < len(data); i += 2 {
		newData[i/2] = data[i] ^ data[i+1] ^ 1
	}
	return calculateChecksum(newData)
}

func Part1(state string, diskSize int) string {
	data := generateData(convertState(state), diskSize)
	checksum := calculateChecksum(data)
	return checksum
}

func Part2(state string, diskSize int) string {
	return ""
}

func main() {
	input, err := os.ReadFile("days/16/input.txt")
	if err != nil {
		panic(err)
	}
	state := strings.Trim(string(input), "\n")

	fmt.Printf("Day 16, part 1: %v\n", Part1(state, 272))
	fmt.Printf("Day 16, part 2: %v\n", Part2(state, 272))
}
