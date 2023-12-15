package thirteen

import (
	"aoc/util"
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
	for i, block := range strings.Split(input, "\n\n") {
		pat := strings.Split(block, "\n")
		n := findReflection(pat) * 100
		if n == 0 {
			pat = rotate(pat)
			slog.Info("Row pattern not found", "block", i+1, "rotated", pat)
			n = findReflection(pat)
		}
		sum += n
	}
	slog.Error("Result", "sum", sum)
}

// part2:
func part2(input string) {
	panic("Unimplemented")
}

func rotate(pattern []string) []string {
	if len(pattern) == 0 || len(pattern[0]) == 0 {
		return pattern
	}

	rows, cols := len(pattern), len(pattern[0])
	result := make([]string, cols)

	for i := 0; i < cols; i++ {
		build := make([]byte, rows)
		for j := 0; j < rows; j++ {
			build[j] = pattern[rows-j-1][i]
		}
		result[i] = string(build)
	}

	return result
}

func findReflection(pattern []string) int {
	for i := 0; i < len(pattern)-1; i++ {
		if isMirrored(pattern, i, i+1) {
			slog.Info("Mirrored", "row 1", i+1, "row 2", i+2,
				"line 1", pattern[i], "line 2", pattern[i+1])
			return i + 1
		}
	}
	return 0
}

func isMirrored(pattern []string, i, j int) bool {
	for ; i > -1 && j < len(pattern); i, j = i-1, j+1 {
		if pattern[i] != pattern[j] {
			return false
		}
	}
	return true
}
