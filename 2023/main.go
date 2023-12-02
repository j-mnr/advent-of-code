package main

import (
	"aoc/1"
	"flag"
	"os"
	"strings"
	"strconv"
	"log/slog"
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
				a.Value = slog.StringValue(a.Value.Time().Format("15:04:05.000"))
				return a
			case "source":
				s := a.Value.String()
				a.Value = slog.StringValue(
					strings.Replace(s[strings.LastIndex(s, "/")+1:strings.LastIndex(s, "}")-1], " ", ":", 1),
				)
				return a
			default:
				return a
			}
		},
	})))
}

func main() {
	var dayF, partF, exampleF uint8
	infoF := flag.Bool("info", false, "Set if you want information on running processes.")
	flag.Var(newUint8Value(1, &dayF), "day", "The day to run")
	flag.Var(newUint8Value(1, &partF), "part", "The part to run")
	flag.Var(newUint8Value(0, &exampleF), "example", "Run the example given from AoC part 1 or 2. Defaults to input")
	flag.Parse()

	if !*infoF {
		programLevel.Set(slog.LevelError)
	}
	switch dayF {
	case 1:
		one.Run(uint8(partF), uint8(exampleF))
	}
}
