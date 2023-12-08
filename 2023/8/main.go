package eight

import (
	"aoc/util"
	_ "embed"
	"log/slog"
	"strings"
)

var (
	// example1: start with AAA and go right (R) by choosing the right element of
	// AAA, CCC. Then, L means to choose the left element of CCC, ZZZ. By
	// following the left/right instructions, you reach ZZZ in 2 steps.
	example1 = strings.NewReader(`
RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)
`[1:])

	// example2:
	example2 = strings.NewReader(`
LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)
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

type node struct {
	name, left, right string
}

// part1: Starting at AAA, follow the left/right instructions. How many steps
// are required to reach ZZZ?
func part1(input string) {
	data := strings.Split(input, "\n")
	instructions := data[0]
	nodes := make(map[string]node, len(data[2:]))
	for _, line := range data[2:] {
		ff := strings.Fields(line)
		nodes[ff[0]] = node{
			name:  ff[0],
			left:  ff[2][1:4],
			right: ff[3][:3],
		}
	}
	slog.Info("Data collected", "instructions", instructions, "nodes", nodes)
	name := "AAA"
	steps := 0
	for {
		for _, dir := range instructions {
			switch dir {
			case 'R':
				slog.Info("Right switch", "node", nodes[name],
					"name", nodes[name].right)
				name = nodes[name].right
			case 'L':
				slog.Info("Left switch", "node", nodes[name],
					"name", nodes[name].left)
				name = nodes[name].left
			default:
				panic("Impossible input")
			}
			steps++
			if name == "ZZZ" {
				slog.Info("Got to ZZZ", "steps", steps)
				return
			}
		}
		slog.Info("Ran through all instructions")
	}
}

// part2:
func part2(input string) {
	panic("Unimplemented")
}
