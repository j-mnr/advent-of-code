package two

import (
	"aoc/util"
	_ "embed"
	"log/slog"
	"strings"
)

var (
	// example1: Scanning box IDs you need to make a checksum
	// - abcdef contains no letters that appear exactly two or three times.
	// - bababc contains two a and three b, so it counts for both.
	// - abbcde contains two b, but no letter appears exactly three times.
	// - abcccd contains three c, but no letter appears exactly two times.
	// - aabcdd contains two a and two d, but it only counts once.
	// - abcdee contains two e.
	// - ababab contains three a and three b, but it only counts once.
	example1 = strings.NewReader(`
abcdef
bababc
abbcde
abcccd
aabcdd
abcdee
ababab
`[1:])

	// example2: The IDs abcde and axcye are close, but they differ by two
	// characters (the second and fourth). However, the IDs fghij and fguij differ
	// by exactly one character, the third (h and u). Those must be the correct
	// boxes.
	example2 = strings.NewReader(`
abcde
fghij
klmno
pqrst
fguij
axcye
wvxyz
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

// part1: You use your fancy wrist device to quickly scan every box's ID and
// produce a list of the likely candidates.
//
//	checksumFn == ntwoCounts * nthreeCounts
//
// What is the checksum for the list?
func part1(input string) {
	var ntwoCounts, nthreeCounts uint
	for _, id := range strings.Split(input, "\n") {
		counter := map[rune]uint8{}
		for _, r := range id {
			counter[r]++
		}
		slog.Info("Counts for box ID", "ID", id, "counter", counter)
		var hasTwo, hasThree bool
		for _, v := range counter {
			switch {
			case v == 2 && !hasTwo:
				hasTwo = true
				ntwoCounts++
			case v == 3 && !hasThree:
				hasThree = true
				nthreeCounts++
			}
			if hasTwo && hasThree {
				break
			}
		}
	}
	slog.Info("Total", "amount of two counts", ntwoCounts,
		"amount of three counts", nthreeCounts, "result", ntwoCounts*nthreeCounts)
}

// part2: The boxes will have IDs which differ by exactly one character at the
// same position in both strings.
//
// What letters are common between the two correct box IDs? (In the example
// above, this is found by removing the differing character from either ID,
// producing fgij.)
func part2(input string) {
	ids := strings.Split(input, "\n")
	var boxID1, boxID2 string
foundIDs:
	for i, id1 := range ids {
		for _, id2 := range ids[i+1:] {
			var oneDiff bool
		nextCandidate:
			for i, r := range id1 {
				switch {
				case !oneDiff && r != rune(id2[i]):
					oneDiff = true // One more chance!
				case r != rune(id2[i]):
					break nextCandidate
				}
				if oneDiff && i == len(id1)-1 {
					boxID1, boxID2 = id1, id2
					break foundIDs
				}
			}
		}
	}
	var sb strings.Builder
	for i := range boxID1 {
		if boxID1[i] != boxID2[i] {
			continue
		}
		sb.WriteByte(boxID1[i])
	}
	slog.Info("Found two IDs", "one", boxID1, "two", boxID2, "result", sb.String())
}
