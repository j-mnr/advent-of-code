package seven

import (
	"aoc/util"
	_ "embed"
	"fmt"
	"log/slog"
	"slices"
	"strconv"
	"strings"
)

var (
	// example1: Each hand is followed by its bid
	// amount. Each hand wins an amount equal to its bid multiplied by its rank,
	// where the weakest hand gets rank 1, the second-weakest hand gets rank 2,
	// and so on up to the strongest hand. Because there are five hands in this
	// example, the strongest hand will have rank 5 and its bid will be multiplied
	// by 5.
	//
	// So, the first step is to put the hands in order of strength:
	//
	// - 32T3K is the only one pair and the other hands are all a stronger type,
	// so it gets rank 1.
	// - KK677 and KTJJT are both two pair. Their first cards both have the same
	// label, but the second card of KK677 is stronger (K vs T), so KTJJT gets
	// rank 2 and KK677 gets rank 3.
	// - T55J5 and QQQJA are both three of a kind. QQQJA has a stronger first
	// card, so it gets rank 5 and T55J5 gets rank 4.
	example1 = strings.NewReader(`
32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
`[1:])

	// example2:
	// - 32T3K is still the only one pair; it doesn't contain any jokers, so its
	// strength doesn't increase.
	// - KK677 is now the only two pair, making it the second-weakest hand.
	// - T55J5, KTJJT, and QQQJA are now all four of a kind! T55J5 gets rank 3,
	// QQQJA gets rank 4, and KTJJT gets rank 5.
	//
	// With the new joker rule, the total winnings in this example are 5905.
	//
	// 4 or greater 'J's is five of a kind.
	// Make sure to account for 3 'J's.
	// Anytime you have 1 'J' and 1 pair you have three of a kind NOT two pair.
	example2 = strings.NewReader(`
32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
JJJJJ 500
2JJJJ 500
8ATTJ 500
88TTJ 500
8JTTJ 500
8ATTJ 500
8AJJJ 500
AAJJJ 500
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

type handRank uint8

const (
	highCard handRank = iota
	onePair
	twoPair
	threeOfKind
	fullHouse
	fourOfKind
	fiveOfKind
)

func (h handRank) String() string {
	switch h {
	case highCard:
		return "highCard"
	case onePair:
		return "onePair"
	case twoPair:
		return "twoPair"
	case threeOfKind:
		return "threeOfKind"
	case fullHouse:
		return "fullHouse"
	case fourOfKind:
		return "fourOfKind"
	case fiveOfKind:
		return "fiveOfKind"
	default:
		panic("Impossible hand rank")
	}
}

type cardRank uint8

// In part 2 order now; cardJ goes after cardT in part 1.
const (
	cardJ cardRank = iota
	card2
	card3
	card4
	card5
	card6
	card7
	card8
	card9
	cardT
	cardQ
	cardK
	cardA
)

func newCardRank(card byte) cardRank {
	switch card {
	case '2':
		return card2
	case '3':
		return card3
	case '4':
		return card4
	case '5':
		return card5
	case '6':
		return card6
	case '7':
		return card7
	case '8':
		return card8
	case '9':
		return card9
	case 'T':
		return cardT
	case 'J':
		return cardJ
	case 'Q':
		return cardQ
	case 'K':
		return cardK
	case 'A':
		return cardA
	}
	panic("Impossible card value")
}

func (c cardRank) String() string {
	switch c {
	case cardJ:
		return "J"
	case card2:
		return "2"
	case card3:
		return "3"
	case card4:
		return "4"
	case card5:
		return "5"
	case card6:
		return "6"
	case card7:
		return "7"
	case card8:
		return "8"
	case card9:
		return "9"
	case cardT:
		return "T"
	case cardQ:
		return "Q"
	case cardK:
		return "K"
	case cardA:
		return "A"
	}
	panic("Impossible card value")
}

type round struct {
	hand string
	bet  int
}

func cmpRound(rankHand func(string) handRank) func(a, b round) int {
	return func(a, b round) int {
		// find higher ranking card
		r1, r2 := rankHand(a.hand), rankHand(b.hand)
		switch {
		case r1 < r2:
			return -1
		case r1 > r2:
			return 1
		default: // Same rank; find high card
			for i := range a.hand {
				c1, c2 := newCardRank(a.hand[i]), newCardRank(b.hand[i])
				// slog.Info("Card ranks", "card1", c1, "rank1", c1,
				// 	"card2", c2, "rank2", c2)
				switch {
				case c1 < c2:
					return -1
				case c1 > c2:
					return 1
				default: // Move to next card or equal hands in every way
				}
			}
		}
		return 0
	}
}

// part1: Find the rank of every hand in your set. What are the total winnings?
func part1(input string) {
	var rounds []round
	for _, raw := range strings.Split(input, "\n") {
		handBet := strings.Fields(raw)
		rounds = append(rounds, round{
			hand: handBet[0],
			bet:  util.Must2(strconv.Atoi(handBet[1])),
		})
	}
	rankHand := func(hand string) handRank {
		counter := map[rune]uint8{}
		for _, card := range hand {
			counter[card]++
		}
		var has3oK, has1pair bool
		for _, count := range counter {
			switch count {
			case 5:
				return fiveOfKind
			case 4:
				return fourOfKind
			case 3:
				has3oK = true
			case 2:
				if has1pair {
					return twoPair
				}
				has1pair = true
			case 1: // nop
			}
		}
		if has3oK || has1pair {
			switch {
			case has3oK && has1pair:
				return fullHouse
			case has3oK:
				return threeOfKind
			case has1pair:
				return onePair
			}
		}
		return highCard
	}
	slices.SortFunc(rounds, cmpRound(rankHand))
	slog.Info("Converted values", "rounds", rounds)
	total := 0
	for i, r := range rounds {
		total += (r.bet * (i + 1))
	}
	slog.Info("Total", "result", total)
}

// part2: Using the new joker rule, find the rank of every hand in your set.
// What are the new total winnings?
func part2(input string) {
	var rounds []round
	for _, raw := range strings.Split(input, "\n") {
		handBet := strings.Fields(raw)
		rounds = append(rounds, round{
			hand: handBet[0],
			bet:  util.Must2(strconv.Atoi(handBet[1])),
		})
	}
	rankHand := func(hand string) handRank {
		counter := map[rune]uint8{}
		highestCard, highestCount := rune('X'), uint8(0)
		for _, card := range hand {
			counter[card]++
			if counter[card] > highestCount && card != 'J' {
				highestCard, highestCount = card, counter[card]
			}
		}

		jokersWild := counter['J']
		if jokersWild == 5 || jokersWild == 4 {
			return fiveOfKind
		}

		var has3oK, has1pair bool
		switch highestCount + jokersWild {
		case 5:
			return fiveOfKind
		case 4:
			return fourOfKind
		case 3:
			has3oK = true
		case 2:
			if has1pair {
				return twoPair
			}
			has1pair = true
		case 1: // nop
		default:
			slog.Error("Impossible score for: " + string(highestCard) + " " +
				strconv.Itoa(int(highestCount)))
		}
		for card, count := range counter {
			if card == highestCard || card == 'J' {
				continue
			}
			switch count {
			case 5:
				return fiveOfKind
			case 4:
				return fourOfKind
			case 3:
				has3oK = true
			case 2:
				if has1pair {
					return twoPair
				}
				has1pair = true
			case 1: // nop
			default:
				panic("Impossible score for: " + string(card) + " " + strconv.Itoa(int(count)))
			}
		}
		if has3oK || has1pair {
			switch {
			case has3oK && has1pair:
				return fullHouse
			case has3oK:
				return threeOfKind
			case has1pair:
				return onePair
			}
		}
		return highCard
	}
	slices.SortFunc(rounds, cmpRound(rankHand))
	for i, r := range rounds {
		slog.Info("Round played", "ranking", fmt.Sprintf("%03d", i+1),
			"hand", r.hand, "bet", fmt.Sprintf("%03d", r.bet), "rank", rankHand(r.hand))
	}
	total := 0
	for i, r := range rounds {
		total += (r.bet * (i + 1))
	}
	slog.Info("Total", "result", total)
}
