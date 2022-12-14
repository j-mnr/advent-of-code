package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type coord [2]int

var f = strings.NewReader(`
498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9`[1:])

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	filled := make(map[coord]struct{})

	scr := bufio.NewScanner(f)
	for scr.Scan() {
		coords := parseLine(scr.Text())
		for i := 1; i < len(coords); i++ {
			cx, cy := coords[i][0], coords[i][1]
			px, py := coords[i-1][0], coords[i-1][1]
			switch {
			case cy == py:
				for x := min(cx, px); x < max(cx, px)+1; x++ {
					filled[coord{x, cy}] = struct{}{}
				}
			case cx == px:
				for y := min(cy, py); y < max(cy, py)+1; y++ {
					filled[coord{cx, y}] = struct{}{}
				}
			default:
				panic("Either both X's should be equal or both Y's")
			}
		}
	}

	grains := 0
	for hasStopped(filled, findFloor(filled)) {
		grains++
	}
	fmt.Printf("grains of sand: %v\n", grains) // 198 too low
}

func hasStopped(filled map[coord]struct{}, floor int) bool {
	x, y := 500, 0
	for y <= floor {
		if _, in := filled[coord{x, y + 1}]; !in {
			y++
			continue
		}
		if _, in := filled[coord{x - 1, y + 1}]; !in {
			x--
			y++
			continue
		}
		if _, in := filled[coord{x + 1, y + 1}]; !in {
			x++
			y++
			continue
		}
		filled[coord{x, y}] = struct{}{}
		return true
	}
	return false
}

func findFloor(filled map[coord]struct{}) int {
	max := 0
	for c := range filled {
		if c[1] > max {
			max = c[1]
		}
	}
	return max
}

func parseLine(s string) []coord {
	var coords []coord
	for _, grp := range strings.Split(s, " -> ") {
		a, b, ok := strings.Cut(grp, ",")
		if !ok {
			panic("could not find , in " + grp)
		}
		x, err := strconv.Atoi(a)
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(b)
		if err != nil {
			panic(err)
		}
		coords = append(coords, coord{x, y})
	}
	return coords
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
