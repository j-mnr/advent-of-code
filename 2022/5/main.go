package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	errMalformedText = errors.New("could not unmarshal text into crane")
	errNilGraph      = errors.New("Cannot unmarshal text into a nil graph")
)

type crane struct{ move, from, to uint8 }

func (c *crane) UnmarshalText(text []byte) error {
	const move, from, to = "move ", "from ", "to "

	midx := bytes.Index(text, []byte(move))
	fidx := bytes.Index(text, []byte(from))
	tidx := bytes.Index(text, []byte(to))
	if midx == -1 || tidx == -1 || fidx == -1 {
		return errMalformedText
	}
	midx = midx + len(move)
	mnum, err := strconv.Atoi(string(text[midx : fidx-1]))
	if err != nil {
		return fmt.Errorf("crane: %w", err)
	}
	fidx = fidx + len(from)
	fnum, err := strconv.Atoi(string(text[fidx : tidx-1]))
	if err != nil {
		return fmt.Errorf("crane: %w", err)
	}
	tidx = tidx + len(to)
	tnum, err := strconv.Atoi(string(text[tidx:]))
	if err != nil {
		return fmt.Errorf("crane: %w", err)
	}
	*c = crane{
		move: uint8(mnum),
		from: uint8(fnum),
		to:   uint8(tnum),
	}
	return nil
}

func (c crane) String() string {
	return fmt.Sprintf("\"move %d from %d to %d\"", c.move, c.from, c.to)
}

type stack []byte

func (s *stack) pop(amount uint8) []byte {
	if s == nil {
		return nil
	}
	b := make(stack, amount)
	for i, j := len(*s)-1, 1; i >= len(*s)-int(amount); i, j = i-1, j+1 {
		b[len(b)-j] = (*s)[i]
	}
	(*s) = (*s)[:len(*s)-int(amount)]
	return b
}

func (s stack) String() string {
	var sb strings.Builder
	sb.WriteByte('[')
	for _, b := range s {
		sb.Write([]byte{' ', b})
	}
	sb.Write([]byte(" ]"))
	return sb.String()
}

type graph map[uint8]stack

func (g *graph) UnmarshalText(text []byte) error {
	if g == nil {
		return errNilGraph
	}
	rows := bytes.Split(text, []byte("\n"))

	// gather column indexes
	indexes := make(map[int]uint8)
	column := uint8(1)
	for i, b := range rows[len(rows)-1] {
		if !(b >= '1' && b <= '9') {
			continue
		}
		indexes[i] = column
		column++
	}
	rows = rows[:len(rows)-1]
	*g = make(graph, len(indexes))

	// gather crates by name
	for i := len(rows) - 1; i >= 0; i-- {
		for i, b := range rows[i] {
			if !(b >= 'A' && b <= 'Z') {
				continue
			}
			(*g)[indexes[i]] = append((*g)[indexes[i]], b)
		}
	}
	return nil
}

func (g graph) String() string {
	var sb strings.Builder
	ordered := make([]stack, len(g))
	for column, stk := range g {
		ordered[column-1] = stk
	}

	for i, stk := range ordered {
		sb.WriteString(strconv.Itoa(i+1) + ":")
		for i := len(stk) - 1; i >= 0; i-- {
			sb.WriteByte(' ')
			sb.WriteByte(stk[i])
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	graphAndSequences := bytes.Split(data, []byte("\n\n"))
	var moves []*crane
	for _, seq := range bytes.Split(graphAndSequences[1], []byte("\n")) {
		c := &crane{}
		if err := c.UnmarshalText(seq); err != nil {
			continue
		}
		moves = append(moves, c)
	}

	g := &graph{}
	err = g.UnmarshalText(graphAndSequences[0])
	shuffle(g, moves)

	fmt.Println(g)
}

func shuffle(g *graph, moves []*crane) {
	for _, m := range moves {
		from, to := (*g)[m.from], (*g)[m.to]
		to = append(to, from.pop(m.move)...)
		(*g)[m.from], (*g)[m.to] = from, to
	}
}
