package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"sort"
)

var test = []byte(`
FBFBBFFRLR
BFFFBBFRRR
FFFBBBFRRR
BBFFBBFRLL`[1:])

func main() {
	f, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}
	lines := bytes.Split(f, []byte("\n"))
	seats := make([]int, 0, len(lines))
	for _, line := range lines[:len(lines)-1] {
		lo, hi := 0, 127
		for _, r := range line[:7] {
			mid := lo + (hi-lo)/2
			switch r {
			case 'F':
				hi = mid - 1
			case 'B':
				lo = mid + 1
			}
		}
		row := lo
		lo, hi = 0, 7
		for _, r := range line[7:] {
			mid := lo + (hi-lo)/2
			switch r {
			case 'R':
				lo = mid + 1
			case 'L':
				hi = mid - 1
			}
		}
		col := lo
		seats = append(seats, row*8+col)
	}
	sort.Ints(seats)
	for prev, i := seats[0], 1; i < len(seats); prev, i = seats[i], i+1 {
		if seats[i] != prev+1 {
			fmt.Println("YOUR SEAT SIR:", prev+1)
			break
		}
	}
}
