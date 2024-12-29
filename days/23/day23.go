package main

import (
	"fmt"
	"github.com/ceridan/advent-of-code-2016/pkg/assembunny"
	"os"
	"strings"
)

func Part1(instructions []string, a int) int64 {
	program := assembunny.LoadProgram(instructions)
	vm := assembunny.NewVM()
	vm.Init(int64(a), 0, 0, 0)
	vm.Run(program)
	return vm.A
}

func Part2(instructions []string, a int) int64 {
	program := assembunny.LoadProgram(instructions)
	vm := assembunny.NewVM()
	vm.Init(int64(a), 0, 0, 0)
	patchFn := func(idx int) int {
		if idx == 4 {
			vm.A = vm.B * vm.D
			vm.C = 0
			vm.D = 0
			return 10
		}
		return idx
	}
	vm.ApplyPatch(patchFn)
	vm.Run(program)
	return vm.A
}

func main() {
	input, err := os.ReadFile("days/23/input.txt")
	if err != nil {
		panic(err)
	}
	instructions := strings.Split(string(input), "\n")
	instructions = instructions[:len(instructions)-1]

	fmt.Printf("Day 23, part 1: %v\n", Part1(instructions, 7))
	fmt.Printf("Day 23, part 2: %v\n", Part2(instructions, 12))
}
