package four

import (
	"aoc/util"
	_ "embed"
	"strings"
	"log/slog"
	"math"
)

var (
	// example1: card 1 has five winning numbers (41, 48, 83, 86, and 17) and
	// eight numbers you have (83, 86, 6, 31, 17, 9, 48, and 53). Of the numbers
	// you have, four of them (48, 83, 17, and 86) are winning numbers! That means
	// card 1 is worth 8 points (1 for the first match, then doubled three times
	// for each of the three matches after the first).
	example1 = strings.NewReader(`
Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
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

// part1: As far as the Elf has been able to figure out, you have to figure out
// which of the numbers you have appear in the list of winning numbers. The
// first match makes the card worth one point and each match after the first
// doubles the point value of that card.
//
// How many points are they worth in total?
// NOTE: points for a card are math.Pow(2, m-1) where m is the number of matches
func part1(input string) {
	sum := 0
	winners := ""
	for _, line := range strings.Split(input, "\n") {
		card, nums, ok := strings.Cut(line, ":")
		winners, nums, ok = strings.Cut(nums, " |")
		slog.Info(card, "winners", winners, "nums", nums, "found", ok)

		matches := -1
		for i := 0; i < len(nums); i += 3 {
			slog.Info(card, "number", nums[i:i+3])
			if strings.Contains(winners, nums[i:i+3]) {
				slog.Info("Found matching number", "match", nums[i:i+3])
				matches++
			}
		}
		if matches >= 0 {
			sum += int(math.Pow(2, float64(matches)))
			slog.Info("Found at least 1 match", "matches", matches, "sum", sum)
		}
	}
	slog.Info("Total points", "result", sum)
}

// part2:
func part2(input string) {
	panic("Unimplemented")
}
