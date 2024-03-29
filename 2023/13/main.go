package thirteen

import (
	"aoc/util"
	"bytes"
	_ "embed"
	"log/slog"
	"strings"
)

const (
	ash  = '.'
	rock = '#'
)

var (
	// example1: In the first pattern, the reflection is across a vertical line
	// between two columns; arrows on each of the two columns point at the line
	// between the columns
	//
	// In this pattern, the line of reflection is the vertical line between
	// columns 5 and 6. Because the vertical line is not perfectly in the middle
	// of the pattern, part of the pattern (column 1) has nowhere to reflect onto
	// and can be ignored; every other column has a reflected column within the
	// pattern and must match exactly: column 2 matches column 9, column 3 matches
	// 8, 4 matches 7, and 5 matches 6.
	//
	// The second pattern reflects across a horizontal line instead
	//
	// This pattern reflects across the horizontal line between rows 4 and 5. Row
	// 1 would reflect with a hypothetical row 8, but since that's not in the
	// pattern, row 1 doesn't need to match anything. The remaining rows match:
	// row 2 matches row 7, row 3 matches row 6, and row 4 matches row 5.
	example1 = strings.NewReader(`
#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#

.#.##.#.#
.##..##..
.#.##.#..
#......##
#......##
.#.##.#..
.##..##.#

#..#....#
###..##..
.##.#####
.##.#####
###..##..
#..#....#
#..##...#

#.##..##.
..#.##.#.
##..#...#
##...#..#
..#.##.#.
..##..##.
#.#.##.#.
`[1:])

	// example2:
	example2 = strings.NewReader(`
#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#

.#.##.#.#
.##..##..
.#.##.#..
#......##
#......##
.#.##.#..
.##..##.#

#..#....#
###..##..
.##.#####
.##.#####
###..##..
#..#....#
#..##...#

#.##..##.
..#.##.#.
##..#...#
##...#..#
..#.##.#.
..##..##.
#.#.##.#.
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

// part1: To find the reflection in each pattern, you need to find a perfect
// reflection across either a horizontal line between two rows or across a
// vertical line between two columns.
//
// To summarize your pattern notes, add up the number of columns to the
// left of each vertical line of reflection; to that, also add 100 multiplied by
// the number of rows above each horizontal line of reflection. In the above
// example, the first pattern's vertical line has 5 columns to its left and the
// second pattern's horizontal line has 4 rows above it, a total of 405.
//
// Find the line of reflection in each of the patterns in your notes. What
// number do you get after summarizing all of your notes?
func part1(input string) {
	sum := 0
	for i, block := range bytes.Split([]byte(input), []byte("\n\n")) {
		pat := bytes.Split(block, []byte("\n"))
		n := findReflection(pat) * 100
		if n == 0 {
			slog.Info("Row pattern not found", "block", i+1)
			n = findReflection(rotate(pat))
		}
		sum += n
	}
	slog.Error("Result", "sum", sum)
}

func findReflection(pattern [][]byte) int {
	for i := 0; i < len(pattern)-1; i++ {
		if isMirrored(pattern, i, i+1) {
			slog.Info("Mirrored", "row 1", i+1, "row 2", i+2,
				"line 1", pattern[i], "line 2", pattern[i+1])
			return i + 1
		}
	}
	return 0
}

// part2:
func part2(input string) {
	sum := 0
	for i, block := range bytes.Split([]byte(input), []byte("\n\n"))[0:2] {
		pat := bytes.Split(block, []byte("\n"))
		n := findSmudge(pat) * 100
		if n == 0 {
			slog.Info("Row pattern not found", "block", i+1)
			n = findSmudge(rotate(pat))
		}
		sum += n
	}
	slog.Error("Result", "sum", sum)
}

func findSmudge(pattern [][]byte) int {
	flip := func(pattern [][]byte, i, j int) {
		switch pattern[i][j] {
		case ash:
			pattern[i][j] = rock
		case rock:
			pattern[i][j] = ash
		default:
			panic("Impossible character: " + string(pattern[i]))
		}
	}
	n, m := len(pattern)-1, len(pattern[0])
	for j := 0; j < m; j++ {
		for i := 0; i < n; i++ {
			original := pattern[i]
			flip(pattern, i, j)
			for k := 0; k < n; k++ {
				slog.Info("Flipped",
					"line 1", pattern[k], "line 2", pattern[k+1],
				)
				if isMirrored(pattern, k, k+1) {
					slog.Info("Mirrored", "row 1", k+1, "row 2", k+2,
						"line 1", pattern[k], "line 2", pattern[k+1])
					pattern[i] = original
					return k + 1
				}
			}
			pattern[i] = original
		}
	}
	return 0
}

func rotate(pattern [][]byte) [][]byte {
	if len(pattern) == 0 || len(pattern[0]) == 0 {
		return pattern
	}

	rows, cols := len(pattern), len(pattern[0])
	result := make([][]byte, cols)
	for i := 0; i < cols; i++ {
		build := make([]byte, rows)
		for j := 0; j < rows; j++ {
			build[j] = pattern[rows-j-1][i]
		}
		result[i] = build
	}
	return result
}

func isMirrored(pattern [][]byte, i, j int) bool {
	for ; i > -1 && j < len(pattern); i, j = i-1, j+1 {
		if !bytes.Equal(pattern[i], pattern[j]) {
			return false
		}
	}
	return true
}
