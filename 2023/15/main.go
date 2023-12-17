package fifteen

import (
	"aoc/util"
	_ "embed"
	"log/slog"
	"strings"
)

var (
	// example1: This initialization sequence specifies 11 individual steps; the
	// result of running the HASH algorithm on each of the steps is as follows:
	//
	// rn=1 becomes 30.
	// cm- becomes 253.
	// qp=3 becomes 97.
	// cm=2 becomes 47.
	// qp- becomes 14.
	// pc=4 becomes 180.
	// ot=9 becomes 9.
	// ab=5 becomes 197.
	// pc- becomes 48.
	// pc=6 becomes 214.
	// ot=7 becomes 231.
	example1 = strings.NewReader(`
rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7
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
// Determine the ASCII code for the current character of the string.
// Increase the current value by the ASCII code you just determined.
// Set the current value to itself multiplied by 17.
// Set the current value to the remainder of dividing itself by 256.
//
// Run the HASH algorithm on each step in the initialization sequence.
// What is the sum of the results? (The initialization sequence is one long
// line; be careful when copy-pasting it.)
func part1(input string) {
	var sums []int
	total := 0
	for _, step := range strings.FieldsFunc(input, func(r rune) bool {
		return r == ','
	}) {

		sum := 0
		for _, r := range step {
			sum += int(r)
			sum *= 17
			sum %= 256
		}
		total += sum
		sums = append(sums, sum)
	}
	slog.Error("Result", "total", total)
}

// part2:
func part2(input string) {
	panic("Unimplemented")
}
