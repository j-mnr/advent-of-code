package main

import (
	"bufio"
	"fmt"
	"log"
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
		log.Fatal(err)
	}
	defer f.Close()
	var items []item
	scr := bufio.NewScanner(f)
	for scr.Scan() {
		n := len(scr.Text()) / 2
		first, second := scr.Text()[:n], scr.Text()[n:]

	out:
		for _, b1 := range first {
			for _, b2 := range second {
				if b1 == b2 {
					items = append(items, item(b1))
					break out
				}
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
