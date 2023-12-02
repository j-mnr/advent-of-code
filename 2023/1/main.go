package main

import (
	"bytes"
	_ "embed"
	"flag"
	"fmt"
	"log/slog"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	// 12
	// 38
	// 15
	// 77
	// == 142
	example1 = `
1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`[1:]
	// 29
	// 83
	// 13
	// 24
	// 42
	// 14
	// 76
	// == 281
	example2 = `
two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`[1:]

	//go:embed input.txt
	input []byte

	info *slog.Logger

	programLevel = new(slog.LevelVar)
)

func init() {
	info = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     programLevel,
		AddSource: true,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			switch a.Key {
			case "level":
				return slog.Attr{}
			case "time":
				a.Value = slog.StringValue(a.Value.Time().Format("15:04:05.000"))
				return a
			case "source":
				s := a.Value.String()
				a.Value = slog.StringValue(strings.Replace(s[strings.LastIndex(s, "/")+1:strings.LastIndex(s, "}")-1], " ", ":", 1))
				return a
			default:
				return a
			}
		},
	}))
}

func main() {
	infoF := flag.Bool("info", false, "Set if you want information on running processes.")
	exampleF := flag.Int("example", 0, "Run the example given from AoC part 1 or 2. Defaults to input")
	partF := flag.Int("part", 1, "Run part 1 or 2. Defaults to 1")
	flag.Parse()

	if !*infoF {
		programLevel.Set(slog.LevelError)
	}

	switch *partF {
	case 1:
		part1(prepareInput(*exampleF))
	case 2:
		part2(prepareInput(*exampleF))
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
		info.Info("digits found", slog.String("first", string(line[first:first+1])),
			slog.String("last", string(line[last:last+1])))
		sum += must2(strconv.Atoi(string(line[first:first+1]) + string(line[last:last+1])))
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
			info.Info("bad digit", slog.String("string", d))
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
			info.Info("ints found", slog.String("ifirstInt", string(line[ifirstInt:ifirstInt+1])),
				slog.String("ilastInt", string(line[ilastInt:ilastInt+1])))
		}

		// See if part 1 has better indexes than digits as letters.
		if ifirstInt != -1 && ifirstInt < minIdx {
			minIdx, ndigitMin = ifirstInt, -1
		}
		if ilastInt > maxIdx {
			maxIdx, ndigitMax = ilastInt, -1
		}
		info.Info("min/max indexes", slog.Int("minIdx", minIdx),
			slog.Int("maxIdx", maxIdx),
			slog.String("encompasses", accountForDigits(line, minIdx, maxIdx, ndigitMax)),
		)

		// Do what problem asks.
		first, last := line[minIdx:minIdx+1], line[maxIdx:maxIdx+1]
		if ndigitMin != -1 {
			first = toNumber(digit(line[minIdx:minIdx+ndigitMin]))
		}
		if ndigitMax != -1 {
			last = toNumber(digit(line[maxIdx:maxIdx+ndigitMax]))
		}
		sum += must2(strconv.Atoi(first + last))
	}

	fmt.Println("Sum of lines is", sum)
}

func prepareInput(exampleCase int) string {
	var buf bytes.Buffer
	defer func() { info.Info("prepareInput", "input", buf.String()) }()

	info.Info("exampleCase chosen", "example", exampleCase)
	switch exampleCase {
	case 0:
		buf.Write(input)
	case 1:
		buf.WriteString(example1)
	case 2:
		buf.WriteString(example2)
	default:
		panic("Only 3 cases 0=input, 1=part1, 2=part2")
	}
	// Remove the trailing newline from the file.
	return buf.String()[:len(buf.String())-1]
}

func must2[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}
