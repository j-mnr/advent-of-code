package sixteen

import (
	"aoc/util"
	_ "embed"
	"fmt"
	"log/slog"
	"strings"
)

type mirror byte

const (
	empty  mirror = '.'
	diagL  mirror = '/'
	dRight mirror = '\\'
	vSplt  mirror = '|'
	hSplt  mirror = '-'
)

type mirrors []mirror

func (m mirrors) String() string { return string(m) }

var (
	// example1:
	//
	// >|<<<\....
	// |v-.\^....
	// .v...|->>>
	// .v...v^.|.
	// .v...v^...
	// .v...v^..\
	// .v../2\\..
	// <->-/vv|..
	// .|<<<2-|.\
	// .v//.|.v..
	//
	// ######....
	// .#...#....
	// .#...#####
	// .#...##...
	// .#...##...
	// .#...##...
	// .#..####..
	// ########..
	// .#######..
	// .#...#.#..
	//
	// Ultimately, in this example, 46 tiles become energized.
	example1 = strings.NewReader(`
.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....
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

// part1:
//
// The beam enters in the top-left corner from the left and heading to the
// right.
//
// Beams do not interact with other beams; a tile can have many beams passing
// through it at the same time. A tile is energized if that tile has at least
// one beam pass through it, reflect in it, or split in it.
func part1(input string) {
	var contraption []mirrors
	for i, line := range strings.Split(input, "\n") {
		contraption = append(contraption, make(mirrors, len(line)))
		for j, r := range line {
			switch mirror(r) {
			case empty, diagL, dRight, vSplt, hSplt:
				contraption[i][j] = mirror(r)
			default:
				panic("Impossible mirror " + string(r))
			}
		}
	}
	energized := make([][]byte, len(contraption))
	for i := range energized {
		energized[i] = make([]byte, len(contraption[i]))
		for k := range energized[i] {
			energized[i][k] = '.'
		}
	}

	type direction struct{ y, x int }
	var (
		up    = direction{y: -1, x: 0}
		down  = direction{y: 1, x: 0}
		left  = direction{y: 0, x: -1}
		right = direction{y: 0, x: 1}
	)

	type position struct{ row, col int }

	type coord struct {
		p position
		d direction
	}

	stack := []coord{{p: position{0, 0}, d: right}}
	for len(stack) != 0 {
		c := stack[0]
		stack = stack[1:]
		if 0 > c.p.row || c.p.row >= len(contraption) ||
			0 > c.p.col || c.p.col >= len(contraption[0]) {
			continue
		}
		switch contraption[c.p.row][c.p.col] {
		case diagL:
			var dir direction
			switch c.d {
			case up:
				dir = right
			case right:
				dir = up
			case left:
				dir = down
			case down:
				dir = left
			}
			stack = append(stack, coord{
				p: position{row: c.p.row + dir.y, col: c.p.col + dir.x},
				d: dir,
			})
			slog.Info("Diagonal Left '/'", "stack", stack)
		case dRight:
			var dir direction
			switch c.d {
			case up:
				dir = left
			case left:
				dir = up
			case right:
				dir = down
			case down:
				dir = right
			}
			stack = append(stack, coord{
				p: position{row: c.p.row + dir.y, col: c.p.col + dir.x},
				d: dir,
			})
			slog.Info("Diagonal Right '\\'", "stack", stack)
		case vSplt:
			if energized[c.p.row][c.p.col] == '#' {
				break
			}
			switch c.d {
			case left, right:
				stack = append(stack, coord{
					p: position{row: c.p.row + up.y, col: c.p.col + up.x},
					d: up,
				})
				stack = append(stack, coord{
					p: position{row: c.p.row + down.y, col: c.p.col + down.x},
					d: down,
				})
			case up, down:
				stack = append(stack, coord{
					p: position{row: c.p.row + c.d.y, col: c.p.col + c.d.x},
					d: c.d,
				})
			}
			slog.Info("Vertical Split", "stack", stack)
		case hSplt:
			if energized[c.p.row][c.p.col] == '#' {
				break
			}
			switch c.d {
			case up, down:
				stack = append(stack, coord{
					p: position{row: c.p.row + left.y, col: c.p.col + left.x},
					d: left,
				})
				stack = append(stack, coord{
					p: position{row: c.p.row + right.y, col: c.p.col + right.x},
					d: right,
				})
			case left, right:
				stack = append(stack, coord{
					p: position{row: c.p.row + c.d.y, col: c.p.col + c.d.x},
					d: c.d,
				})
			}
			slog.Info("Horziontal Split", "stack", stack)
		case empty:
			stack = append(stack, coord{
				p: position{row: c.p.row + c.d.y, col: c.p.col + c.d.x},
				d: c.d,
			})
			slog.Info("Empty", "stack", stack)
		}
		energized[c.p.row][c.p.col] = '#'
	}

	slog.Info("Contraption")
	for _, line := range contraption {
		fmt.Println(string(line))
	}

	var sum uint
	slog.Info("Energized")
	for _, line := range energized {
		fmt.Println(string(line))
		for _, b := range line {
			if b == '#' {
				sum++
			}
		}
	}
	slog.Error("Result", "sum", sum)
}

// part2:
func part2(input string) {
	panic("Unimplemented")
}
