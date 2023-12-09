package nine

import (
	"aoc/util"
	_ "embed"
	"log/slog"
	"strconv"
	"strings"
)

var (
	// example1: tl;dr: find the derivatives and give next value in list.
	// - f'(3n) == 3, f''(3) == 0
	// - f'(1/2n^2+1/2n) == n+1/2, f''(n+1/2) == 1, f'''(1) == 0
	// - f'(1/3 (15 + 20 n - 6 n^2 + n^3)) == 20/3 - 4 n + n^2,
	// f''(20/3 - 4 n + // n^2) == 2(-2 + n), f'''(2(-2 + n)) == 2, f''''(2) == 0
	example1 = strings.NewReader(`
0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45
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
func part1(input string) {
	lines := strings.Split(input, "\n")
	series := make([][]int, len(lines))
	for i, line := range lines {
		series[i] = append(series[i], util.SlicesMap[[]string, []int](
			strings.Fields(line), func(a string) int {
				return util.Must2(strconv.Atoi(a))
			})...)
	}
	slog.Info("Mapped values", "series", series)
	var results []int
	for _, s := range series {
		results = append(results, funcopop(s))
	}
	slog.Info("Last of series collected", "sum", util.SlicesReduce(results,
	func(a, b int) int {
		return a+b
	}, 0))
}

func funcopop(series []int) int {
	lastVals := []int{series[len(series)-1]}
	next := series
	allZero := func(diffs []int) bool {
		for _, d := range diffs {
			if d != 0 {
				return false
			}
		}
		return true
	}
	for {
		var diffs []int
		for i := 1; i < len(next); i++ {
			diffs = append(diffs, next[i]-next[i-1])
		}
		slog.Info("Diffs found", "diffs", diffs)
		if allZero(diffs) {
			return util.SlicesReduce(lastVals, func(a, b int) int {
				slog.Info("reducing values", "a", a, "b", b)
				return a + b
			}, 0)
		}
		next = diffs
		lastVals = append(lastVals, diffs[len(diffs)-1])
	}
	panic("Impossible to get here!")
}

// part2:
func part2(input string) {
	panic("Unimplemented")
}
