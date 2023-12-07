package five

import (
	"aoc/util"
	_ "embed"
	"log/slog"
	"math"
	"slices"
	"strconv"
	"strings"
)

var (
	// example1: An almanac; the first line is the starting numbers we look at
	// seed-to-soil mapping.
	//
	// The first line has a destination range start of 50, a source range start of
	// 98, and a range length of 2. This line means that the source range starts
	// at 98 and contains two values: 98 and 99. The destination range is the same
	// length, but it starts at 50, so its two values are 50 and 51. With this
	// information, you know that seed number 98 corresponds to soil number 50 and
	// that seed number 99 corresponds to soil number 51.
	//
	// The second line means that the source range starts at 50 and contains 48
	// values: 50, 51, ..., 96, 97. This corresponds to a destination range
	// starting at 52 and also containing 48 values: 52, 53, ..., 98, 99. So, seed
	// number 53 corresponds to soil number 55.
	//
	// Any source numbers that aren't mapped correspond to the same destination
	// number. So, seed number 10 corresponds to soil number 10.
	example1 = strings.NewReader(`
seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4
`[1:])

	// example2:
	// The values on the initial seeds: line come in pairs. Within each pair, the
	// first value is the start of the range and the second value is the length of
	// the range. So, in the first line of the example above:
	//
	//	seeds: 79 14 55 13
	//
	// This line describes two ranges of seed numbers to be planted in the garden.
	// The first range starts with seed number 79 and contains 14 values: 79, 80,
	// ..., 91, 92. The second range starts with seed number 55 and contains 13
	// values: 55, 56, ..., 66, 67.
	//
	// In the above example, the lowest location number can be obtained from seed
	// number 82, which corresponds to soil 84, fertilizer 84, water 84, light 77,
	// temperature 45, humidity 46, and location 46. So, the lowest location
	// number is 46.
	example2 = strings.NewReader(`
seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4
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

// part1: Using these maps (of the almanac), find the lowest location number
// that corresponds to any of the initial seeds. To do this, you'll need to
// convert each seed number through other categories until you can find its
// corresponding location number.
func part1(input string) {
	almanac := strings.Split(input, "\n\n")
	var transitions []int
	for _, s := range strings.Fields(almanac[0])[1:] {
		transitions = append(transitions, util.Must2(strconv.Atoi(s)))
	}

	breakApart := func(line string) (dst, src, rng int) {
		f := strings.Fields(line)
		atoi := func(s string) int { return util.Must2(strconv.Atoi(s)) }
		return atoi(f[0]), atoi(f[1]), atoi(f[2])
	}
	for _, mapping := range almanac[1:] { // almanac[0] is seeds
		lines := strings.Split(mapping, "\n")
		slog.Info("Before "+lines[0], "transitions", transitions)
		for i, t := range transitions {
			for _, line := range lines[1:] {
				dst, src, rng := breakApart(line)
				if between := t - src; 0 <= between && between < rng {
					transitions[i] = between + dst
					slog.Info("A conversion was made",
						"destination start", dst, "source start", src, "range", rng,
						"old transition", t, "new transition", transitions[i],
					)
					break
				}
			}
		}
		slog.Info("After "+lines[0], "transitions", transitions)
	}
	slog.Info("Lowest location number",
		"result", slices.Min(transitions))
}

type almanacRange struct {
	start, end int64
	processed  bool
}

type table struct {
	name     string
	mappings []struct{ start, end, conversion int64 }
}

// part2: The values on the initial seeds: line come in pairs. Within each pair,
// the first value is the start of the range and the second value is the length
// of the range.
//
// Now, rather than considering four seed numbers, you need to consider a total
// of 27 seed numbers.
//
// Consider all of the initial seed numbers listed in the ranges on the first
// line of the almanac. What is the lowest location number that corresponds to
// any of the initial seed numbers?
//
// XXX: Doesn't work; finds lower number.
func part2(input string) {
	almanac := strings.Split(input, "\n\n")
	var seedRanges []*almanacRange
	ranges := strings.Fields(almanac[0])[1:]
	for i := 0; i < len(ranges); i += 2 {
		seed := util.Must2(strconv.Atoi(ranges[i]))
		seedRanges = append(seedRanges, &almanacRange{
			start: int64(seed),
			end:   int64(seed + util.Must2(strconv.Atoi(ranges[i+1])) - 1),
		})
	}
	slog.Info("Created ranges to check", "seedRanges", seedRanges)

	breakApart := func(line string) (dst, src, rng int64) {
		f := strings.Fields(line)
		atoi := func(s string) int64 { return int64(util.Must2(strconv.Atoi(s))) }
		return atoi(f[0]), atoi(f[1]), atoi(f[2])
	}
	var tables []table
	for i, almTable := range almanac[1:] {
		lines := strings.Split(almTable, "\n")
		tables = append(tables, table{
			name:     lines[0],
			mappings: make([]struct{ start, end, conversion int64 }, len(lines[1:])),
		})
		for j, line := range lines[1:] {
			dst, src, rng := breakApart(line)
			tables[i].mappings[j] = struct{ start, end, conversion int64 }{
				start:      src,
				end:        src + rng - 1,
				conversion: dst - src,
			}
		}
		slog.Info("Created Table", "name", tables[i].name, "mappings", tables[i].mappings)
	}

	transitions := make([]*almanacRange, len(seedRanges))
	copy(transitions, seedRanges)
	for _, t := range tables {
		transitions = convert(transitions, t)
	}

	low := int64(math.MaxInt64)
	for _, v := range transitions {
		if v.start < low {
			low = v.start
		}
	}
	slog.Info("Lowest location number", "result", low)
}

func convert(input []*almanacRange, t table) (output []*almanacRange) {
	for hasRemaining := true; hasRemaining; {
		hasRemaining = false
		for _, rng := range input {
			if rng.processed { // Valid only when we append remaining
				continue
			}
			// Check for fitting or overlapping values in mapping
			for _, m := range t.mappings {
				if rng.start >= m.start && rng.end <= m.end {
					rng.processed = true
					o := &almanacRange{
						start: rng.start + m.conversion,
						end:   rng.end + m.conversion,
					}
					output = append(output, o)
					slog.Info("Fit values", "mapping", m, "range", rng, "output", o)
					break
				}
				if rng.start >= m.start && rng.start <= m.end {
					rng.processed = true
					o := &almanacRange{
						start: rng.start + m.conversion,
						end:   m.end + m.conversion,
					}
					output = append(output, o)
					slog.Info("Overlap values", "mapping", m, "range", rng, "output", o)

					hasRemaining = true
					remaining := &almanacRange{
						start: m.end + 1,
						end:   rng.end,
					}
					input = append(input, remaining)
					slog.Info("Remainder values", "remaining", remaining)
					break
				}
			}
			if !rng.processed {
				slog.Info("This input couldn't find a mapping",
					"range", rng)
				rng.processed = true
				output = append(output, &almanacRange{
					start: rng.start,
					end:   rng.end,
				})
			}
		}
	}
	return output
}
