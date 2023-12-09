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

	// example2: Here, there are two starting nodes, 11A and 22A (because they
	// both end with A). As you follow each left/right instruction, use that
	// instruction to simultaneously navigate away from both nodes you're
	// currently on. Repeat this process until all of the nodes you're currently
	// on end with Z. (If only some of the nodes you're on end with Z, they act
	// like any other node and you continue as normal.) In this example, you would
	// proceed as follows:
	//
	// - Step 0: You are at 11A and 22A.
	// - Step 1: You choose all of the left paths, leading you to 11B and 22B.
	// - Step 2: You choose all of the right paths, leading you to 11Z and 22C.
	// - Step 3: You choose all of the left paths, leading you to 11B and 22Z.
	// - Step 4: You choose all of the right paths, leading you to 11Z and 22B.
	// - Step 5: You choose all of the left paths, leading you to 11B and 22C.
	// - Step 6: You choose all of the right paths, leading you to 11Z and 22Z.
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

// part2: Simultaneously start on every node that ends with A. How many steps
// does it take before you're only on nodes that end with Z?
func part2(input string) {
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

	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}

	lcm := func(a, b int) int {
		return (a * b) / gcd(a, b)
	}

	findSteps := func(instructions, name string, nodes map[string]node) int {
		steps := 0
		for {
			for _, dir := range instructions {
				switch dir {
				case 'R':
					name = nodes[name].right
				case 'L':
					name = nodes[name].left
				default:
					panic("Impossible input")
				}
				steps++
				if strings.HasSuffix(name, "Z") {
					slog.Info("Got to '..Z'", "steps", steps)
					return steps
				}
			}
		}
	}

	var allSteps []int
	for name := range nodes {
		if !strings.HasSuffix(name, "A") {
			continue
		}
		allSteps = append(allSteps, findSteps(instructions, name, nodes))
		slog.Info("Gathered another set of steps", "node", name,
			"Steps so far", allSteps)
	}
	slog.Info("Reduced steps",
		"least common multiple", util.SlicesReduce(allSteps, lcm, 1))
}
