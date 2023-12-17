package fourteen

import (
	"aoc/util"
	"bytes"
	_ "embed"
	"log/slog"
	"slices"
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

// part2: Each cycle tilts the platform four times so that the rounded rocks
// roll north, then west, then south, then east. After each tilt, the rounded
// rocks roll as far as they can before the platform tilts in the next
// direction. After one cycle, the platform will have finished rolling the
// rounded rocks in those four directions in that order.
//
// Run the spin cycle for 1_000_000_000 cycles. Afterward, what is the total load
// on the north support beams?
// ANSWER: 93742
func part2(input string) {
	platform := bytes.Split([]byte(input), []byte("\n"))
	var cycleOfCycles []int
	initial := 0
	start, end := -1, -2
	for i := 1; i <= 1500; i++ {
		sum := cycle(platform, i)

		if start != -1 {
			start++
			if sum != cycleOfCycles[start] {
				slog.Info("Broke cycle")
				start = -1
				end = -2
			}
			if start == end {
				slog.Info("Found a cycle!", "cycle", i, "start", initial, "end", end,
					"nums", cycleOfCycles[initial:end])
				break
			}
		}
		if i := slices.Index(cycleOfCycles, sum); i != -1 && start == -1 {
			start, initial = i, i

			end = len(cycleOfCycles)
			slog.Info("Found a repeat", "start", start, "end", end,
				"nums", cycleOfCycles)
		}
		cycleOfCycles = append(cycleOfCycles, sum)
		// Reliably break when cycle found.
	}
	blindGuess := (1_000_000_000-16)%(end-initial)
	slog.Error("I hate math", "NO idea", blindGuess,
	"between these values", cycleOfCycles[initial:end])
}

func cycle(platform [][]byte, times int) int {
	for n := 0; n < times; n++ {
		for i := 1; i < len(platform); i++ { // Move north
			for k, r := range platform[i] {
				if r != rndRock {
					continue
				}

				for row := i - 1; row > -1; row-- {
					if platform[row][k] != empty {
						break
					}
					platform[row][k] = rndRock
					platform[row+1][k] = empty
				}
			}
		}

		for k := 1; k < len(platform[0]); k++ { // Move west
			for i := 0; i < len(platform); i++ {
				if platform[i][k] != rndRock {
					continue
				}

				for col := k - 1; col > -1; col-- {
					if platform[i][col] != empty {
						break
					}
					//slog.Info("Moving", "row", i, "col", col, "line", platform[i])
					platform[i][col] = rndRock
					platform[i][col+1] = empty
				}
			}
		}

		for i := len(platform) - 2; i > -1; i-- { // Move south
			for k, r := range platform[i] {
				if r != rndRock {
					continue
				}

				for row := i + 1; row < len(platform); row++ {
					if platform[row][k] != empty {
						break
					}
					//slog.Info("Moving", "row", row, "col", k, "line", platform[i])
					platform[row][k] = rndRock
					platform[row-1][k] = empty
				}
			}
		}

		for k := len(platform[0]) - 2; k > -1; k-- { // Move east
			for i := 0; i < len(platform); i++ {
				if platform[i][k] != rndRock {
					continue
				}

				for col := k + 1; col < len(platform[0]); col++ {
					if platform[i][col] != empty {
						break
					}
					//slog.Info("Moving", "row", i, "col", col, "line", platform[i])
					platform[i][col] = rndRock
					platform[i][col-1] = empty
				}
			}
		}
	}
	sum := 0
	for i, line := range platform {
		c := bytes.Count(line, []byte{rndRock})
		//slog.Info("Found rocks", "count", c, "row", len(platform)-i)
		sum += (c * (len(platform) - i))
	}
	slog.Info("cycle", "sum", sum)
	return sum
}
