package three

import (
	_ "embed"
	"strings"
)

var (
	example1 = strings.NewReader(``)
	example2 = strings.NewReader(``)

	//go:embed input.txt
	input string
)

func Run() {
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
