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

var f = strings.NewReader(`
    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`[1:])

func main() {
	data, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	// data, err := io.ReadAll(f)
	// if err != nil {
	// 	panic(err)
	// }
	graphAndSequences := bytes.Split(data, []byte("\n\n"))
	var moves []*crane
	for _, seq := range bytes.Split(graphAndSequences[1], []byte("\n")) {
		c := &crane{}
		if err := c.UnmarshalText(seq); err != nil {
			continue
		}
		moves = append(moves, c)
	}
	// fmt.Println(moves)

	g := &graph{}
	err = g.UnmarshalText(graphAndSequences[0])
	// fmt.Println(g, err)
	shuffle(g, moves)
	fmt.Println(g)
}

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

func (s *stack) pop() byte {
	if s == nil {
		return 0
	}
	// fmt.Println(*s)
	b := (*s)[len(*s)-1]
	(*s) = (*s)[:len(*s)-1]
	return b
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
	fmt.Println(indexes)
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
	for col, stk := range *g {
		fmt.Println(col, stk)
	}
	// fmt.Printf("%s %s %s\n", (*g)[1], (*g)[2], (*g)[3])
	return nil
}

func (g graph) String() string {
	var sb strings.Builder
	ordered := make([]stack, len(g))
	for column, stk := range g {
		fmt.Println(column, stk)
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

func shuffle(g *graph, moves []*crane) {
	for _, m := range moves {
		fmt.Println(m)
		from, to := (*g)[m.from], (*g)[m.to]
		for i := uint8(0); i < m.move; i++ {
			to = append(to, from.pop())
		}
		(*g)[m.from], (*g)[m.to] = from, to
		// fmt.Println(g)
	}
}
