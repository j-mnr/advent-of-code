package TODO

import (
	_ "embed"
	"strings"
)

var (
	// example1
	example1 = strings.NewReader(`
`[1:])

	// example2
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
	panic("Unimplemented")
}

// part2:
func part2(input string) {
	panic("Unimplemented")
}
