package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	name string
	arg1 interface{}
	arg2 interface{}
}

type VM struct {
	a int64
	b int64
	c int64
	d int64
}

func newVM() *VM {
	return &VM{a: 0, b: 0, c: 0, d: 0}
}

func (vm *VM) init(a int64, b int64, c int64, d int64) {
	vm.a = a
	vm.b = b
	vm.c = c
	vm.d = d
}

func (vm *VM) run(program []Instruction) {
	idx := 0
	for idx < len(program) {
		instr := program[idx]
		switch instr.name {
		case "cpy":
			cpy(vm, instr.arg1, instr.arg2)
		case "inc":
			inc(vm, instr.arg1.(string))
		case "dec":
			dec(vm, instr.arg1.(string))
		case "jnz":
			idx += jnz(vm, instr.arg1, instr.arg2.(int))
			continue
		}
		idx++
	}
}

func chooseRegister(vm *VM, r interface{}) *int64 {
	switch r {
	case "a":
		return &vm.a
	case "b":
		return &vm.b
	case "c":
		return &vm.c
	case "d":
		return &vm.d
	default:
		return nil
	}
}

func cpy(vm *VM, arg1 interface{}, arg2 interface{}) {
	var val int64
	r1 := chooseRegister(vm, arg1)
	if r1 != nil {
		val = *r1
	} else {
		val = arg1.(int64)
	}
	r2 := chooseRegister(vm, arg2)
	*r2 = val
}

func inc(vm *VM, arg string) {
	r := chooseRegister(vm, arg)
	*r += 1
}

func dec(vm *VM, arg string) {
	r := chooseRegister(vm, arg)
	*r -= 1
}

func jnz(vm *VM, arg1 interface{}, arg2 int) int {
	var val int64
	r := chooseRegister(vm, arg1)
	if r != nil {
		val = *r
	} else {
		val = arg1.(int64)
	}

	if val != 0 {
		return arg2
	}

	return 1
}

func parseInput(instructions []string) []Instruction {
	instr := make([]Instruction, len(instructions))
	for i, line := range instructions {
		split := strings.Split(line, " ")
		var arg1 interface{}
		arg1, err := strconv.ParseInt(split[1], 10, 64)
		if err != nil {
			arg1 = split[1]
		}
		var arg2 interface{}
		if len(split) > 2 {
			arg, err := strconv.ParseInt(split[2], 10, 32)
			if err != nil {
				arg2 = split[2]
			} else {
				arg2 = int(arg)
			}
		}
		instr[i] = Instruction{name: split[0], arg1: arg1, arg2: arg2}
	}
	return instr
}

func Part1(instructions []string) int64 {
	program := parseInput(instructions)
	vm := newVM()
	vm.run(program)
	return vm.a
}

func Part2(instructions []string) int64 {
	program := parseInput(instructions)
	vm := newVM()
	vm.init(0, 0, 1, 0)
	vm.run(program)
	return vm.a
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
