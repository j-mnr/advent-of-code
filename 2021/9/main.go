package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var r io.Reader
	var err error
	if false {
		r = strings.NewReader(`
2199943210
3987894921
9856789892
8767896789
9899965678
`[1:])
	} else {
		r, err = os.Open("input")
		if err != nil {
			panic(err)
		}
	}
	scr := bufio.NewScanner(r)
	hm := make(heightMap, 0, 100)
	var i uint
	for scr.Scan() {
		hm = append(hm, make([]uint8, 0, 100))
		for _, r := range scr.Text() {
			hm[i] = append(hm[i], uint8(r-'0'))
		}
		i++
	}
	lows := make([]uint8, 0, 10)
	for y, rows := range hm {
		for x, val := range rows {
			p := point{y: uint8(y), x: uint8(x)}
			for _, h := range hm.surrounding(p) {
				if uint8(h) <= hm[y][x] {
					goto next
				}
			}
			lows = append(lows, val)
		next:
		}
	}
	fmt.Println(lows)
	fmt.Println(riskLevel(lows))
}

func riskLevel(points []uint8) uint64 {
	var sum uint64
	for _, v := range points {
		sum = sum + 1 + uint64(v)
	}
	return sum
}

type heightMap [][]uint8

type point struct{ y, x uint8 }
type height uint8

func (h heightMap) surrounding(p point) []height {
	heights := make([]height, 0, 8)
	x, y := p.x, p.y
	if x != 0 && y != 0 { // top-left
		heights = append(heights, height(h[y-1][x-1]))
	}
	if y != 0 { // top
		heights = append(heights, height(h[y-1][x]))
	}
	if int(x+1) < len(h[0]) && y != 0 { // top-right
		heights = append(heights, height(h[y-1][x+1]))
	}
	if x != 0 { // left
		heights = append(heights, height(h[y][x-1]))
	}
	if int(x+1) < len(h[0]) { // right
		heights = append(heights, height(h[y][x+1]))
	}
	if x != 0 && int(y+1) < len(h) { // bottom-left
		heights = append(heights, height(h[y+1][x-1]))
	}
	if int(y+1) < len(h) { // bottom
		heights = append(heights, height(h[y+1][x]))
	}
	if int(x+1) < len(h[0]) && int(y+1) < len(h) { // bottom-right
		heights = append(heights, height(h[y+1][x+1]))
	}
	return heights
}
