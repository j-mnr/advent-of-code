package fourteen

import (
	"aoc/util"
	"bytes"
	_ "embed"
	"log/slog"
	"strings"
)

const (
	rndRock  = 'O'
	cubeRock = '#'
	empty    = '.'
)

var (
	// example1:
	example1 = strings.NewReader(`
O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....
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
// Tilt the platform so that the rounded rocks all roll north. Afterward, what
// is the total load on the north support beams?
//
// The amount of load caused by a single rounded rock (O) is equal to the number
// of rows from the rock to the south edge of the platform, including the row
// the rock is on.
// The total load is the sum of the load caused by all of the rounded rocks.
func part1(input string) {
	platform := bytes.Split([]byte(input), []byte("\n"))
	for i := 1; i < len(platform); i++ {
		for k, r := range platform[i] {
			if r != rndRock {
				continue
			}

			for row := i - 1; row > -1; row-- { // Move upwards
				if platform[row][k] != empty {
					break
				}
				platform[row][k] = rndRock
				platform[row+1][k] = empty
			}
		}
	}
	sum := 0
	for i, line := range platform {
		c := bytes.Count(line, []byte{rndRock})
		slog.Info("Found rocks", "count", c, "row", len(platform)-i)
		sum += (c * (len(platform) - i))
	}
	slog.Error("Result", "total load", sum)
}

// part2:
func part2(input string) {
	panic("Unimplemented")
}
