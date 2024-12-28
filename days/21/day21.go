package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Letter struct {
	ch   string
	next *Letter
	prev *Letter
}

type LinkedList struct {
	head *Letter
	size int
}

func newLinkedList(password string) *LinkedList {
	head := &Letter{
		ch: password[0:1],
	}
	curr := head
	for i := 1; i < len(password); i++ {
		next := &Letter{
			ch:   password[i : i+1],
			prev: curr,
		}
		curr.next = next
		curr = next
	}
	curr.next = head
	head.prev = curr
	return &LinkedList{head: head, size: len(password)}
}

func (ll *LinkedList) swapPosition(x, y int) {
	var lx, ly *Letter
	pos := 0
	curr := ll.head
	for lx == nil || ly == nil {
		if pos == x {
			lx = curr
		}
		if pos == y {
			ly = curr
		}
		curr = curr.next
		pos++
	}
	lx.ch, ly.ch = ly.ch, lx.ch
}

func (ll *LinkedList) swapLetter(x, y string) {
	var lx, ly *Letter
	curr := ll.head
	for lx == nil || ly == nil {
		if curr.ch == x {
			lx = curr
		}
		if curr.ch == y {
			ly = curr
		}
		curr = curr.next
	}

	lx.ch, ly.ch = ly.ch, lx.ch
}

func (ll *LinkedList) rotateLeft(steps int) {
	for i := 0; i < steps; i++ {
		ll.head = ll.head.next
	}
}

func (ll *LinkedList) rotateRight(steps int) {
	for i := 0; i < steps; i++ {
		ll.head = ll.head.prev
	}
}

func (ll *LinkedList) rotateLetter(x string) {
	pos := 0
	curr := ll.head
	for {
		if curr.ch == x {
			break
		}
		curr = curr.next
		pos++
	}

	steps := pos + 1
	if pos >= 4 {
		steps += 1
	}
	ll.rotateRight(steps)
}

func (ll *LinkedList) reverse(x, y int) {
	if x > y {
		x, y = y, x
	}
	var lx, ly *Letter
	pos := 0
	curr := ll.head
	for lx == nil || ly == nil {
		if pos == x {
			lx = curr
		}
		if pos == y {
			ly = curr
		}
		curr = curr.next
		pos++
	}

	for x < y {
		lx.ch, ly.ch = ly.ch, lx.ch
		lx, ly = lx.next, ly.prev
		x++
		y--
	}
}
func (ll *LinkedList) move(x, y int) {
	pos := 0
	var lx *Letter
	curr := ll.head
	for lx == nil {
		if pos == x {
			lx = curr
		}
		curr = curr.next
		pos++
	}

	if lx == ll.head {
		ll.head = lx.next
	}
	lx.prev.next = lx.next
	lx.next.prev = lx.prev

	curr = ll.head
	for i := 0; i < y; i++ {
		curr = curr.next
	}

	lx.prev = curr.prev
	curr.prev.next = lx
	curr.prev = lx
	lx.next = curr

	if y == 0 {
		ll.head = lx
	}
}

func (ll *LinkedList) String() string {
	var sb strings.Builder
	curr := ll.head
	for i := 0; i < ll.size; i++ {
		sb.WriteString(curr.ch)
		curr = curr.next
	}
	return sb.String()
}

func (ll *LinkedList) Copy() *LinkedList {
	return newLinkedList(ll.String())
}

