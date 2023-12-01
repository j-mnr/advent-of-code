package main

import (
	"bytes"
	_ "embed"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

var (
	// 12
	// 38
	// 15
	// 77
	// Should equal 142 all together.
	example = strings.NewReader(`
1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`[1:])

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
	exampleF := flag.Bool("example", false, "Run the example given from AoC.")
	flag.Parse()

	if !*infoF {
		programLevel.Set(slog.LevelError)
	}
	buf := &bytes.Buffer{}
	if *exampleF {
		buf = nil
	}

	part1(buf)
}

func part1(buf *bytes.Buffer) {
	sum := 0
	for _, line := range strings.Split(prepareInput(buf), "\n") {
		first := strings.IndexAny(line, "0123456789")
		last := strings.LastIndexAny(line, "0123456789")
		info.Info("digits found", "first", string(line[first:first+1]), "last", string(line[last:last+1]))
		sum += must2(strconv.Atoi(string(line[first:first+1]) + string(line[last:last+1])))

	}
	fmt.Println("Sum of lines is", sum)
}

func prepareInput(buf *bytes.Buffer) string {
	defer func() { info.Info("prepareInput", "input", buf.String()) }()

	if buf == nil {
		var buf bytes.Buffer
		example.WriteTo(&buf)
		return buf.String()
	}

	buf.Write(input)
	// Remove the trailing newline from the file.
	return buf.String()[:len(buf.String())-1]
}

func must2[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}
