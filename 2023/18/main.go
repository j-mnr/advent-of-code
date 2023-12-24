package eighteen

import (
	"aoc/util"
	_ "embed"
	"log/slog"
	"strconv"
	"strings"
)

var (
	// example1:
	example1 = strings.NewReader(`
R 6 (#70c710)
D 5 (#0dc571)
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
// https://www.themathdoctors.org/polygon-coordinates-and-areas/
// XXX: 54632 too high
func part1(input string) {
	type coord struct{ x, y int }

	coords := []coord{{0, 0}}
	var x, y int
	for _, line := range strings.Split(input, "\n") {
		f := strings.Fields(line)
		switch f[0][0] {
		case left:
			x -= util.Must2(strconv.Atoi(f[1]))
		case right:
			x += util.Must2(strconv.Atoi(f[1]))
		case up:
			y -= util.Must2(strconv.Atoi(f[1]))
		case down:
			y += util.Must2(strconv.Atoi(f[1]))
		default:
			panic("Impossible direction " + string(f[0][0]))
		}
		coords = append(coords, coord{x: x, y: y})
	}
	// coords = []coord{{-2, -2}, {0, 4}, {3, -1}, {1, -1}}
	sum1, sum2 := 0, 0
	for i := 0; i < len(coords)-1; i++ {
		c1, c2 := coords[i], coords[i+1]
		// xn-1 * yn - xn * yn-1
		sum1 += c1.x*c2.y
		sum2 += c1.y*c2.x
	}
	slog.Error("Result", "coords", coords, "sum1", sum1, "sum2", sum2, "minus",
	(sum1-sum2)/2)
}

// part2:
func part2(input string) {
	panic("Unimplemented")
}
