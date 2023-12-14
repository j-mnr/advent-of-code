package twelve

import (
	"aoc/util"
	_ "embed"
	"log/slog"
	"strings"
	"strconv"
)

const (
	unknown = '?'
	damaged = '#'
	working = '.'
)

var (
	// example1:
	// ???.### 1,1,3
	// .??..??...?##. 1,1,3
	// ?#?#?#?#?#?#?#? 1,3,1,6
	// ????.#...#... 4,1,1
	// ????.######..#####. 1,6,5
	// ?###???????? 3,2,1
	example1 = strings.NewReader(`
???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1
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

type conditionRecord struct {
	record string
	groups []int
}

// part1:
func part1(input string) {
	var crs []conditionRecord
	crCache := map[string]int{}
	sum := 0
	for i, line := range strings.Split(input, "\n") {
		rNg := strings.Fields(line)
		crs = append(crs, conditionRecord{
			record: rNg[0],
			groups: util.SlicesMap[[]string, []int](
				strings.FieldsFunc(rNg[1], func(r rune) bool {
					return r == ','
				}),
				func(s string) int { return util.Must2(strconv.Atoi(s)) },
			),
		})

		possibilities := calc(crs[i])
		sum += possibilities
		slog.Info("Parsing", "record", crs[i].record, "groups", crs[i].groups,
			"possibilities", possibilities)
	}
	slog.Info("All possibilities", "result", sum)

	for _, cr := range crs {
		if _, found := crCache[cr.record]; found {
			slog.Error("FOUND SAME RECORD!!!!", "record", cr.record, "groups", cr.groups)
		}
		crCache[cr.record] = 0
	}
}

func calc(cr conditionRecord) int {
	processDot := func(cr conditionRecord) int {
		cr.record = cr.record[1:]
		slog.Info("processDot", "record", cr.record, "groups", cr.groups)
		return calc(cr)
	}

	processPound := func(cr conditionRecord) int {
		next := cr.groups[0]
		if len(cr.record) < next {
			return 0
		}
		grouping := strings.ReplaceAll(cr.record[:next], "?", "#")
		c := strings.Count(grouping, "#")
		slog.Info("processPound", "count", c, "next count", next,
			"returning 0", c != next)
		if c != next {
			return 0
		}

		if len(cr.record) == next {
			slog.Info("processPound", "record", cr.record, "groups", cr.groups)
			if len(cr.groups) == 1 {
				slog.Info("Obvious flaw", "cr", cr)
				return 1
			}
			return 0
		}

		slog.Info("processPound", "next rune", string(cr.record[next]),
			"groups", cr.groups)
		switch cr.record[next] {
		case unknown, working:
			cr.record = cr.record[next+1:]
			cr.groups = cr.groups[1:]
			slog.Info("processPound before Calc", "record", cr.record,
				"groups", cr.groups)
			return calc(cr)
		}
		return 0
	}

	if len(cr.groups) == 0 {
		// If rest are not damaged we count this as valid; empty is valid.
		if func(s string) bool {
			for _, r := range s {
				if r == damaged {
					return false
				}
			}
			return true
		}(cr.record) {
			slog.Info("Obvious flaw", "cr", cr)
			return 1
		}

		return 0
	}

	if len(cr.record) == 0 {
		return 0
	}

	slog.Info("Calc", "record", cr.record, "groups", cr.groups)
	switch cr.record[0] {
	case damaged:
		return processPound(cr)
	case working:
		return processDot(cr)
	case unknown:
		return processPound(cr) + processDot(cr)
	}
	panic("WHAT DA HELL IS DAT? " + string(cr.record[0]))
}

// part2:
func part2(input string) {
	panic("Unimplemented")
}
