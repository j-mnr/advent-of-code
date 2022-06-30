package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var f = strings.NewReader(`
..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#`[1:])

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	scr := bufio.NewScanner(f)
	area := make([]string, 0, 323)
	for scr.Scan() {
		area = append(area, scr.Text())
	}

	strip := len(area[0])
	trees := 0
	for x, y := 3, 1; y != len(area); x, y = x+3, y+1 {
		if area[y][x%strip] == '#' {
			trees++
		}
	}
	fmt.Println(trees)
}
