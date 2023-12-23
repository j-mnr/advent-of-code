package fifteen

import (
	"aoc/util"
	_ "embed"
	"log/slog"
	"slices"
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

	// example2:  Here is the contents of every box after each step in the example
	// initialization sequence above:
	//
	// After "rn=1":
	// Box 0: [rn 1]
	//
	// After "cm-":
	// Box 0: [rn 1]
	//
	// After "qp=3":
	// Box 0: [rn 1]
	// Box 1: [qp 3]
	//
	// After "cm=2":
	// Box 0: [rn 1] [cm 2]
	// Box 1: [qp 3]
	//
	// After "qp-":
	// Box 0: [rn 1] [cm 2]
	//
	// After "pc=4":
	// Box 0: [rn 1] [cm 2]
	// Box 3: [pc 4]
	//
	// After "ot=9":
	// Box 0: [rn 1] [cm 2]
	// Box 3: [pc 4] [ot 9]
	//
	// After "ab=5":
	// Box 0: [rn 1] [cm 2]
	// Box 3: [pc 4] [ot 9] [ab 5]
	//
	// After "pc-":
	// Box 0: [rn 1] [cm 2]
	// Box 3: [ot 9] [ab 5]
	//
	// After "pc=6":
	// Box 0: [rn 1] [cm 2]
	// Box 3: [ot 9] [ab 5] [pc 6]
	//
	// After "ot=7":
	// Box 0: [rn 1] [cm 2]
	// Box 3: [ot 7] [ab 5] [pc 6]
	example2 = strings.NewReader(`
rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7
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
		hash(step)
		total += sum
		sums = append(sums, sum)
	}
	slog.Error("Result", "total", total)
}

// part2: The book goes on to describe a series of 256 boxes numbered 0 through
// 255. The boxes are arranged in a line starting from the point where light
// enters the facility. The boxes have holes that allow light to pass from one
// box to the next all the way down the line.
//
// The book goes on to explain how to perform each step in the initialization
// sequence, a process it calls the Holiday ASCII String Helper Manual
// Arrangement Procedure, or HASHMAP for short.
//
// The focusing power of a single lens is the result of multiplying together:
//
// - One plus the box number of the lens in question.
// - The slot number of the lens within the box: 1 for the first lens, 2 for the
// second lens, and so on.
// - The focal length of the lens.
//
// What is the focusing power of the resulting lens configuration?
func part2(input string) {
	const (
		opEq   = '='
		opDash = '-'
	)
	type lens struct {
		label  string
		length uint8 // length 1-9
	}
	boxes := [256][]lens{}
	for _, step := range strings.FieldsFunc(input, func(r rune) bool {
		return r == ','
	}) {
		i := strings.LastIndexAny(step, string([]byte{opEq, opDash}))
		if i == -1 {
			panic("Bad step: " + step)
		}
		lbl := step[:i]
		var op, leng byte
		op = step[i]
		if op == opEq {
			leng = step[i+1] - '0'
		}
		slog.Info("Hashing", "result", hash(lbl))

		hsh := hash(lbl)
		i = slices.IndexFunc(boxes[hsh], func(l lens) bool {
			return l.label == step[:i]
		})
		switch op {
		case opEq:
			if i == -1 {
				boxes[hsh] = append(boxes[hsh], lens{label: lbl, length: leng})
				continue
			}
			boxes[hsh][i] = lens{label: lbl, length: leng}
		case opDash:
			if i == -1 {
				continue
			}
			boxes[hsh] = append(boxes[hsh][:i], boxes[hsh][i+1:]...)
		default:
			panic("Bad Operator: " + string(op))
		}
	}

	total := 0
	for i, box := range boxes {
		if len(box) == 0 {
			continue
		}
		slog.Info("Found", "n", i+1, "box", box)
		for k, lens := range box {
			total += (i + 1) * (k + 1) * int(lens.length)
		}
	}
	slog.Error("Result", "total", total)
}

func hash(step string) int {
	result := 0
	for _, r := range step {
		result += int(r)
		result *= 17
		result %= 256
	}
	return result
}
