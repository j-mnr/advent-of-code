package two

import (
	"aoc/util"
	"bytes"
	_ "embed"
	"log/slog"
	"strconv"
	"strings"
)

var (
	// 12 red, 13 green, 14 blue == 1+2+5 == 8
	example1 = strings.NewReader(`
Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
`[1:])

	example2 = strings.NewReader(``)

	//go:embed input.txt
	input []byte
)

func Run(part uint8, example bool) {
	switch part {
	case 1:
		if example {
			part1(util.PrepareInput(example1))
			return
		}
		part1(util.PrepareInput(bytes.NewReader(input)))
	case 2:
		if example {
			part2(util.PrepareInput(example2))
			return
		}
		part2(util.PrepareInput(bytes.NewReader(input)))
	}
}

// part1 Determine which games would have been possible if the bag had been
// loaded with only 12 red cubes, 13 green cubes, and 14 blue cubes. What is the
// sum of the IDs of those games?
func part1(input string) {
	sum := 0
	for id, line := range strings.Split(input, "\n") {
		id := id + 1
		sum += id
		slog.Info("Current Game", "id", id, "line", line, "sum", sum)
		sum = findImpossibleGameIn(line, id, sum)
	}
	slog.Info("Reached end", "sum", sum)
}

func findImpossibleGameIn(line string, gameID, sum int) (newSum int) {
	const (
		red, green, blue       = " red", " green", " blue"
		rlimit, glimit, blimit = 12, 13, 14
	)
	hasSetSum := func(limit, rgblimit int, color, rgbcolor string) bool {
		if strings.Contains(color, rgbcolor) && limit > rgblimit {
			slog.Info("Limit too great for"+rgbcolor, "limit", limit, "color", color)
			sum -= gameID
			return true
		}
		return false
	}
	for i, set := range strings.Split(line[strings.Index(line, ": ")+2:], "; ") {
		slog.Info("Current Set", "set", set, "set #", i+1)
		for j, rgb := range strings.SplitN(set, ", ", 3) {
			slog.Info("Current Color", "gameID", gameID, "rgb #", j+1, "color", rgb)
			for _, color := range []struct {
				rgblimit int
				color    string
			}{
				{rlimit, red}, {glimit, green}, {blimit, blue},
			} {
				if hasSetSum(
					util.Must2(strconv.Atoi(rgb[:strings.Index(rgb, " ")])),
					color.rgblimit,
					rgb,
					color.color,
				) {
					return sum
				}
			}
		}
		slog.Info("After findImpossibleGameIn", "sum", sum)
	}
	return sum
}

func part2(input string) {
	panic("Unimplemented")
}
