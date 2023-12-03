package three

import (
	"aoc/util"
	_ "embed"
	"log/slog"
	"strconv"
	"strings"
	"unicode"
)

var (
	// example1: In this schematic, two numbers are not part numbers because they
	// are not adjacent to a symbol: 114 (top right) and 58 (middle right). Every
	// other number is adjacent to a symbol and so is a part number; their sum is
	// 4361.
	example1 = strings.NewReader(`
467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`[1:])

	// example2: In this schematic, there are two gears. The first is in the top
	// left; it has part numbers 467 and 35, so its gear ratio is 16345. The
	// second gear is in the lower right; its gear ratio is 451490. (The *
	// adjacent to 617 is not a gear because it is only adjacent to one part
	// number.) Adding up all of the gear ratios produces 467835.
	example2 = strings.NewReader(`
467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`[1:])

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

// part1: There are lots of numbers and symbols you don't really understand, but
// apparently any number adjacent to a symbol, even diagonally, is a "part
// number" and should be included in your sum.
//
// What is the sum of all of the part numbers in the engine schematic?
func part1(input string) {
	schematic := strings.Split(input, "\n")
	used := make([][]bool, len(schematic))
	for i := range used {
		used[i] = make([]bool, len(schematic[i]))
	}
	sum := 0
	for y, line := range schematic {
		for x, symbol := range line {
			if unicode.IsDigit(symbol) || symbol == '.' {
				continue
			}
			slog.Info("Symbol found", slog.String("symbol", string(symbol)),
				slog.Int("line", y+1), slog.Int("position", x+1))
			for _, delta := range []struct{ y, x int }{
				{-1, -1}, {-1, 0}, {-1, 1},
				{0, -1} /*{0,0}*/, {0, 1},
				{1, -1}, {1, 0}, {1, 1},
			} {
				dy, dx := delta.y+y, delta.x+x
				if !(dy >= 0 && dy < len(schematic)) ||
					!(dx >= 0 && dx < len(schematic[dy])) {
					continue
				}
				if used[dy][dx] { // redundant check for logging
					slog.Info("Used digit found", slog.String("digit", string(schematic[dy][dx])),
						slog.Int("line", dy+1), slog.Int("position", dx+1))
				}
				d := rune(schematic[dy][dx])
				if used[dy][dx] || !unicode.IsDigit(d) {
					continue
				}
				slog.Info("Digit found", slog.String("digit", string(d)),
					slog.Int("line", dy+1), slog.Int("position", dx+1))
				used[dy][dx] = true

				left, right := dx+-1, dx+1
				for ; left >= 0 && unicode.IsDigit(rune(schematic[dy][left])); left-- {
					used[dy][left] = true
				}
				for ; right < len(schematic[dy]) &&
					unicode.IsDigit(rune(schematic[dy][right])); right++ {
					used[dy][right] = true
				}
				slog.Info("Number found",
					slog.Int("line", dy+1),
					slog.Int("left", left+1), slog.Int("right", right),
					slog.String("number", schematic[dy][left+1:right]),
				)
				sum += util.Must2(strconv.Atoi(schematic[dy][left+1 : right]))
			}
		}
	}

	for y, line := range used {
		for x, used := range line {
			if !used {
				continue
			}
			slog.Info("Used digit", slog.String("digit", string(schematic[y][x])),
				slog.Int("line", y+1), slog.Int("position", x+1))
		}
	}

	slog.Info("Sum of all part numbers", slog.Int("sum", sum))
}

// part2: one of the gears in the engine is wrong. A gear is any '*' symbol that
// is adjacent to exactly two part numbers. Its gear ratio is the result of
// multiplying those two numbers together.
//
// This time, you need to find the gear ratio of every gear and add them all up
// so that the engineer can figure out which gear needs to be replaced.
func part2(input string) {
	schematic := strings.Split(input, "\n")
	used := make([][]bool, len(schematic))
	for i := range used {
		used[i] = make([]bool, len(schematic[i]))
	}
	sum := uint(0)
	for y, line := range schematic {
		for x, gear := range line {
			if gear != '*' {
				continue
			}
			slog.Info("Gear found", slog.Int("line", y+1), slog.Int("position", x+1))
			var numbers []uint // Follow this to see where part 2 updates
			for _, delta := range []struct{ y, x int }{
				{-1, -1}, {-1, 0}, {-1, 1},
				{0, -1} /*{0,0}*/, {0, 1},
				{1, -1}, {1, 0}, {1, 1},
			} {
				if len(numbers) > 2 {
					break
				}
				dy, dx := delta.y+y, delta.x+x
				if !(dy >= 0 && dy < len(schematic)) ||
					!(dx >= 0 && dx < len(schematic[dy])) {
					continue
				}
				if used[dy][dx] { // redundant check for logging
					slog.Info("Used digit found", slog.String("digit", string(schematic[dy][dx])),
						slog.Int("line", dy+1), slog.Int("position", dx+1))
				}
				isDigitB := func(b byte) bool { return unicode.IsDigit(rune(b)) }
				if used[dy][dx] || !isDigitB(schematic[dy][dx]) {
					continue
				}

				left, right := dx+-1, dx+1
				for ; left >= 0 && isDigitB(schematic[dy][left]); left-- {
					used[dy][left] = true
				}
				for ; right < len(schematic[dy]) && isDigitB(schematic[dy][right]); right++ {
					used[dy][right] = true
				}
				slog.Info("Number found",
					slog.Int("line", dy+1),
					slog.Int("left", left+1), slog.Int("right", right),
					slog.String("number", schematic[dy][left+1:right]),
				)
				numbers = append(numbers,
					uint(util.Must2(strconv.Atoi(schematic[dy][left+1:right]))))
			}
			if len(numbers) != 2 {
				slog.Info("Gear found, but not 2 numbers",
					slog.Int("line", y+1), slog.Int("position", x+1),
					"numbers", numbers)
				continue
			}
			sum += numbers[0] * numbers[1]
		}
	}

	slog.Info("Sum of all gear ratios", slog.Int("sum", int(sum)))
}
