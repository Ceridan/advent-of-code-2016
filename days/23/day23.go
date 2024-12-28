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
			inc(vm, instr.arg1)
		case "dec":
			dec(vm, instr.arg1)
		case "tgl":
			tgl(vm, instr.arg1, program, idx)
		case "jnz":
			idx += jnz(vm, instr.arg1, instr.arg2)
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
	val := getIntValue(vm, arg1)
	r2 := chooseRegister(vm, arg2)
	if r2 != nil {
		*r2 = val
	}
}

func inc(vm *VM, arg interface{}) {
	r := chooseRegister(vm, arg)
	if r != nil {
		*r++
	}
}

func dec(vm *VM, arg interface{}) {
	r := chooseRegister(vm, arg)
	if r != nil {
		*r--
	}
}

func jnz(vm *VM, arg1 interface{}, arg2 interface{}) int {
	val := getIntValue(vm, arg1)
	loc := getIntValue(vm, arg2)

	if val != 0 {
		return int(loc)
	}
	return 1
}

func tgl(vm *VM, arg interface{}, program []Instruction, idx int) {
	val := getIntValue(vm, arg)
	tmpIdx := val + int64(idx)
	if tmpIdx < 0 || tmpIdx >= int64(len(program)) {
		return
	}
	newIdx := int(tmpIdx)

	switch program[newIdx].name {
	case "inc":
		program[newIdx].name = "dec"
	case "dec":
		program[newIdx].name = "inc"
	case "tgl":
		program[newIdx].name = "inc"
	case "jnz":
		program[newIdx].name = "cpy"
	case "cpy":
		program[newIdx].name = "jnz"
	}
}

func getIntValue(vm *VM, arg interface{}) int64 {
	r := chooseRegister(vm, arg)
	if r != nil {
		return *r
	}
	return arg.(int64)
}

func parseArg(arg string) interface{} {
	val, err := strconv.ParseInt(arg, 10, 64)
	if err != nil {
		return arg
	}
	return val
}

func parseInput(instructions []string) []Instruction {
	instr := make([]Instruction, len(instructions))
	for i, line := range instructions {
		split := strings.Split(line, " ")
		arg1 := parseArg(split[1])
		var arg2 interface{}
		if len(split) > 2 {
			arg2 = parseArg(split[2])
		}
		instr[i] = Instruction{name: split[0], arg1: arg1, arg2: arg2}
	}
	return instr
}

func Part1(instructions []string, a int) int64 {
	instr := parseInput(instructions)
	vm := newVM()
	vm.init(int64(a), 0, 0, 0)
	vm.run(instr)
	return vm.a
}

func Part2(instructions []string, a int) int64 {
	return 0
}

func main() {
	input, err := os.ReadFile("days/23/input.txt")
	if err != nil {
		panic(err)
	}
	instructions := strings.Split(string(input), "\n")
	instructions = instructions[:len(instructions)-1]

	fmt.Printf("Day 23, part 1: %v\n", Part1(instructions, 7))
	fmt.Printf("Day 23, part 2: %v\n", Part2(instructions, 7))
}
