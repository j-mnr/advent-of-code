package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
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
	max := 0
	lines := bytes.Split(f, []byte("\n"))
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
		if val := row*8 + col; max < val {
			max = val
		}
	}
	fmt.Println(max)
}
