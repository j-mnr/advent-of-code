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
	area := make([]string, 323)
	for i := 0; scr.Scan(); i++ {
		area[i] = scr.Text()
	}

	strip := len(area[0])
	trees := make(map[int]int)
	for _, z := range []int{1, 3, 5, 7} {
		if z == 1 {
			for x, y := 1, 2; y < len(area); x, y = x+1, y+2 {
				if y < len(area)-1 && area[y][x%strip] == '#' {
					fmt.Println(y, x, z)
					trees[2]++
				}
			}
		}
		for x, y := z, 1; y != len(area); x, y = x+z, y+1 {
			if area[y][x%strip] == '#' {
				trees[z]++
			}
		}
	}
	fmt.Println(trees)
	product := 1
	for _, k := range trees {
		product *= k
	}
	fmt.Println(product)
}
