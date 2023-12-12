package eleven

import (
	"aoc/util"
	"bytes"
	_ "embed"
	"fmt"
	"math"
	"log/slog"
	"strings"
)

const (
	galaxy = '#'
	space  = '.'
)

var (
	// example1:
	// column 3, 5, 8 need to expand. row 3, 7 need to expand.
	// Shortest path is done by moving up, left, right, and down **NOT** diagonal.
	example1 = strings.NewReader(`
...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....
`[1:])

	// example2:
	example2 = strings.NewReader(`
...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....
`[1:])

	//go:embed input.txt
	input string
)

type coord struct{ y, x int }

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

// part1: The image includes empty space (.) and galaxies (#)
//
// Only some space expands. In fact, the result is that any rows or columns that
// contain no galaxies should all actually be twice as big.
//
// Equipped with this expanded universe, the shortest path between every pair of
// galaxies can be found.
//
// Expand the universe, then find the length of the shortest path between every
// pair of galaxies. What is the sum of these lengths?
func part1(input string) {
	data := bytes.Split([]byte(input), []byte("\n"))
	img := make([][]byte, 0, len(data))
	for _, line := range data {
		if !bytes.Contains(line, []byte{galaxy}) {
			img = append(img, line)
		}
		img = append(img, line)
	}

	for col := len(data[0]) - 1; col > -1; col-- {
		var hasGalaxy bool
		for _, row := range data {
			if row[col] == galaxy {
				hasGalaxy = true
				break
			}
		}
		if !hasGalaxy {
			for row := range img {
				img[row] = append(img[row][:col], append([]byte{space}, img[row][col:]...)...)
			}
		}
	}

	var coords []coord
	for row, line := range img {
		for col, r := range line {
			if byte(r) == galaxy {
				coords = append(coords, coord{y: row, x: col})
			}
		}
	}
	sum := 0
	for i := 0; i < len(coords)-1; i++ {
		for _, c2 := range coords[i+1:] {
			y := util.Abs(coords[i].y - c2.y)
			x := util.Abs(coords[i].x - c2.x)
			sum += y + x
		}
	}
	slog.Info("Results", "sum", sum)
}

func printImg(img [][]byte) {
	for _, row := range img {
		fmt.Println(string(row))
	}
}

// part2: Now, instead of the expansion you did before, make each empty row or
// column one million times larger. That is, each empty row should be replaced
// with 1000000 empty rows, and each empty column should be replaced with
// 1000000 empty columns.
//
// Find the pattern!
//
// 1   sum=9155896  0
// 10  sum=14277292 1
// 100 sum=65491252 2
// 1,000            3
// 10,000           4
// 100,000          5
// 1,000,000        6
// factor=5,121,396
func part2(input string) {
	sum1 := sumUp(expandBy(input, 0))
	sum10 := sumUp(expandBy(input, 9))

	constant := float64(sum10 - sum1)
	sum := float64(sum1)
	for i := float64(0); i < 6; i++ {
		sum += (constant * math.Pow(10, i))
	}
	slog.Info("Results", "constant", int64(constant), "scaled", int64(sum))
}

func expandBy(input string, expand int) [][]byte {
	data := bytes.Split([]byte(input), []byte("\n"))
	img := make([][]byte, 0, len(data))
	for _, line := range data {
		if !bytes.Contains(line, []byte{galaxy}) {
			for i := 0; i < expand; i++ {
				img = append(img, line)
			}
		}
		img = append(img, line)
	}

	for col := len(data[0]) - 1; col > -1; col-- {
		var hasGalaxy bool
		for _, row := range data {
			if row[col] == galaxy {
				hasGalaxy = true
				break
			}
		}
		if !hasGalaxy {
			for row := range img {
					img[row] = append(img[row][:col], append(bytes.Repeat([]byte{space},
					expand), img[row][col:]...)...)
			}
		}
	}
	return img
}

func sumUp(img [][]byte) int {
	var coords []coord
	for row, line := range img {
		for col, r := range line {
			if byte(r) == galaxy {
				coords = append(coords, coord{y: row, x: col})
			}
		}
	}
	sum := 0
	for i := 0; i < len(coords)-1; i++ {
		for _, c2 := range coords[i+1:] {
			y := util.Abs(coords[i].y - c2.y)
			x := util.Abs(coords[i].x - c2.x)
			sum += y + x
		}
	}
	return sum
}
