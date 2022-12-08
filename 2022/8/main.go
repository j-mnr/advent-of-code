package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var f = strings.NewReader(`
30373
25512
65332
33549
35390`[1:])

type height uint8

type direction uint8

const (
	unknown = iota
	left
	right
	down
	up
)

func (d direction) String() string {
	switch d {
	case left:
		return "left"
	case right:
		return "right"
	case up:
		return "up"
	case down:
		return "down"
	default:
		return "unknown"
	}
}

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var outline [][]height
	idx := 0
	scr := bufio.NewScanner(f)
	for scr.Scan() {
		outline = append(outline, make([]height, len(scr.Bytes())))
		for i, b := range scr.Bytes() {
			outline[idx][i] = height(b - '0')
		}
		idx++
	}
	visible := 0
	for i, heights := range outline {
		for j := range heights {
			if i == 0 || j == 0 || i == len(outline)-1 || j == len(heights)-1 { // edges
				visible++
				continue
			}

			for _, dir := range []direction{left, right, up, down} {
				if isVisibleFrom(dir, i, j, outline) {
					visible++
					break
				}
			}
		}
	}
	fmt.Printf("visible trees: %v\n", visible)
}

func isVisibleFrom(d direction, y, x int, outline [][]height) bool {
	switch d {
	case up:
		for i := y - 1; i >= 0; i-- {
			if outline[i][x] >= outline[y][x] {
				return false
			}
		}
	case right:
		for i := x + 1; i < len(outline[y]); i++ {
			if outline[y][i] >= outline[y][x] {
				return false
			}
		}
	case down:
		for i := y + 1; i < len(outline); i++ {
			if outline[i][x] >= outline[y][x] {
				return false
			}
		}
	case left:
		for i := x - 1; i >= 0; i-- {
			if outline[y][i] >= outline[y][x] {
				return false
			}
		}
	}
	return true
}
