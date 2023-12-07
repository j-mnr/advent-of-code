package six

import (
	"aoc/util"
	_ "embed"
	"log/slog"
	"strings"
	"strconv"
)

var (
	// example1: This document describes three races:
	//
	//	The first race lasts 7 milliseconds. The record distance in this race is 9
	//	millimeters.
	//	The second race lasts 15 milliseconds. The record distance in this race is
	//	40 millimeters.
	//	The third race lasts 30 milliseconds. The record distance in this race is
	//	200 millimeters.
	// res:  4 * 8 * 9 == 288
	example1 = strings.NewReader(`
Time:      7  15   30
Distance:  9  40  200
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

// part1: Your toy boat has a starting speed of zero millimeters per
// millisecond. For each whole millisecond you spend at the beginning of the
// race holding down the button, the boat's speed increases by one millimeter
// per millisecond.
//
// Determine the number of ways you could beat the record in each race. What do
// you get if you multiply these numbers together?
func part1(input string) {
	tnd := strings.Split(input, "\n")
	times := slicesMap[[]string, []int](strings.Fields(tnd[0])[1:], func(s string) int {
		return util.Must2(strconv.Atoi(s))
	})
	dists := slicesMap[[]string, []int](strings.Fields(tnd[1])[1:], func(s string) int {
		return util.Must2(strconv.Atoi(s))
	})
	slog.Info("Starting values", "times", times, "distances", dists)

	wins := make([]int, len(times))
	for i, t := range times {
		for j := 0; j < t+1; j++ {
			win := (t - j) * j
			if win > dists[i] {
				slog.Info("Found winning speed", "win", win)
				wins[i]++
			}
		}
	}

	product := 1
	for _, w := range wins {
		product *= w
	}
	slog.Info("Calculated wins", "wins", wins, "result", product)
}

func slicesMap[S1 ~[]E1, S2 ~[]E2, E1, E2 any](s1 S1, mapFn func(E1) E2) S2 {
	s2 := make([]E2, len(s1))
	for i, e1 := range s1 {
		s2[i] = mapFn(e1)
	}
	return s2
}

// part2:
func part2(input string) {
	panic("Unimplemented")
}
