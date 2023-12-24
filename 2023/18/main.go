package eighteen

import (
	"aoc/util"
	"bytes"
	_ "embed"
	"fmt"
	"log/slog"
	"strconv"
	"strings"
)

var (
	// example1:
	example1 = strings.NewReader(`
U 6 (#70c710)
R 6 (#70c710)
D 11 (#0dc571)
L 2 (#5713f0)
D 2 (#d2c081)
R 2 (#59c680)
D 2 (#411b91)
L 5 (#8ceee2)
U 2 (#caa173)
L 1 (#1b58a2)
U 2 (#caa171)
R 2 (#7807d2)
U 3 (#a77fa3)
L 2 (#015232)
U 2 (#7a21e3)
`[1:])

	// example2:
	example2 = strings.NewReader(`
`[1:])

	//go:embed input.txt
	input string
)

func Run(part uint8, example bool) {
	data := util.PrepareInput(strings.NewReader(input))
	switch part {
	case 1:
		if example {
			data = util.PrepareInput(example1)
		}
		part1(data)
	case 2:
		if example {
			data = util.PrepareInput(example2)
		}
		part2(data)
	}
}

type direction byte

const (
	left  = 'L'
	right = 'R'
	up    = 'U'
	down  = 'D'
)

func (d direction) String() string {
	switch d {
	case left:
		return "left"
	case right:
		return "right"
	case up:
		return "up"
	case down:
		return "down"
	default:
		panic("Impossible direction " + string(d))
	}
}

// part1:
// XXX: 54632 too high
func part1(input string) {
	type order struct {
		dir  direction
		move int
	}

	newOrder := func(fields []string) order {
		var o order
		switch direction(fields[0][0]) {
		case left, right, up, down:
			o.dir = direction(fields[0][0])
		default:
			panic("Invalid direction " + string(fields[0]))
		}
		o.move = util.Must2(strconv.Atoi(fields[1]))
		return o
	}

	var orders []order
	position := struct{ y, x int }{}
	var minX, minY, maxX, maxY int
	for _, line := range strings.Split(input, "\n") {
		o := newOrder(strings.Fields(line))
		orders = append(orders, o)
		slog.Info("New world", "order", o)
		switch o.dir {
		case left:
			position.x -= o.move
			if position.x < minX {
				minX = position.x
			}
		case right:
			position.x += o.move
			if position.x > maxX {
				maxX = position.x
			}
		case up:
			position.y -= o.move
			if position.y < minY {
				minY = position.y
			}
		case down:
			position.y += o.move
			if position.y > maxY {
				maxY = position.y
			}
		}
	}

	digPlan := make([][]byte, maxY-minY+1)
	for i := range digPlan {
		digPlan[i] = make([]byte, maxX-minX+1)
		for k := range digPlan[i] {
			digPlan[i][k] = '.'
		}
	}

	// Impossible to go OOB for digPlan at this point.
	row, col := -minY, -minX
	for _, order := range orders {
		switch order.dir {
		case up:
			r := row
			for ; r >= row-order.move; r-- {
				digPlan[r][col] = '#'
			}
			row = r + 1
		case down:
			r := row
			for ; r <= row+order.move; r++ {
				digPlan[r][col] = '#'
			}
			row = r - 1
		case left:
			// slog.Info("Made it here", "row", row, "col", col, "move", order.move,
			// "condition GT", col - order.move)
			c := col
			for ; c >= col-order.move; c-- {
				digPlan[row][c] = '#'
			}
			col = c + 1
		case right:
			c := col
			for ; c <= col+order.move; c++ {
				digPlan[row][c] = '#'
			}
			col = c - 1
		}
	}

	for _, line := range digPlan {
		fmt.Println(string(line))
	}
	sum := 0
	marker := []byte{'#'}
	for _, line := range digPlan {
		sum += (bytes.LastIndex(line, marker) - bytes.Index(line, marker) + 1)
	}
	slog.Error("Result", "sum", sum)
}

// part2:
func part2(input string) {
	panic("Unimplemented")
}
