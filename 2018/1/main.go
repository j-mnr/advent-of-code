package one

import (
	"aoc/util"
	_ "embed"
	"log/slog"
	"strconv"
	"strings"
)

var (
	// example1:
	// Starting from a frequency of zero, the following changes would occur:
	// - Current frequency  0, change of +1; resulting frequency  1.
	// - Current frequency  1, change of -2; resulting frequency -1.
	// - Current frequency -1, change of +3; resulting frequency  2.
	// - Current frequency  2, change of +1; resulting frequency  3.
	// In this example, the resulting frequency is 3.
	example1 = strings.NewReader(`
+1
-2
+3
+1
`[1:])

	// example2:
	// Starting from a frequency of zero
	//
	// Current frequency  0, change of +1; resulting frequency  1.
	// Current frequency  1, change of -2; resulting frequency -1.
	// Current frequency -1, change of +3; resulting frequency  2.
	// Current frequency  2, change of +1; resulting frequency  3.
	// (At this point, the device continues from the start of the list.)
	// Current frequency  3, change of +1; resulting frequency  4.
	// Current frequency  4, change of -2; resulting frequency  2, which has already been seen.
	example2 = strings.NewReader(`
+1
-2
+3
+1
+1
-2
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

// part1: Starting with a frequency of zero, what is the resulting frequency
// after all of the changes in frequency have been applied?
func part1(input string) {
	sum := 0
	for _, line := range strings.Split(input, "\n") {
		v := util.Must2(strconv.Atoi(line))
		slog.Info("Parsed line", "value", v, "current sum", sum)
		sum += v
	}
	slog.Info("Calculated frequency", "result", sum)
}

// part2: You notice that the device repeats the same frequency change list over
// and over. To calibrate the device, you need to find the first frequency it
// reaches twice.
//
// NOTE: your device might need to repeat its list of frequency changes many
// times before a duplicate frequency is found
func part2(input string) {
	sum := 0
	seen := map[int]struct{}{}
	for {
		for _, line := range strings.Split(input, "\n") {
			v := util.Must2(strconv.Atoi(line))
			if _, ok := seen[sum]; ok {
				slog.Info("Found repeat frequency", "result", sum)
				return
			}
			seen[sum] = struct{}{}
			sum += v
			slog.Info("Put in seen", "value", sum)
		}
	}
}