func forwardOperations(operations []string, password string) string {
	rd := regexp.MustCompile("(\\d+)")
	rl := regexp.MustCompile("letter (\\w)")

	ll := newLinkedList(password)
	for _, operation := range operations {
		if strings.HasPrefix(operation, "swap position") {
			matches := rd.FindAllStringSubmatch(operation, -1)
			x, _ := strconv.Atoi(matches[0][1])
			y, _ := strconv.Atoi(matches[1][1])
			ll.swapPosition(x, y)
		} else if strings.HasPrefix(operation, "swap letter") {
			matches := rl.FindAllStringSubmatch(operation, -1)
			x := matches[0][1]
			y := matches[1][1]
			ll.swapLetter(x, y)
		} else if strings.HasPrefix(operation, "rotate left") {
			matches := rd.FindAllStringSubmatch(operation, -1)
			x, _ := strconv.Atoi(matches[0][1])
			ll.rotateLeft(x)
		} else if strings.HasPrefix(operation, "rotate right") {
			matches := rd.FindAllStringSubmatch(operation, -1)
			x, _ := strconv.Atoi(matches[0][1])
			ll.rotateRight(x)
		} else if strings.HasPrefix(operation, "rotate based") {
			matches := rl.FindAllStringSubmatch(operation, -1)
			x := matches[0][1]
			ll.rotateLetter(x)
		} else if strings.HasPrefix(operation, "reverse") {
			matches := rd.FindAllStringSubmatch(operation, -1)
			x, _ := strconv.Atoi(matches[0][1])
			y, _ := strconv.Atoi(matches[1][1])
			ll.reverse(x, y)
		} else if strings.HasPrefix(operation, "move") {
			matches := rd.FindAllStringSubmatch(operation, -1)
			x, _ := strconv.Atoi(matches[0][1])
			y, _ := strconv.Atoi(matches[1][1])
			ll.move(x, y)
		}
	}

	return ll.String()
}

func backwardOperations(operations []string, password string) string {
	rd := regexp.MustCompile("(\\d+)")
	rl := regexp.MustCompile("letter (\\w)")

	ll := newLinkedList(password)
	for i := len(operations) - 1; i >= 0; i-- {
		operation := operations[i]
		if strings.HasPrefix(operation, "swap position") {
			matches := rd.FindAllStringSubmatch(operation, -1)
			x, _ := strconv.Atoi(matches[0][1])
			y, _ := strconv.Atoi(matches[1][1])
			ll.swapPosition(x, y)
		} else if strings.HasPrefix(operation, "swap letter") {
			matches := rl.FindAllStringSubmatch(operation, -1)
			x := matches[0][1]
			y := matches[1][1]
			ll.swapLetter(x, y)
		} else if strings.HasPrefix(operation, "rotate left") {
			matches := rd.FindAllStringSubmatch(operation, -1)
			x, _ := strconv.Atoi(matches[0][1])
			ll.rotateRight(x)
		} else if strings.HasPrefix(operation, "rotate right") {
			matches := rd.FindAllStringSubmatch(operation, -1)
			x, _ := strconv.Atoi(matches[0][1])
			ll.rotateLeft(x)
		} else if strings.HasPrefix(operation, "rotate based") {
			matches := rl.FindAllStringSubmatch(operation, -1)
			x := matches[0][1]
			pass := ll.String()
			for {
				ll.rotateLeft(1)
				tmp := ll.Copy()
				tmp.rotateLetter(x)
				if tmp.String() == pass {
					break
				}
			}
		} else if strings.HasPrefix(operation, "reverse") {
			matches := rd.FindAllStringSubmatch(operation, -1)
			x, _ := strconv.Atoi(matches[0][1])
			y, _ := strconv.Atoi(matches[1][1])
			ll.reverse(x, y)
		} else if strings.HasPrefix(operation, "move") {
			matches := rd.FindAllStringSubmatch(operation, -1)
			x, _ := strconv.Atoi(matches[0][1])
			y, _ := strconv.Atoi(matches[1][1])
			ll.move(y, x)
		}
	}

	return ll.String()
}

func Part1(operations []string, password string) string {
	return forwardOperations(operations, password)
}

func Part2(operations []string, password string) string {
	return backwardOperations(operations, password)
}

func main() {
	input, err := os.ReadFile("days/21/input.txt")
	if err != nil {
		panic(err)
	}
	operations := strings.Split(string(input), "\n")
	operations = operations[:len(operations)-1]

	fmt.Printf("Day 21, part 1: %v\n", Part1(operations, "abcdefgh"))
	fmt.Printf("Day 21, part 2: %v\n", Part2(operations, "fbgdceah"))
}
