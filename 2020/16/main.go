package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var f = []byte(`
class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12`[1:])

type rule struct {
	name    string
	minmaxs [2][2]int
}

func main() {
	f, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	notes := bytes.Split(f, []byte("\n\n"))
	_, notes[1], _ = bytes.Cut(notes[1], []byte("\n"))
	_, notes[2], _ = bytes.Cut(notes[2], []byte("\n"))

	rules := parseRules(notes[0])
	ticket := ticketValues(notes[1])
	_ = ticket
	otickets := [][]int{}
	for _, t := range bytes.Split(notes[2], []byte("\n")) {
		otickets = append(otickets, ticketValues(t))
	}
	// fmt.Println(otickets)
	// fmt.Println(rules)

	var invalids []int
	for _, tvals := range otickets {
		for _, n := range tvals {
			if invalidNumber(n, rules) {
				invalids = append(invalids, n)
			}
		}
	}
	fmt.Println(invalids)

	sum := 0
	for _, n := range invalids {
		sum += n
	}
	fmt.Println(sum)
}

func invalidNumber(n int, rules []rule) bool {
	for _, r := range rules {
		if r.minmaxs[0][0] <= n && n <= r.minmaxs[0][1] {
			return false
		} else if r.minmaxs[1][0] <= n && n <= r.minmaxs[1][1] {
			return false
		}
	}
	return true
}

func ticketValues(input []byte) (nums []int) {
	for _, d := range bytes.Split(input, []byte(",")) {
		n, err := strconv.Atoi(string(d))
		if err != nil {
			log.Println(err)
		}
		nums = append(nums, n)
	}
	return nums
}

func parseRules(input []byte) []rule {
	minmaxs := [][2]int{}
	rules := []rule{}
	for _, line := range bytes.Split(input, []byte("\n")) {
		m := regexp.MustCompile(`(\w+): (\d+)-(\d+) or (\d+)-(\d+)`).
			FindStringSubmatch(string(line))
		r := rule{name: m[1]}
		isMin := true
		mm := [2]int{}
		for _, n := range m[2:] {
			d, err := strconv.Atoi(n)
			if err != nil {
				log.Fatal(err)
			}
			if isMin {
				mm[0] = d
			} else {
				mm[1] = d
				minmaxs = append(minmaxs, mm)
			}
			isMin = !isMin
		}
		r.minmaxs[0], r.minmaxs[1] = minmaxs[0], minmaxs[1]
		minmaxs = minmaxs[:0]
		rules = append(rules, r)
	}
	return rules
}
