package main

import (
	"fmt"
	"github.com/ceridan/advent-of-code-2016/pkg/assembunny"
	"os"
	"strings"
)

func Part1(instructions []string) int64 {
	program := assembunny.LoadProgram(instructions)
	vm := assembunny.NewVM()
	vm.Run(program)
	return vm.A
}

func Part2(instructions []string) int64 {
	program := assembunny.LoadProgram(instructions)
	vm := assembunny.NewVM()
	vm.Init(0, 0, 1, 0)
	vm.Run(program)
	return vm.A
}

func main() {
	input, err := os.ReadFile("days/12/input.txt")
	if err != nil {
		panic(err)
	}

	var instructions []string
	for _, in := range strings.Split(string(input), "\n") {
		if in == "" {
			continue
		}
		instructions = append(instructions, strings.Trim(in, " "))
	}

	fmt.Printf("Day 12, part 1: %v\n", Part1(instructions))
	fmt.Printf("Day 12, part 2: %v\n", Part2(instructions))
}
