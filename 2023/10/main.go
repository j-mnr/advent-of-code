package ten

import (
	"aoc/util"
	_ "embed"
	"fmt"
	"log/slog"
	"strconv"
	"strings"
)

var (
	// example1:
	//
	//
	// Same as below:
	// 7-F7-
	// .FJ|7
	// SJLL7
	// |F--J
	// LJ.LJ
	example1 = strings.NewReader(`
.....
.S-7.
.|.|.
.L-J.
.....
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

var (
	above = coord{y: -1, x: 0}
	below = coord{y: 1, x: 0}
	left  = coord{y: 0, x: -1}
	right = coord{y: 0, x: 1}
)

type coord struct{ y, x int }

type pipe struct {
	// typ is the allowed connections to the pipe.
	//  | is a vertical pipe connecting north and south.
	//  - is a horizontal pipe connecting east and west.
	//  L is a 90-degree bend connecting north and east.
	//  J is a 90-degree bend connecting north and west.
	//  7 is a 90-degree bend connecting south and west.
	//  F is a 90-degree bend connecting south and east.
	//  S is the starting position of the animal; there is a pipe on this tile,
	//  but your sketch doesn't show what shape the pipe has.
	typ rune
	coord
}

func (p pipe) String() string {
	return "(" + strconv.Itoa(p.y) + "," + strconv.Itoa(p.x) + ") " +
		strconv.QuoteRune(p.typ)
}

// part1: Find the single giant loop starting at S. How many steps along the
// loop does it take to get from the starting position to the point *farthest*
// from the starting position?
func part1(input string) {
	diagram := strings.Split(input, "\n")
	var pipes []*pipe
	var depth uint
	for row, line := range diagram {
		for col, r := range line {
			if r != 'S' {
				continue
			}
			pipes = append(pipes, &pipe{
				typ:   r,
				coord: coord{y: row, x: col},
			})
			depth = func() uint {
				for _, search := range []coord{above, below, left, right} {
					nextc := coord{y: row + search.y, x: col + search.x}
					nextPipe := diagram[nextc.y][nextc.x]
					var nextdir coord
					switch search {
					case above:
						switch nextPipe {
						case '|':
							nextdir = above
						case '7':
							nextdir = left
						case 'F':
							nextdir = right
						}
					case below:
						switch nextPipe {
						case '|':
							nextdir = below
						case 'L':
							nextdir = right
						case 'J':
							nextdir = left
						}
					case right:
						switch nextPipe {
						case '-':
							nextdir = right
						case '7':
							nextdir = below
						case 'J':
							nextdir = above
						}
					case left:
						switch nextPipe {
						case '-':
							nextdir = left
						case 'F':
							nextdir = below
						case 'L':
							nextdir = above
						}
					default:
						panic("Impossible coordinate")
					}
					var zero coord
					if nextdir != zero {
						pipes = append(pipes, &pipe{
							typ:   rune(nextPipe),
							coord: nextc,
						})
						return dfs(diagram, &pipes, depth+1, nextc, nextdir)
					}
				}
				return 2 << 32
			}()
		}
	}
	printGrid(diagram, pipes)
	f := depth / 2
	if depth%2 != 0 {
		f++ // Round up
	}
	slog.Info("Looped", "depth", depth, "farthest", f)
}

func dfs(
	diagram []string, pipes *[]*pipe, depth uint, next, search coord,
) uint {
	if diagram[next.y][next.x] == 'S' {
		return depth + 1
	}

	dy, dx := search.y+next.y, search.x+next.x
	nextc := coord{y: dy, x: dx}
	p := &pipe{typ: rune(diagram[dy][dx]), coord: nextc}
	// NOTE(jay): Noisy -- slog.Info("In DFS", "pipe", p, "depth", depth)
	*pipes = append(*pipes, p)
	nextPipe := p.typ
	switch search {
	case above:
		switch nextPipe {
		case '|':
			depth = dfs(diagram, pipes, depth+1, nextc, above)
		case '7':
			depth = dfs(diagram, pipes, depth+1, nextc, left)
		case 'F':
			depth = dfs(diagram, pipes, depth+1, nextc, right)
		}
	case below:
		switch nextPipe {
		case '|':
			depth = dfs(diagram, pipes, depth+1, nextc, below)
		case 'L':
			depth = dfs(diagram, pipes, depth+1, nextc, right)
		case 'J':
			depth = dfs(diagram, pipes, depth+1, nextc, left)
		}
	case right:
		switch nextPipe {
		case '-':
			depth = dfs(diagram, pipes, depth+1, nextc, right)
		case '7':
			depth = dfs(diagram, pipes, depth+1, nextc, below)
		case 'J':
			depth = dfs(diagram, pipes, depth+1, nextc, above)
		}
	case left:
		switch nextPipe {
		case '-':
			depth = dfs(diagram, pipes, depth+1, nextc, left)
		case 'F':
			depth = dfs(diagram, pipes, depth+1, nextc, below)
		case 'L':
			depth = dfs(diagram, pipes, depth+1, nextc, above)
		}
	default:
		panic("Impossible coordinate")
	}
	return depth
}

func printGrid(diagram []string, pipes []*pipe) {
	contains := func(pipes []*pipe, x, y int) (int, bool) {
		for i, p := range pipes {
			if p.x == x && p.y == y {
				return i, true
			}
		}
		return -1, false
	}
	var sb strings.Builder
	for y, line := range diagram {
		for x, r := range line {
			if i, ok := contains(pipes, x, y); ok {
				r = pipes[i].typ
				switch r {
				case '|':
					r = '┃'
				case '-':
					r = '━'
				case 'L':
					r = '┗'
				case 'J':
					r = '┛'
				case '7':
					r = '┓'
				case 'F':
					r = '┏'
				case '.':
					r = '░'
				}
			}
			sb.WriteRune(r)
		}
		sb.WriteByte('\n')
	}
	fmt.Println(sb.String())
}

// part2:
func part2(input string) {
	panic("Unimplemented")
}
