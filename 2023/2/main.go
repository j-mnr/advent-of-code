package two

import (
	"aoc/util"
	_ "embed"
	"log/slog"
	"strconv"
	"strings"
)

const red, green, blue = " red", " green", " blue"

var (
	// part1: 12 red, 13 green, 14 blue == 1+2+5 == 8
	// part2: (Game 1: 4 red * 2 green * 6 blue)48 + 2 + 1560 + 630 + 36 == 2286
	example1 = strings.NewReader(`
Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
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
			data = util.PrepareInput(example1)
		}
		part2(data)
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
	const rlimit, glimit, blimit = 12, 13, 14
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

// part2 For each game, find the minimum set of cubes that must have been
// present. What is the sum of the power of these sets?
//
// The power of a set of cubes is equal to the numbers of red, green, and blue
// cubes multiplied together.
func part2(input string) {
	sum := 0
	for id, line := range strings.Split(input, "\n") {
		id := id + 1
		slog.Info("Current Game", "id", id, "line", line, "sum", sum)

		minSet := struct{ r, g, b uint }{}
		for i, set := range strings.Split(line[strings.Index(line, ": ")+2:], "; ") {
			slog.Info("Current Set", "set", set, "set #", i+1)
			findLowestSetPossible(set, &minSet)
			slog.Info("After findLowestSetPossible", "minSet", minSet, "sum", sum)
		}

		sum += int(minSet.r * minSet.g * minSet.b)
	}
	slog.Info("Reached end", "sum", sum)
}

func findLowestSetPossible(set string, minSet *struct{ r, g, b uint }) {
	for j, rgb := range strings.SplitN(set, ", ", 3) {
		slog.Info("Current Color", "rgb #", j+1, "color", rgb)
		cand := uint(util.Must2(strconv.Atoi(rgb[:strings.Index(rgb, " ")])))
		switch {
		case strings.Contains(rgb, red):
			if cand > minSet.r {
				slog.Info("Setting lower value", red, minSet.r, "color", rgb,
					"candidate", cand)
				minSet.r = cand
			}
		case strings.Contains(rgb, green):
			if cand > minSet.g {
				slog.Info("Setting lower value", green, minSet.g, "color", rgb,
					"candidate", cand)
				minSet.g = cand
			}
		case strings.Contains(rgb, blue):
			if cand > minSet.b {
				slog.Info("Setting lower value", blue, minSet.b, "color", rgb,
					"candidate", cand)
				minSet.b = cand
			}
		}
	}
}
