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
.....
.S-7.
.|.|.
.L-J.
.....
`[1:])

	//FF7FSF7F7F7F7F7F---7
	//L|LJ||||||||||||F--J
	//FL-7LJLJ||||||LJL-77
	//F--JF--7||LJLJ7F7FJ-
	//L---JF-JLJ.||-FJLJJ7
	//|F|F-JF---7F7-L7L|7|
	//|FFJF7L7F-JF7|JL---7
	//7-L-JL7||F7|L7F-7F7|
	//L.L7LFJ|||||FJL7||LJ
	//L7JLJL-JLJLJL--JLJ.L

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
					nextdir := followPipe(search, rune(nextPipe))
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
		return depth
	}

	c := coord{y: search.y + next.y, x: search.x + next.x}
	t := rune(diagram[c.y][c.x])
	*pipes = append(*pipes, &pipe{typ: t, coord: c})
	return dfs(diagram, pipes, depth+1, c, followPipe(search, t))
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

// part2: Need to iterate for each row and keep an `isInside` bool that
// starts false. We then iterate for each tile, and for each pipe that has a
// side pointing north (And belonging to a hashset that tells us which is the
// loop path, to remove the junk pipes) we invert the `isInside` variable.
//
// For each tile that is not a pipe belonging to the path hashset, we check if
// that tile is inside or not, if that's the case, increment your `insideCount`
// variable.
//
// NOTE: This requires that we transform our starting point into a pipe, so we
// don't encounter any weird bugs with the tile counting.
//
// NOTE(jay): Also, this algo is used for drawing vector graphics!
func part2(input string) {
	diagram := strings.Split(input, "\n")
	pipes := map[coord]*pipe{}
	for row, line := range diagram {
		for col, r := range line {
			if r != 'S' {
				continue
			}

			var dirToStart1, dirToStart2 coord
			var nextc, nextdir coord
			for _, search := range []coord{above, below, left, right} {
				dy, dx := row+search.y, col+search.x
				if !(0 <= dy && dy < len(diagram)) ||
					!(0 <= dx && dx < len(diagram[dy])) {
					continue
				}
				nextc = coord{y: dy, x: dx}
				nextPipe := diagram[dy][dx]
				followDir := followPipe(search, rune(nextPipe))
				slog.Info("Follow Pipe", "way", search, "pipe", string(nextPipe),
					"got", nextdir)
				var zero coord
				if dirToStart1 == zero {
					dirToStart1 = search
				}
				dirToStart2 = search
				if followDir != zero {
					nextdir = followDir
					slog.Info("Found", "way", nextdir, "dir1", dirToStart1, "dir2", dirToStart2)
					pipes[nextc] = &pipe{
						typ:   rune(nextPipe),
						coord: nextc,
					}
				}
			}
			// Need to map 'S' to it's actual value
			pipes[coord{y: row, x: col}] = &pipe{
				typ:   mapTypeFrom(dirToStart1, dirToStart2),
				coord: coord{y: row, x: col},
			}
			slog.Info("WTF", "pipes", pipes)
			dfs2(diagram, pipes, nextc, nextdir)
		}
	}

	// part2 begins
	var insideCount uint
	for row, line := range diagram {
		var isInside bool
		for col, r := range line {
			p, ok := pipes[coord{y: row, x: col}]
			if !ok && isInside {
				slog.Info("Adding", "rune", string(r))
				insideCount++
				continue
			}
			if !ok {
				continue
			}
			switch p.typ {
			case '|', 'J', 'L':
				isInside = !isInside
			}
		}
	}
	slog.Info("I want the", "pipes", pipes, "Inside should be", insideCount)
}

func mapTypeFrom(d1, d2 coord) rune {
	switch {
	default:
		panic("Invalid directions for starting position 'S' " +
			fmt.Sprintf("%+v %+v", d1, d2))
	case d1 == above && d2 == below:
		return '|'
	case d1 == above && d2 == right:
		return 'F'
	case d1 == above && d2 == left:
		return '7'
	case d1 == below && d2 == left:
		return 'J'
	case d1 == below && d2 == right:
		return 'L'
	case d1 == left && d2 == right:
		return '-'
	}
}

func followPipe(dir coord, nextPipe rune) coord {
	var nextdir coord
	switch dir {
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
	return nextdir
}

func dfs2(diagram []string, pipes map[coord]*pipe, next, search coord) {
	if diagram[next.y][next.x] == 'S' {
		return
	}

	c := coord{y: search.y + next.y, x: search.x + next.x}
	t := rune(diagram[c.y][c.x])
	if _, ok := pipes[c]; !ok {
		pipes[c] = &pipe{typ: t, coord: c}
	}
	// NOTE(jay): Noisy -- slog.Info("DFS", "pipe", pipes[c])
	dfs2(diagram, pipes, c, followPipe(search, t))
}
