package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

var f = strings.NewReader(`
vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`[1:])

type item byte

func (i item) String() string {
	return string(i)
}

func (i item) priority() uint8 {
	switch {
	case 'a' <= i && i <= 'z':
		return uint8(i - 'a' + 1)
	case 'A' <= i && i <= 'Z':
		return uint8(i - 'A' + 27)
	}
	panic("Should be unreachable")
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}
	var groups [][3]map[item]struct{}
	fields := bytes.Fields(b)
	for i := 0; i < len(fields); i += 3 {
		groups = append(groups,
			[3]map[item]struct{}{itemize(fields[i]), itemize(fields[i+1]), itemize(fields[i+2])})
	}

	var items []item
	for _, group := range groups {
		elf := 0
		for item := range group[elf] {
			_, ok1 := group[elf+1][item]
			_, ok2 := group[elf+2][item]
			if ok1 && ok2 {
				items = append(items, item)
			}
		}
	}

	sum := uint32(0)
	for _, it := range items {
		sum += uint32(it.priority())
	}
	fmt.Printf("items: %v\n", items)
	fmt.Printf("sum: %v\n", sum)
}

func itemize(data []byte) map[item]struct{} {
	items := make(map[item]struct{}, len(data))
	for _, b := range data {
		items[item(b)] = struct{}{}
	}
	return items
}
