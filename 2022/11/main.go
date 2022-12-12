package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

var f = strings.NewReader(`
Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3

Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0

Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

Monkey 3:
  Starting items: 74
  Operation: new = old + 3
  Test: divisible by 17
    If true: throw to monkey 0
		If false: throw to monkey 1`[1:])

type monkey struct {
	items     []uint64
	inspect   func(uint64) uint64
	throw     func(uint64) int
	inspected uint64
}

func main() {
	data, err := io.ReadAll(f)
	data, err = os.ReadFile("input")
	check(err)

	lcm := 1
	monkeys := populate(data, &lcm)

	for i := 0; i < 10000; i++ {
		for j, m := range monkeys {
			for _, it := range m.items {
				newItem := m.inspect(it) % uint64(lcm)
				monkeys[j].inspected++
				to := m.throw(newItem)
				monkeys[to].items = append(monkeys[to].items, newItem)
			}
			monkeys[j].items = m.items[:0]
		}
	}

	first, second := uint64(0), uint64(0)
	for i, m := range monkeys {
		ins := m.inspected
		switch {
		case ins > first:
			first, second = ins, first
		case ins > second:
			second = ins
		}
		fmt.Println("monkey", i, "inspected", m.inspected, "items")
	}

	fmt.Println("Level of monkey business:", first*second)
}

func populate(data []byte, lcm *int) []monkey {
	var monkeys []monkey
	for _, mData := range strings.Split(string(data), "\n\n") {
		m := monkey{}
		details := strings.Split(mData, "\n")

		// gather items
		_, items, _ := strings.Cut(details[1], ": ")
		for _, item := range strings.Split(items, ", ") {
			n, err := strconv.Atoi(item)
			check(err)
			m.items = append(m.items, uint64(n))
		}

		// create inspection function
		_, fn, _ := strings.Cut(details[2], "new = old ")
		f := strings.Fields(fn)
		m.inspect = func(u uint64) uint64 {
			n, err := strconv.Atoi(f[1])
			if err != nil {
				n = int(u)
			}
			switch f[0] {
			case "+":
				return (u + uint64(n))
			case "*":
				return (u * uint64(n))
			default:
				panic("Not implemented!")
			}
		}

		// create test function
		_, num, _ := strings.Cut(details[3], ": divisible by ")
		n, err := strconv.Atoi(num)
		// Part 2
		*lcm *= n
		check(err)
		_, num, _ = strings.Cut(details[4], ": throw to monkey ")
		m1, err := strconv.Atoi(num)
		check(err)
		_, num, _ = strings.Cut(details[5], ": throw to monkey ")
		m2, err := strconv.Atoi(num)
		check(err)
		m.throw = func(u uint64) int {
			if u%uint64(n) == 0 {
				return m1
			}
			return m2
		}

		monkeys = append(monkeys, m)
	}
	return monkeys
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
