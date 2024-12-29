package main

import (
	"fmt"
	"github.com/ceridan/advent-of-code-2016/pkg/assembunny"
	"os"
	"strings"
)

func provideExitCondition(program []assembunny.Instruction, vm *assembunny.VM, valid *bool) func(int) int {
	return func(idx int) int {
		if idx > 0 && program[idx-1].Name == "out" {
			n := len(vm.GetOutput())
			if int64(n%2) == vm.GetOutput()[n-1] {
				*valid = false
				return -1
			}
			if n == 8 {
				*valid = true
				return -1
			}
		}
		return idx
	}
}

func Part1(instructions []string) int64 {
	program := assembunny.LoadProgram(instructions)
	vm := assembunny.NewVM()
	var valid bool
	patchFn := provideExitCondition(program, vm, &valid)
	var a int64 = 0
	for {
		vm.Init(a, 0, 0, 0)
		vm.ApplyPatch(patchFn)
		vm.Run(program)
		if valid {
			return a
		}
		a++
	}
}

func Part2(instructions []string) string {
	return "-"
}

func main() {
	input, err := os.ReadFile("days/25/input.txt")
	if err != nil {
		panic(err)
	}
	instructions := strings.Split(string(input), "\n")
	instructions = instructions[:len(instructions)-1]

	fmt.Printf("Day 25, part 1: %v\n", Part1(instructions))
	fmt.Printf("Day 25, part 2: %s\n", Part2(instructions))
}
