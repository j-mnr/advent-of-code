package four

import (
	"aoc/util"
	_ "embed"
	"log/slog"
	"math"
	"strings"
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
	//
	// - Card 1 has four matching numbers, so you win one copy each of the next
	// four cards: cards 2, 3, 4, and 5.
	// - Your original card 2 has two matching numbers, so you win one copy each
	// of cards 3 and 4.
	// - Your copy of card 2 also wins one copy each of cards 3 and 4.
	// - Your four instances of card 3 (one original and three copies) have two
	// matching numbers, so you win four copies each of cards 4 and 5.
	// - Your eight instances of card 4 (one original and seven copies) have one
	// matching number, so you win eight copies of card 5.
	// - Your fourteen instances of card 5 (one original and thirteen copies) have
	// no matching numbers and win no more cards.
	// - Your one instance of card 6 (one original) has no matching numbers and
	// wins no more cards.
	example2 = strings.NewReader(`
Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
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

// part2: There's no such thing as "points". Instead, scratchcards only cause
// you to win more scratchcards equal to the number of winning numbers you have.
// Process all of the original and copied scratchcards until no more
// scratchcards are won. Including the original set of scratchcards, how many
// total scratchcards do you end up with?
func part2(input string) {
	cards := strings.Split(input, "\n")
	type cardMatch struct{ matches, copies uint }
	cm := make([]cardMatch, len(cards))
	winners := ""
	for i, line := range cards {
		card, nums, ok := strings.Cut(line, ":")
		winners, nums, ok = strings.Cut(nums, " |")
		slog.Info(card, "winners", winners, "nums", nums, "found", ok)

		matches := 0
		for i := 0; i < len(nums); i += 3 {
			slog.Info(card, "number", nums[i:i+3])
			if strings.Contains(winners, nums[i:i+3]) {
				slog.Info("Found matching number", "match", nums[i:i+3])
				matches++
			}
		}
		cm[i] = cardMatch{matches: uint(matches), copies: 1}
	}

	sum := uint(0)
	for i, v := range cm {
		sum += v.copies
		slog.Info("Adding copies", "card", i+1, "copies", v.copies, "matches", v.matches)
		if v.matches == 0 {
			continue
		}
		for n := uint(0); n < v.copies; n++ {
			for addCopy := uint(1); addCopy <= v.matches; addCopy++ {
				cm[uint(i)+addCopy].copies++
			}
		}
	}
	slog.Info("Total copies", "result", sum)
}
