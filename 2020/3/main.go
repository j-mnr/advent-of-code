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

	prd := 1
	for _, xy := range [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}} {
		prd *= traverse(xy[0], xy[1], area)
	}
	fmt.Println(prd)
}

func traverse(rightStep, downStep int, area []string) (total int) {
	for right, down := rightStep, downStep; down < len(area); {
		row := area[down]
		if row[right%len(row)] == '#' {
			total++
		}
		right, down = right+rightStep, down+downStep
	}
	return total
}
