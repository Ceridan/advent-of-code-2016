package main

import (
	"container/list"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type outputInstruction struct {
	outputType string
	id         int
}

type bot struct {
	id              int
	low             int
	high            int
	lowInstruction  outputInstruction
	highInstruction outputInstruction
}

type botFactory map[int]*bot

func Part1(instructions []string, targetChipLow int, targetChipHigh int) int {
	bots := buildBots(instructions)
	queue := list.New()
	initQueue(queue, instructions, bots)

	for queue.Len() > 0 {
		id := queue.Remove(queue.Front()).(int)
		b := bots[id]

		if b.low == targetChipLow && b.high == targetChipHigh {
			return b.id
		}

		if b.lowInstruction.outputType == "bot" {
			bots[b.lowInstruction.id].SetValue(b.low)
			if bots[b.lowInstruction.id].IsReady() {
				queue.PushBack(b.lowInstruction.id)
			}
		}
		if b.highInstruction.outputType == "bot" {
			bots[b.highInstruction.id].SetValue(b.high)
			if bots[b.highInstruction.id].IsReady() {
				queue.PushBack(b.highInstruction.id)
			}
		}
	}

	return -1
}

func Part2(instructions []string) int {
	bots := buildBots(instructions)
	queue := list.New()
	initQueue(queue, instructions, bots)
	output := make(map[int]int)

	for queue.Len() > 0 {
		id := queue.Remove(queue.Front()).(int)
		b := bots[id]

		if b.lowInstruction.outputType == "bot" {
			bots[b.lowInstruction.id].SetValue(b.low)
			if bots[b.lowInstruction.id].IsReady() {
				queue.PushBack(b.lowInstruction.id)
			}
		} else {
			output[b.lowInstruction.id] = b.low
		}

		if b.highInstruction.outputType == "bot" {
			bots[b.highInstruction.id].SetValue(b.high)
			if bots[b.highInstruction.id].IsReady() {
				queue.PushBack(b.highInstruction.id)
			}
		} else {
			output[b.highInstruction.id] = b.high
		}
	}

	return output[0] * output[1] * output[2]
}

func initQueue(queue *list.List, instructions []string, bots botFactory) {
	r := regexp.MustCompile("value (\\d+) goes to bot (\\d+)")

	for _, instr := range instructions {
		matches := r.FindStringSubmatch(instr)
		if len(matches) == 0 {
			continue
		}

		val, _ := strconv.Atoi(matches[1])
		id, _ := strconv.Atoi(matches[2])
		bots[id].SetValue(val)
		if bots[id].IsReady() {
			queue.PushBack(id)
		}
	}
}

func buildBots(instructions []string) botFactory {
	r := regexp.MustCompile("bot (\\d+) gives low to (bot|output) (\\d+) and high to (bot|output) (\\d+)")
	bots := make(map[int]*bot)

	for _, instr := range instructions {
		matches := r.FindStringSubmatch(instr)
		if len(matches) == 0 {
			continue
		}

		low, _ := strconv.Atoi(matches[3])
		lowInstr := outputInstruction{outputType: matches[2], id: low}
		high, _ := strconv.Atoi(matches[5])
		highInstr := outputInstruction{outputType: matches[4], id: high}
		id, _ := strconv.Atoi(matches[1])
		bots[id] = &bot{id: id, lowInstruction: lowInstr, highInstruction: highInstr}
	}

	return bots
}

func (b *bot) IsReady() bool {
	return b.low > 0 && b.high > 0
}

func (b *bot) SetValue(value int) {
	if b.low == 0 {
		b.low = value
	} else if b.low > value {
		b.high = b.low
		b.low = value
	} else {
		b.high = value
	}
}

func main() {
	input, err := os.ReadFile("days/10/input.txt")
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

	fmt.Printf("Day 10, part 1: %v\n", Part1(instructions, 17, 61))
	fmt.Printf("Day 10, part 2: %v\n", Part2(instructions))
}
