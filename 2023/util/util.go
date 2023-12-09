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

func SlicesMap[S1 ~[]E1, S2 ~[]E2, E1, E2 any](s1 S1, mapFn func(E1) E2) S2 {
	s2 := make([]E2, len(s1))
	for i, e1 := range s1 {
		s2[i] = mapFn(e1)
	}
	return s2
}

func SlicesReduce[S ~[]E, E any](s S, f func(E, E) E, initVal E) E {
	acc := initVal
	for _, v := range s {
		acc = f(acc, v)
	}
	return acc
}
