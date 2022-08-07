package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type display struct {
	rows int
	cols int
	data [][]int
}

func Part1(operations []string, width int, height int) int {
	d := newDisplay(width, height)
	processCardInput(d, operations)
	return d.Checksum()
}

func Part2(operations []string, width int, height int) string {
	d := newDisplay(width, height)
	processCardInput(d, operations)
	return d.Print()
}

func processCardInput(d *display, operations []string) {
	rRect := regexp.MustCompile("rect (\\d+)x(\\d+)")
	rCol := regexp.MustCompile("rotate column x=(\\d+) by (\\d+)")
	rRow := regexp.MustCompile("rotate row y=(\\d+) by (\\d+)")

	for _, op := range operations {
		var matches []string
		matches = rRect.FindStringSubmatch(op)
		if len(matches) > 0 {
			cols, _ := strconv.Atoi(matches[1])
			rows, _ := strconv.Atoi(matches[2])
			d.Rect(cols, rows)
			continue
		}

		matches = rCol.FindStringSubmatch(op)
		if len(matches) > 0 {
			col, _ := strconv.Atoi(matches[1])
			rot, _ := strconv.Atoi(matches[2])
			d.RotateByColumn(col, rot)
			continue
		}

		matches = rRow.FindStringSubmatch(op)
		if len(matches) > 0 {
			row, _ := strconv.Atoi(matches[1])
			rot, _ := strconv.Atoi(matches[2])
			d.RotateByRow(row, rot)
			continue
		}
	}
}

func newDisplay(columns int, rows int) *display {
	data := make([][]int, rows)
	for row := 0; row < rows; row++ {
		data[row] = make([]int, columns)
	}
	return &display{
		rows: rows,
		cols: columns,
		data: data,
	}
}

func (d *display) RotateByColumn(col int, rot int) {
	rot = rot % d.rows
	newCol := make([]int, d.rows)
	for row := 0; row < d.rows; row++ {
		newCol[(row+rot)%d.rows] = d.data[row][col]
	}
	for row := 0; row < d.rows; row++ {
		d.data[row][col] = newCol[row]
	}
}

func (d *display) RotateByRow(row int, rot int) {
	rot = rot % d.cols
	newRow := make([]int, d.cols)
	for col := 0; col < d.cols; col++ {
		newRow[(col+rot)%d.cols] = d.data[row][col]
	}
	d.data[row] = newRow
}

func (d *display) Rect(cols int, rows int) {
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			d.data[row][col] = 1
		}
	}
}

func (d *display) Checksum() int {
	lit := 0
	for row := 0; row < d.rows; row++ {
		for col := 0; col < d.cols; col++ {
			lit += d.data[row][col]
		}
	}
	return lit
}

func (d *display) Print() string {
	var sb strings.Builder
	sb.WriteString("\n--------------------------------------------------\n")
	for row := 0; row < d.rows; row++ {
		for col := 0; col < d.cols; col++ {
			if d.data[row][col] == 0 {
				sb.WriteRune('.')
			} else {
				sb.WriteRune('#')
			}
		}
		sb.WriteRune('\n')
	}
	sb.WriteString("--------------------------------------------------\n")
	return sb.String()
}

func main() {
	input, err := os.ReadFile("days/08/input.txt")
	if err != nil {
		panic(err)
	}

	var operations []string
	for _, in := range strings.Split(string(input), "\n") {
		if in == "" {
			continue
		}
		operations = append(operations, strings.Trim(in, " "))
	}

	fmt.Printf("Day 08, part 1: %v\n", Part1(operations, 50, 6))
	fmt.Printf("Day 08, part 2: %v\n", Part2(operations, 50, 6))
}
