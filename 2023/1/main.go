package one

import (
	"aoc/util"
	"bytes"
	_ "embed"
	"fmt"
	"log/slog"
	"math"
	"strconv"
	"strings"
)

var (
	// example1 == 142
	example1 = strings.NewReader(
		"1abc2\n" + // 12
			"pqr3stu8vwx\n" + // 38
			"a1b2c3d4e5f\n" + // 15
			"treb7uchet", // 77
	)
	// example2 == 281
	example2 = strings.NewReader(
		"two1nine\n" + // 29
			"eightwothree\n" + // 83
			"abcone2threexyz\n" + // 13
			"xtwone3four\n" + // 24
			"4nineeightseven2\n" + // 42
			"zoneight234\n" + // 14
			"7pqrstsixteen", // 76
	)

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

// part1: On each line, the calibration value can be found by combining the
// first digit and the last digit (in that order) to form a single two-digit
// number.
func part1(buf string) {
	sum := 0
	for _, line := range strings.Split(buf, "\n") {
		first := strings.IndexAny(line, "123456789")
		last := strings.LastIndexAny(line, "123456789")
		slog.Info("digits found", slog.String("first", string(line[first:first+1])),
			slog.String("last", string(line[last:last+1])))
		sum += util.Must2(strconv.Atoi(string(line[first:first+1]) + string(line[last:last+1])))
	}
	fmt.Println("Sum of lines is", sum)
}

// part2: Your calculation isn't quite right. It looks like some of the digits
// are actually spelled out with letters: one, two, three, four, five, six,
// seven, eight, and nine also count as valid "digits".
func part2(buf string) {
	type digit string
	const (
		one   digit = "one"
		two   digit = "two"
		three digit = "three"
		four  digit = "four"
		five  digit = "five"
		six   digit = "six"
		seven digit = "seven"
		eight digit = "eight"
		nine  digit = "nine"
	)

	toNumber := func(d digit) string {
		switch d {
		case one:
			return "1"
		case two:
			return "2"
		case three:
			return "3"
		case four:
			return "4"
		case five:
			return "5"
		case six:
			return "6"
		case seven:
			return "7"
		case eight:
			return "8"
		case nine:
			return "9"
		default:
			slog.Info("bad digit", slog.String("string", string(d)))
			panic("bad digit")
		}
	}

	accountForDigits := func(line string, imin, imax, dmax int) string {
		if dmax == -1 {
			dmax = 0
		}
		return line[imin : imax+dmax]
	}

	sum := 0
	for _, line := range strings.Split(buf, "\n") {
		// Find first and last digits as letters.
		minIdx, maxIdx := math.MaxInt, -1
		ndigitMin, ndigitMax := -1, -1
		for _, d := range []digit{one, two, three, four, five, six, seven, eight, nine} {
			ifirstDigit := strings.Index(line, string(d))
			ilastDigit := strings.LastIndex(line, string(d))
			if ifirstDigit != -1 && ifirstDigit < minIdx {
				minIdx, ndigitMin = ifirstDigit, len(d)
			}
			if ilastDigit > maxIdx {
				maxIdx, ndigitMax = ilastDigit, len(d)
			}
		}

		// Do part 1.
		ifirstInt := strings.IndexAny(line, "123456789")
		ilastInt := strings.LastIndexAny(line, "123456789")
		if ifirstInt > 0 && ilastInt > 0 {
			slog.Info("ints found", slog.String("ifirstInt", string(line[ifirstInt:ifirstInt+1])),
				slog.String("ilastInt", string(line[ilastInt:ilastInt+1])))
		}

		// See if part 1 has better indexes than digits as letters.
		if ifirstInt != -1 && ifirstInt < minIdx {
			minIdx, ndigitMin = ifirstInt, -1
		}
		if ilastInt > maxIdx {
			maxIdx, ndigitMax = ilastInt, -1
		}
		slog.Info("min/max indexes", slog.Int("minIdx", minIdx),
			slog.Int("maxIdx", maxIdx),
			slog.String("encompasses", accountForDigits(line, minIdx, maxIdx, ndigitMax)),
		)

		// Do what problem asks.
		first, last := line[minIdx:minIdx+1], line[maxIdx:maxIdx+1]
		if ndigitMin != -1 {
			first = toNumber(digit(line[minIdx : minIdx+ndigitMin]))
		}
		if ndigitMax != -1 {
			last = toNumber(digit(line[maxIdx : maxIdx+ndigitMax]))
		}
		sum += util.Must2(strconv.Atoi(first + last))
	}

	fmt.Println("Sum of lines is", sum)
}
