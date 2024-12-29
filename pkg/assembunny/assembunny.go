// For year 2016, days: 12, 23, 25

package assembunny

import (
	"strconv"
	"strings"
)

type Instruction struct {
	Name string
	Arg1 interface{}
	Arg2 interface{}
}

type VM struct {
	A       int64
	B       int64
	C       int64
	D       int64
	output  []int64
	patchFn func(int) int
}

func NewVM() *VM {
	return &VM{A: 0, B: 0, C: 0, D: 0, output: []int64{}}
}

func (vm *VM) Init(a int64, b int64, c int64, d int64) {
	vm.A = a
	vm.B = b
	vm.C = c
	vm.D = d
	vm.output = []int64{}
	vm.patchFn = nil
}

func (vm *VM) GetOutput() []int64 {
	return vm.output
}

func (vm *VM) Run(program []Instruction) {
	idx := 0
	for idx >= 0 && idx < len(program) {
		if vm.patchFn != nil {
			idx = vm.patchFn(idx)
			if idx < 0 || idx >= len(program) {
				break
			}
		}
		instr := program[idx]
		switch instr.Name {
		case "cpy":
			vm.cpy(instr.Arg1, instr.Arg2)
		case "inc":
			vm.inc(instr.Arg1)
		case "dec":
			vm.dec(instr.Arg1)
		case "out":
			vm.out(instr.Arg1)
		case "tgl":
			vm.tgl(instr.Arg1, program, idx)
		case "jnz":
			idx += vm.jnz(instr.Arg1, instr.Arg2)
			continue
		}
		idx++
	}
}

func (vm *VM) ApplyPatch(patch func(int) int) {
	vm.patchFn = patch
}

func (vm *VM) cpy(arg1 interface{}, arg2 interface{}) {
	val := vm.getIntValue(arg1)
	r2 := vm.chooseRegister(arg2)
	if r2 != nil {
		*r2 = val
	}
}

func (vm *VM) inc(arg interface{}) {
	r := vm.chooseRegister(arg)
	if r != nil {
		*r++
	}
}

func (vm *VM) dec(arg interface{}) {
	r := vm.chooseRegister(arg)
	if r != nil {
		*r--
	}
}

func (vm *VM) jnz(arg1 interface{}, arg2 interface{}) int {
	val := vm.getIntValue(arg1)
	loc := vm.getIntValue(arg2)

	if val != 0 {
		return int(loc)
	}
	return 1
}

func (vm *VM) tgl(arg interface{}, program []Instruction, idx int) {
	val := vm.getIntValue(arg)
	tmpIdx := val + int64(idx)
	if tmpIdx < 0 || tmpIdx >= int64(len(program)) {
		return
	}
	newIdx := int(tmpIdx)

	switch program[newIdx].Name {
	case "inc":
		program[newIdx].Name = "dec"
	case "dec":
		program[newIdx].Name = "inc"
	case "tgl":
		program[newIdx].Name = "inc"
	case "out":
		program[newIdx].Name = "inc"
	case "jnz":
		program[newIdx].Name = "cpy"
	case "cpy":
		program[newIdx].Name = "jnz"
	}
}

func (vm *VM) out(arg interface{}) {
	val := vm.getIntValue(arg)
	vm.output = append(vm.output, val)
}

func (vm *VM) chooseRegister(r interface{}) *int64 {
	switch r {
	case "a":
		return &vm.A
	case "b":
		return &vm.B
	case "c":
		return &vm.C
	case "d":
		return &vm.D
	default:
		return nil
	}
}

func (vm *VM) getIntValue(arg interface{}) int64 {
	r := vm.chooseRegister(arg)
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

func LoadProgram(data []string) []Instruction {
	instr := make([]Instruction, len(data))
	for i, line := range data {
		split := strings.Split(line, " ")
		arg1 := parseArg(split[1])
		var arg2 interface{}
		if len(split) > 2 {
			arg2 = parseArg(split[2])
		}
		instr[i] = Instruction{Name: split[0], Arg1: arg1, Arg2: arg2}
	}
	return instr
}
