package main

import (
	"aoc/1"
	"aoc/2"
	"aoc/3"
	"aoc/4"
	"aoc/5"
	"aoc/6"
	"aoc/7"
	"aoc/8"
	"aoc/9"
	"flag"
	"log/slog"
	"os"
	"strconv"
)

var programLevel = new(slog.LevelVar)

type uint8Value uint8

func newUint8Value(val uint8, p *uint8) *uint8Value {
	*p = val
	return (*uint8Value)(p)
}

func (u uint8Value) String() string {
	return strconv.Itoa(int(u))
}

func (u *uint8Value) Set(s string) error {
	i, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	*u = uint8Value(i)
	return nil
}

func init() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     programLevel,
		AddSource: true,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			switch a.Key {
			case "level":
				return slog.Attr{}
			case "time":
				// a.Value = slog.StringValue(a.Value.Time().Format("15:04:05.000"))
				// return a
				return slog.Attr{}
			case "source":
				// s := a.Value.String()
				// a.Value = slog.StringValue(
				// 	strings.Replace(s[strings.LastIndex(s, "/")+1:strings.LastIndex(s, "}")-1], " ", ":", 1),
				// )
				return slog.Attr{}
			default:
				return a
			}
		},
	})))
}

func main() {
	var dayF, partF uint8
	infoF := flag.Bool("info", false, "Gives information on the running process")
	flag.Var(newUint8Value(1, &dayF), "day", "The day to run")
	flag.Var(newUint8Value(1, &partF), "part", "The part to run")
	exampleF := flag.Bool("example", false, "Run the example given from AoC")
	flag.Parse()

	if !*infoF {
		programLevel.Set(slog.LevelError)
	}
	switch dayF {
	case 1:
		one.Run(uint8(partF), *exampleF)
	case 2:
		two.Run(uint8(partF), *exampleF)
	case 3:
		three.Run(uint8(partF), *exampleF)
	case 4:
		four.Run(uint8(partF), *exampleF)
	case 5:
		five.Run(uint8(partF), *exampleF)
	case 6:
		six.Run(uint8(partF), *exampleF)
	case 7:
		seven.Run(uint8(partF), *exampleF)
	case 8:
		eight.Run(uint8(partF), *exampleF)
	case 9:
		nine.Run(uint8(partF), *exampleF)
	}
}
