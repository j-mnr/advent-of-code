package util

import (
	"bytes"
	"io"
	"log/slog"
)

func PrepareInput(r io.Reader) string {
	var buf bytes.Buffer
	defer func() { slog.Info("prepareInput", "input", buf.String()) }()

	b, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}

	buf.Write(b)
	return buf.String()[:len(buf.String())-1]
}

func Must2[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}
